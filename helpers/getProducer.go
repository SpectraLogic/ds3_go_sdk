package helpers

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "github.com/SpectraLogic/ds3_go_sdk/helpers/ranges"
    "github.com/SpectraLogic/ds3_go_sdk/sdk_log"
    "io"
    "sync"
    "time"
)

const timesToRetryGettingPartialBlob = 5

type getProducer struct {
    JobMasterObjectList  *ds3Models.MasterObjectList //MOL from put bulk job creation
    GetObjects           *[]helperModels.GetObject
    queue                *chan TransferOperation
    strategy             *ReadTransferStrategy
    client               *ds3.Client
    waitGroup            *sync.WaitGroup
    readObjectMap        map[string]helperModels.GetObject
    processedBlobTracker blobTracker
    deferredBlobQueue    BlobDescriptionQueue // queue of blobs whose channels are not yet ready for transfer
    rangeFinder          ranges.BlobRangeFinder
    sdk_log.Logger

    // Conditional value that gets triggered when a blob has finished being transferred
    doneNotifier NotifyBlobDone
}

func newGetProducer(
    jobMasterObjectList *ds3Models.MasterObjectList,
    getObjects *[]helperModels.GetObject,
    queue *chan TransferOperation,
    strategy *ReadTransferStrategy,
    client *ds3.Client,
    waitGroup *sync.WaitGroup,
    doneNotifier NotifyBlobDone) *getProducer {

    return &getProducer{
        JobMasterObjectList:  jobMasterObjectList,
        GetObjects:           getObjects,
        queue:                queue,
        strategy:             strategy,
        client:               client,
        waitGroup:            waitGroup,
        readObjectMap:        toReadObjectMap(getObjects),
        processedBlobTracker: newProcessedBlobTracker(),
        deferredBlobQueue:    NewBlobDescriptionQueue(),
        rangeFinder:          ranges.NewBlobRangeFinder(getObjects),
        Logger:               client.Logger, //use the same logger as the client
        doneNotifier:         doneNotifier,
    }
}

// Creates a map of object name to the GetObject struct
func toReadObjectMap(getObjects *[]helperModels.GetObject) map[string]helperModels.GetObject {
    objectMap := make(map[string]helperModels.GetObject)

    if getObjects == nil {
        return objectMap
    }

    for _, obj := range *getObjects {
        objectMap[obj.Name] = obj
    }

    return objectMap
}

// Processes all the blobs in a chunk that are ready for transfer from BP
// Returns the number of blobs queued for process
func (producer *getProducer) processChunk(curChunk *ds3Models.Objects, bucketName string, jobId string) int {
    producer.Debugf("begin chunk processing %s", curChunk.ChunkId)

    processedCount := 0
    // transfer blobs that are ready, and queue those that are waiting for channel
    for _, curObj := range curChunk.Objects {
        producer.Debugf("queuing object in waiting to be processed %s offset=%d length=%d", *curObj.Name, curObj.Offset, curObj.Length)
        blob := helperModels.NewBlobDescription(*curObj.Name, curObj.Offset, curObj.Length)
        if producer.queueBlobForTransfer(&blob, bucketName, jobId) {
            processedCount++
        }
    }
    return processedCount
}

// Information required to perform a get operation of a blob with BP as data source and channelBuilder as destination
type getObjectInfo struct {
    blob           *helperModels.BlobDescription
    channelBuilder helperModels.WriteChannelBuilder
    bucketName     string
    jobId          string
}

