package helpers

import (
    ds3Models "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/ds3"
    "sync"
    "log"
    "io"
    "spectra/ds3_go_sdk/helpers/ranges"
    helperModels "spectra/ds3_go_sdk/helpers/models"
)

type getProducer struct {
    JobMasterObjectList    *ds3Models.MasterObjectList //MOL from put bulk job creation
    GetObjects             *[]helperModels.GetObject
    queue                  *chan TransferOperation
    strategy               *ReadTransferStrategy
    client                 *ds3.Client
    wg                     *sync.WaitGroup
    readObjectMap          map[string]helperModels.GetObject
    processedBlobTracker   blobTracker
    waitingToBeTransferred BlobDescriptionQueue
    rangeFinder            ranges.BlobRangeFinder
}

func newGetProducer(jobMasterObjectList *ds3Models.MasterObjectList, getObjects *[]helperModels.GetObject, queue *chan TransferOperation, strategy *ReadTransferStrategy, client *ds3.Client, wg *sync.WaitGroup) *getProducer {
    return &getProducer{
        JobMasterObjectList:    jobMasterObjectList,
        GetObjects:             getObjects,
        queue:                  queue,
        strategy:               strategy,
        client:                 client,
        wg:                     wg,
        readObjectMap:          toReadObjectMap(getObjects),
        processedBlobTracker:   newProcessedBlobTracker(),
        waitingToBeTransferred: NewBlobDescriptionQueue(),
        rangeFinder:            ranges.NewBlobRangeFinder(getObjects),
    }
}

//todo delete
func toReadObjectMap(readObjects *[]helperModels.GetObject) map[string]helperModels.GetObject {
    objectMap := make(map[string]helperModels.GetObject)

    if readObjects == nil {
        return objectMap
    }

    for _, obj := range *readObjects {
        objectMap[obj.Name] = obj
    }

    return objectMap
}

func (producer *getProducer) processChunk(curChunk *ds3Models.Objects, bucketName string, jobId string) {
    log.Printf("DEBUG begin chunk processing %s", curChunk.ChunkId) //todo delete

    // transfer blobs that are ready, and queue those that are waiting for channel
    for _, curObj := range curChunk.Objects {
        log.Printf("DEBUG queuing object in waiting to be processed %s offset=%d length=%d", *curObj.Name, curObj.Offset, curObj.Length) //todo delete
        blob := helperModels.NewBlobDescription(*curObj.Name, curObj.Offset, curObj.Length)
        producer.transferBlob(&blob, bucketName, jobId)
    }
}

//TODO clean up struct
type getBlobInfo struct {
    blob           *helperModels.BlobDescription
    channelBuilder helperModels.WriteChannelBuilder
    bucketName     string
    jobId          string
}

func (producer *getProducer) transferOperationBuilder(info getBlobInfo) TransferOperation {
    return func() {

        blobRanges := producer.rangeFinder.GetRanges(info.blob.Name(), info.blob.Offset(), info.blob.Length())

        log.Printf("TRANSFER: objectName='%s' offset=%d ranges=%v", info.blob.Name(), info.blob.Offset(), blobRanges)

        getObjRequest := ds3Models.NewGetObjectRequest(info.bucketName, info.blob.Name()).
            WithOffset(info.blob.Offset()).
            WithJob(info.jobId)

        if len(blobRanges) > 0 {
            getObjRequest = getObjRequest.WithRanges(blobRanges...)
        }

        getObjResponse, err := producer.client.GetObject(getObjRequest)
        if err != nil {
            log.Printf("ERROR during retrieval of %s: %s", info.blob.Name(), err.Error())
            return
        }

        if len(blobRanges) == 0 {
            writer, err := info.channelBuilder.GetChannel(info.blob.Offset())
            if err != nil {
                //todo handle error
                log.Printf("ERROR when copying content for object '%s' at offset '%d': %s", info.blob.Name(), info.blob.Offset(), err.Error())
                return
            }
            defer info.channelBuilder.OnDone(writer)
            io.Copy(writer, getObjResponse.Content) //copy all content from response reader to destination writer
            return
        }

        // write the content of each range
        for _, r := range blobRanges {
            err := writeRangeToDestination(info.channelBuilder, r, getObjResponse.Content)
            if err != nil {
                log.Printf("ERROR when writing to destination channel for object '%s' with range '%v': %s", info.blob.Name(), r, err.Error())
            }
        }

    }
}