// Creates the transfer operation that will perform the data retrieval of the specified blob from BP
func (producer *getProducer) transferOperationBuilder(info getObjectInfo) TransferOperation {
    return func() {
        // has this file fatally errored while transferring a different blob?
        if info.channelBuilder.HasFatalError() {
            // skip performing this blob transfer
            producer.Warningf("fatal error occurred previously on this file, skipping retrieval of blob objectName='%s' offset=%d", info.blob.Name(), info.blob.Offset())
            return
        }
        blobRanges := producer.rangeFinder.GetRanges(info.blob.Name(), info.blob.Offset(), info.blob.Length())

        producer.Debugf("transferring objectName='%s' offset=%d ranges=%v", info.blob.Name(), info.blob.Offset(), blobRanges)

        getObjRequest := ds3Models.NewGetObjectRequest(info.bucketName, info.blob.Name()).
            WithOffset(info.blob.Offset()).
            WithJob(info.jobId)

        if len(blobRanges) > 0 {
            getObjRequest = getObjRequest.WithRanges(blobRanges...)
        }

        getObjResponse, err := producer.client.GetObject(getObjRequest)
        if err != nil {
            producer.strategy.Listeners.Errored(info.blob.Name(), err)
            info.channelBuilder.SetFatalError(err)
            producer.Errorf("unable to retrieve object '%s' at offset %d: %s", info.blob.Name(), info.blob.Offset(), err.Error())
            return
        }
        defer func() {
           err := getObjResponse.Content.Close()
           if err != nil {
               producer.Warningf("unable to close response body for object '%s' at offset '%d': %s", info.blob.Name(), info.blob.Offset(), err)
           }
        }()

        if len(blobRanges) == 0 {
            writer, err := info.channelBuilder.GetChannel(info.blob.Offset())
            if err != nil {
                producer.strategy.Listeners.Errored(info.blob.Name(), err)
                info.channelBuilder.SetFatalError(err)
                producer.Errorf("unable to read contents of object '%s' at offset '%d': %s", info.blob.Name(), info.blob.Offset(), err.Error())
                return
            }
            defer info.channelBuilder.OnDone(writer)
            bytesWritten, err := io.Copy(writer, getObjResponse.Content) //copy all content from response reader to destination writer
            if err != nil && err != io.ErrUnexpectedEOF {
                producer.strategy.Listeners.Errored(info.blob.Name(), err)
                info.channelBuilder.SetFatalError(err)
                producer.Errorf("unable to copy content of object '%s' at offset '%d' from source to destination: %s", info.blob.Name(), info.blob.Offset(), err.Error())
                return
            }
            if bytesWritten != info.blob.Length() {
                producer.Errorf("failed to copy all content of object '%s' at offset '%d': only wrote %d of %d bytes", info.blob.Name(), info.blob.Offset(), bytesWritten, info.blob.Length())
                err := GetRemainingBlob(producer.client, info.bucketName, info.blob, bytesWritten, writer, producer.Logger)
                if err != nil {
                    producer.strategy.Listeners.Errored(info.blob.Name(), err)
                    info.channelBuilder.SetFatalError(err)
                    producer.Errorf("unable to copy content of object '%s' at offset '%d' from source to destination: %s", info.blob.Name(), info.blob.Offset(), err.Error())
                }
            }
            return
        }

        // write the content of each range
        for _, r := range blobRanges {
            err := writeRangeToDestination(info.channelBuilder, r, getObjResponse.Content)
            if err != nil {
                producer.strategy.Listeners.Errored(info.blob.Name(), err)
                info.channelBuilder.SetFatalError(err)
                producer.Errorf("unable to write to destination channel for object '%s' with range '%v': %s", info.blob.Name(), r, err.Error())
            }
        }
    }
}

func GetRemainingBlob(client *ds3.Client, bucketName string, blob *helperModels.BlobDescription, amountAlreadyRetrieved int64, writer io.Writer, logger sdk_log.Logger) error {
    logger.Debugf("starting retry for fetching partial blob '%s' at offset '%d': amount to retrieve %d", blob.Name(), blob.Offset(), blob.Length() - amountAlreadyRetrieved)
    bytesRetrievedSoFar := amountAlreadyRetrieved
    timesRetried := 0
    rangeEnd := blob.Offset() + blob.Length() -1
    for bytesRetrievedSoFar < blob.Length() && timesRetried < timesToRetryGettingPartialBlob {
        rangeStart := blob.Offset() + bytesRetrievedSoFar
        bytesRetrievedThisRound, err := RetryGettingBlobRange(client, bucketName, blob.Name(), blob.Offset(), rangeStart, rangeEnd, writer, logger)
        if err != nil {
            logger.Errorf("failed to get object '%s' at offset '%d', range %d=%d attempt %d: %s", blob.Name(), blob.Offset(), rangeStart, rangeEnd, timesRetried, err.Error())
        }
        bytesRetrievedSoFar+= bytesRetrievedThisRound
        timesRetried++
    }

    if bytesRetrievedSoFar < blob.Length() {
        return fmt.Errorf("failed to copy all content of object '%s' at offset '%d': only wrote %d of %d bytes", blob.Name(), blob.Offset(), bytesRetrievedSoFar, blob.Length())
    }
    return nil
}

func RetryGettingBlobRange(client *ds3.Client, bucketName string, objectName string, blobOffset int64, rangeStart int64, rangeEnd int64, writer io.Writer, logger sdk_log.Logger) (int64, error) {
    // perform a naked get call for the rest of the blob that we originally failed to get
    partOfBlobToFetch := ds3Models.Range{
        Start: rangeStart,
        End:   rangeEnd,
    }
    getObjRequest := ds3Models.NewGetObjectRequest(bucketName, objectName).
        WithOffset(blobOffset).
        WithRanges(partOfBlobToFetch)

    getObjResponse, err := client.GetObject(getObjRequest)
    if err != nil {
        return 0, err
    }
    defer func() {
        err := getObjResponse.Content.Close()
        if err != nil {
            logger.Warningf("failed to close response body for get object '%s' with range %d-%d: %v", objectName, rangeStart, rangeEnd, err)
        }
    }()

    bytesWritten, err := io.Copy(writer, getObjResponse.Content) //copy all content from response reader to destination writer
    return bytesWritten, err
}

// Writes a range of a blob to its destination channel
func writeRangeToDestination(channelBuilder helperModels.WriteChannelBuilder, blobRange ds3Models.Range, content io.Reader) error {
    writer, err := channelBuilder.GetChannel(blobRange.Start)
    if err != nil {
        return err
    }
    defer channelBuilder.OnDone(writer)

    _, err = io.CopyN(writer, content, blobRange.End - blobRange.Start + 1) // copies the range from response reader into destination writer

    return err
}