func writeRangeToDestination(channelBuilder helperModels.WriteChannelBuilder, r ds3Models.Range, content io.Reader) error {
    //todo handle when channel is not available yet
    writer, err := channelBuilder.GetChannel(r.Start)
    defer channelBuilder.OnDone(writer)

    if err != nil {
        return err
    }

    _, err = io.CopyN(writer, content, r.End - r.Start + 1) // copies the range from response reader into destination writer

    if err != nil {
        return err
    }

    return nil
}

// Attempts to transfer a single blob from the BP to the client. If the blob is not ready for transfer,
// then it is added to the waiting to transfer queue
func (producer *getProducer) transferBlob(blob *helperModels.BlobDescription, bucketName string, jobId string) {
    if producer.processedBlobTracker.IsProcessed(*blob) {
        return
    }

    curReadObj := producer.readObjectMap[blob.Name()]

    if !curReadObj.ChannelBuilder.IsChannelAvailable(blob.Offset()) {
        log.Printf("DEBUG channel is NOT available for getting blob %s offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())
        producer.waitingToBeTransferred.Push(blob)
        return
    }

    log.Printf("DEBUG channel is available for getting blob %s offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())

    // Create transfer operation
    objInfo := getBlobInfo{
        blob:           blob,
        channelBuilder: curReadObj.ChannelBuilder,
        bucketName:     bucketName,
        jobId:          jobId,
    }

    var transfer TransferOperation = producer.transferOperationBuilder(objInfo)

    // Increment wait group, and enqueue transfer operation
    producer.wg.Add(1)
    *producer.queue <- transfer

    // Mark blob as processed
    producer.processedBlobTracker.MarkProcessed(*blob)
}

//TODO reusable?
func (producer *getProducer) transferWaitingBlobs(bucketName string, jobId string) {
    // attempt to process all blobs in waiting to be transferred
    waitingChunks := producer.waitingToBeTransferred.Size()
    for i := 0; i < waitingChunks; i++ {
        //attempt transfer
        curBlob, err := producer.waitingToBeTransferred.Pop()
        log.Printf("DEBUG attempting to process %s offset=%d length=%d", curBlob.Name(), curBlob.Offset(), curBlob.Length()) //todo delete
        if err != nil {
            //todo should not be possible to get here
            log.Printf("ERROR when attempting blob transfer: %s", err.Error())
        }
        producer.transferBlob(curBlob, bucketName, jobId)
    }
}

//TODO can re-arrange to be composable, where run() and transferWaitingBlobs() implementation is shared by producers, and processChunk() and transferBlob() is unique to get/put
func (producer *getProducer) run() {
    defer producer.wg.Done()
    defer close(*producer.queue)

    // determine number of chunks to be processed
    var totalBlobCount int64 = producer.totalBlobCount()
    log.Printf("DEBUG totalBlobs=%d processedBlobs=%d", totalBlobCount, producer.processedBlobTracker.NumberOfProcessedBlobs()) //todo delete

    // process all chunks and make sure all blobs are queued for transfer
    for producer.processedBlobTracker.NumberOfProcessedBlobs() < totalBlobCount || producer.waitingToBeTransferred.Size() > 0 {
        // Get the list of available chunks that the server can receive. The server may
        // not be able to receive everything, so not all chunks will necessarily be
        // returned
        chunksReady := ds3Models.NewGetJobChunksReadyForClientProcessingSpectraS3Request(producer.JobMasterObjectList.JobId)
        chunksReadyResponse, err := producer.client.GetJobChunksReadyForClientProcessingSpectraS3(chunksReady)
        if err != nil {
            log.Fatal(err)
        }

        // Check to see if any chunks can be processed
        numberOfChunks := len(chunksReadyResponse.MasterObjectList.Objects)
        if numberOfChunks > 0 {
            // Loop through all the chunks that are available for processing, and send
            // the files that are contained within them.
            for _, curChunk := range chunksReadyResponse.MasterObjectList.Objects {
                producer.processChunk(&curChunk, *chunksReadyResponse.MasterObjectList.BucketName, chunksReadyResponse.MasterObjectList.JobId)
            }

            // Attempt to transfer waiting blobs
            producer.transferWaitingBlobs(*chunksReadyResponse.MasterObjectList.BucketName, chunksReadyResponse.MasterObjectList.JobId)
        } else {
            // When no chunks are returned we need to sleep to allow for cache space to
            // be freed.
            producer.strategy.BlobStrategy.delay()
        }
    }
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