// Attempts to transfer a single blob from the BP to the client. If the blob is not ready for transfer,
// then it is added to the waiting to transfer queue
// Returns whether or not the blob was queued for transfer
func (producer *getProducer) queueBlobForTransfer(blob *helperModels.BlobDescription, bucketName string, jobId string) bool {
    if producer.processedBlobTracker.IsProcessed(*blob) {
        return false // already been processed
    }

    curReadObj := producer.readObjectMap[blob.Name()]

    if curReadObj.ChannelBuilder.HasFatalError() {
        // a fatal error happened on a previous blob for this file, skip processing
        producer.Warningf("fatal error occurred while transferring previous blob on this file, skipping blob '%s' offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())
        producer.processedBlobTracker.MarkProcessed(*blob)
        return false // not going to process
    }

    if !curReadObj.ChannelBuilder.IsChannelAvailable(blob.Offset()) {
        producer.Debugf("channel is not currently available for getting blob '%s' offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())
        producer.deferredBlobQueue.Push(blob)
        return false // not ready to be processed
    }

    producer.Debugf("channel is available for getting blob '%s' offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())

    // Create transfer operation
    objInfo := getObjectInfo{
        blob:           blob,
        channelBuilder: curReadObj.ChannelBuilder,
        bucketName:     bucketName,
        jobId:          jobId,
    }

    var transfer TransferOperation = producer.transferOperationBuilder(objInfo)

    // Increment wait group, and enqueue transfer operation
    producer.waitGroup.Add(1)
    *producer.queue <- transfer

    // Mark blob as processed
    producer.processedBlobTracker.MarkProcessed(*blob)

    return true
}

// Attempts to process all blobs whose channels were not available for transfer.
// Blobs whose channels are still not available are placed back on the queue.
// Returns the number of blobs queued for processing.
func (producer *getProducer) processWaitingBlobs(bucketName string, jobId string) int {
    processedCount := 0

    // attempt to process all blobs in waiting to be transferred
    waitingBlobs := producer.deferredBlobQueue.Size()
    for i := 0; i < waitingBlobs; i++ {
        //attempt transfer
        curBlob, err := producer.deferredBlobQueue.Pop()
        producer.Debugf("attempting to process '%s' offset=%d length=%d", curBlob.Name(), curBlob.Offset(), curBlob.Length())
        if err != nil {
            //should not be possible to get here
            producer.Errorf("failure during blob transfer '%s' at offset %d: %s", curBlob.Name(), curBlob.Offset(), err.Error())
            break
        }
        if producer.queueBlobForTransfer(curBlob, bucketName, jobId) {
            processedCount++
        }
    }
    return processedCount
}

// This initiates the production of the transfer operations which will be consumed by a consumer running in a separate go routine.
// Each transfer operation will retrieve one blob of content from the BP.
// Once all blobs have been queued to be transferred, the producer will finish, even if all operations have not been consumed yet.
func (producer *getProducer) run() error {
    defer close(*producer.queue)

    // determine number of blobs to be processed
    var totalBlobCount int64 = producer.totalBlobCount()
    producer.Debugf("job status totalBlobs=%d processedBlobs=%d", totalBlobCount, producer.processedBlobTracker.NumberOfProcessedBlobs())

    // process all chunks and make sure all blobs are queued for transfer
    for producer.hasMoreToProcess(totalBlobCount) {
        processedCount, err := producer.queueBlobsReadyForTransfer(totalBlobCount)
        if err != nil {
            return err
        }

        // If the last operation processed blobs, then wait for something to finish
        if processedCount > 0 {
            producer.doneNotifier.Wait()
        } else if producer.hasMoreToProcess(totalBlobCount) {
            // nothing could be processed, cache is probably full, wait a bit before trying again
            time.Sleep(producer.strategy.BlobStrategy.delay())
        }
    }
    return nil
}

func (producer *getProducer) hasMoreToProcess(totalBlobCount int64) bool {
    return producer.processedBlobTracker.NumberOfProcessedBlobs() < totalBlobCount || producer.deferredBlobQueue.Size() > 0
}

// Returns the number of blobs that have been queued for transfer
func (producer *getProducer) queueBlobsReadyForTransfer(totalBlobCount int64) (int, error) {
    // Attempt to transfer waiting blobs
    processedCount := producer.processWaitingBlobs(*producer.JobMasterObjectList.BucketName, producer.JobMasterObjectList.JobId)

    // Check if we need to query the BP for allocated blobs, or if we already know everything is allocated.
    if int64(producer.deferredBlobQueue.Size()) + producer.processedBlobTracker.NumberOfProcessedBlobs() >= totalBlobCount {
        // Everything is already allocated, no need to query BP for allocated chunks
        return processedCount, nil
    }

    // Get the list of available chunks that the server can receive. The server may
    // not be able to receive everything, so not all chunks will necessarily be
    // returned
    chunksReady := ds3Models.NewGetJobChunksReadyForClientProcessingSpectraS3Request(producer.JobMasterObjectList.JobId)
    chunksReadyResponse, err := producer.client.GetJobChunksReadyForClientProcessingSpectraS3(chunksReady)
    if err != nil {
        producer.Errorf("unrecoverable error: %v", err)
        return processedCount, err
    }

    // Check to see if any chunks can be processed
    numberOfChunks := len(chunksReadyResponse.MasterObjectList.Objects)
    if numberOfChunks > 0 {
        // Loop through all the chunks that are available for processing, and send
        // the files that are contained within them.
        for _, curChunk := range chunksReadyResponse.MasterObjectList.Objects {
            processedCount += producer.processChunk(&curChunk, *chunksReadyResponse.MasterObjectList.BucketName, chunksReadyResponse.MasterObjectList.JobId)
        }
    }
    return processedCount, nil
}

// Determines the number of blobs to be transferred.
func (producer *getProducer) totalBlobCount() int64 {
    if producer.JobMasterObjectList.Objects == nil || len(producer.JobMasterObjectList.Objects) == 0 {
        return 0
    }

    var count int64 = 0
    for _, chunk := range producer.JobMasterObjectList.Objects {
        for range chunk.Objects {
            count++
        }
    }
    return count
}