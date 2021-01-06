package helpers

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "github.com/SpectraLogic/ds3_go_sdk/sdk_log"
    "sync"
    "time"
)

type putProducer struct {
    JobMasterObjectList  *ds3Models.MasterObjectList //MOL from put bulk job creation
    WriteObjects         *[]helperModels.PutObject
    queue                *chan TransferOperation
    strategy             *WriteTransferStrategy
    client               *ds3.Client
    waitGroup            *sync.WaitGroup
    writeObjectMap       map[string]helperModels.PutObject
    processedBlobTracker blobTracker
    deferredBlobQueue    BlobDescriptionQueue // queue of blobs whose channels are not yet ready for transfer
    sdk_log.Logger

    // Conditional value that gets triggered when a blob has finished being transferred
    doneNotifier NotifyBlobDone
}

func newPutProducer(
    jobMasterObjectList *ds3Models.MasterObjectList,
    putObjects *[]helperModels.PutObject,
    queue *chan TransferOperation,
    strategy *WriteTransferStrategy,
    client *ds3.Client,
    waitGroup *sync.WaitGroup,
    doneNotifier NotifyBlobDone) *putProducer {

    return &putProducer{
        JobMasterObjectList:  jobMasterObjectList,
        WriteObjects:         putObjects,
        queue:                queue,
        strategy:             strategy,
        client:               client,
        waitGroup:            waitGroup,
        writeObjectMap:       toWriteObjectMap(putObjects),
        deferredBlobQueue:    NewBlobDescriptionQueue(),
        processedBlobTracker: newProcessedBlobTracker(),
        Logger:               client.Logger, // use the same logger as the client
        doneNotifier:         doneNotifier,
    }
}

// Creates a map of object name to PutObject struct
func toWriteObjectMap(putObjects *[]helperModels.PutObject) map[string]helperModels.PutObject {
    objectMap := make(map[string]helperModels.PutObject)

    if putObjects == nil {
        return objectMap
    }

    for _, obj := range *putObjects {
        objectMap[obj.PutObject.Name] = obj
    }

    return objectMap
}

// Information required to perform a put operation of a blob using the source channelBuilder to BP destination
type putObjectInfo struct {
    blob           helperModels.BlobDescription
    channelBuilder helperModels.ReadChannelBuilder
    bucketName     string
    jobId          string
}

// Creates the transfer operation that will perform the data upload of the specified blob to BP
func (producer *putProducer) transferOperationBuilder(info putObjectInfo) TransferOperation {
    return func() {
        // has this file fatally errored while transferring a different blob?
        if info.channelBuilder.HasFatalError() {
            // skip performing this blob transfer
            producer.Warningf("fatal error occurred previously on this file, skipping sending blob name='%s' offset=%d length=%d", info.blob.Name(), info.blob.Offset(), info.blob.Length())
            return
        }
        reader, err := info.channelBuilder.GetChannel(info.blob.Offset())
        if err != nil {
            producer.strategy.Listeners.Errored(info.blob.Name(), err)

            info.channelBuilder.SetFatalError(err)
            producer.Errorf("could not get reader for object with name='%s' offset=%d length=%d: %v", info.blob.Name(), info.blob.Offset(), info.blob.Length(), err)
            return
        }
        defer info.channelBuilder.OnDone(reader)

        sizedReader := NewIoReaderWithSizeDecorator(reader, info.blob.Length())
        putObjRequest := ds3Models.NewPutObjectRequest(info.bucketName, info.blob.Name(), sizedReader).
            WithJob(info.jobId).
            WithOffset(info.blob.Offset())

		producer.maybeAddMetadata(info, putObjRequest)

        _, err = producer.client.PutObject(putObjRequest)
        if err != nil {
            producer.strategy.Listeners.Errored(info.blob.Name(), err)

            info.channelBuilder.SetFatalError(err)
            producer.Errorf("problem during transfer of %s: %s", info.blob.Name(), err.Error())
        }
    }
}

func (producer *putProducer) maybeAddMetadata(info putObjectInfo, putObjRequest *ds3Models.PutObjectRequest) {
	metadataMap := producer.metadataFrom(info)

	if len(metadataMap) == 0 {
		return
	}

	for key, value := range metadataMap {
		putObjRequest.WithMetaData(key, value)
	}
}

func (producer *putProducer) metadataFrom(info putObjectInfo) map[string]string {
	result := map[string]string{}

	for _, objectToPut := range *producer.WriteObjects {
		if objectToPut.PutObject.Name == info.blob.Name() {
			result = objectToPut.Metadata
			break
		}
	}

	return result
}

// Processes all the blobs in a chunk and attempts to add them to the transfer queue.
// If a blob is not ready for transfer, then it is added to the waiting to be transferred queue.
// Returns the number of  blobs added to queue.
func (producer *putProducer) processChunk(curChunk *ds3Models.Objects, bucketName string, jobId string) (int, error) {
    processedCount := 0
    producer.Debugf("begin chunk processing %s", curChunk.ChunkId)

    // transfer blobs that are ready, and queue those that are waiting for channel
    for _, curObj := range curChunk.Objects {
        producer.Debugf("queuing object in waiting to be processed %s offset=%d length=%d", *curObj.Name, curObj.Offset, curObj.Length)
        blob := helperModels.NewBlobDescription(*curObj.Name, curObj.Offset, curObj.Length)
        blobQueued, err := producer.queueBlobForTransfer(&blob, bucketName, jobId)
        if err != nil {
            return 0, err
        }
        if blobQueued {
            processedCount++
        }
    }
    return processedCount, nil
}

// Iterates through blobs that are waiting to be transferred and attempts to transfer.
// If successful, blob is removed from queue. Else, it is re-queued.
// Returns the number of blobs added to queue.
func (producer *putProducer) processWaitingBlobs(bucketName string, jobId string) (int, error) {
    processedCount := 0

    // attempt to process all blobs in waiting to be transferred
    waitingBlobs := producer.deferredBlobQueue.Size()
    for i := 0; i < waitingBlobs; i++ {
        //attempt transfer
        curBlob, err := producer.deferredBlobQueue.Pop()
        if err != nil {
            //should not be possible to get here
            producer.Errorf("problem when getting next blob to be transferred: %s", err.Error())
            break
        }
        producer.Debugf("attempting to process %s offset=%d length=%d", curBlob.Name(), curBlob.Offset(), curBlob.Length())
        blobQueued, err := producer.queueBlobForTransfer(curBlob, bucketName, jobId)
        if err != nil {
            return 0, err
        }
        if blobQueued {
            processedCount++
        }
    }

    return processedCount, nil
}

// Attempts to transfer a single blob. If the blob is not ready for transfer,
// it is added to the waiting to transfer queue.
// Returns whether or not the blob was queued for transfer.
func (producer *putProducer) queueBlobForTransfer(blob *helperModels.BlobDescription, bucketName string, jobId string) (bool, error) {
    if producer.processedBlobTracker.IsProcessed(*blob) {
        return false, nil // this was already processed
    }

    curWriteObj, ok := producer.writeObjectMap[blob.Name()]
    if !ok {
        err := fmt.Errorf("failed to find object associated with blob in object map: %s offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())
        producer.Errorf("unrecoverable error: %v", err)
        producer.processedBlobTracker.MarkProcessed(*blob)
        return false, err // fatal error occurred
    }

    if curWriteObj.ChannelBuilder == nil {
        err := fmt.Errorf("failed to transfer object, it does not have a channel builder: %s", curWriteObj.PutObject.Name)
        producer.Errorf("unrecoverable error: %v", err)
        producer.processedBlobTracker.MarkProcessed(*blob)
        return false, err // fatal error occurred
    }

    if curWriteObj.ChannelBuilder.HasFatalError() {
        // a fatal error happened on a previous blob for this file, skip processing
        producer.Warningf("fatal error occurred while transferring previous blob on this file, skipping blob %s offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())
        producer.processedBlobTracker.MarkProcessed(*blob)
        return false, nil // not actually transferring this blob
    }

    if !curWriteObj.ChannelBuilder.IsChannelAvailable(blob.Offset()) {
        producer.Debugf("channel is not currently available for blob %s offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())
        // Not ready to be transferred
        producer.deferredBlobQueue.Push(blob)
        return false, nil // not ready to be sent
    }

    producer.Debugf("channel is available for blob %s offset=%d length=%d", curWriteObj.PutObject.Name, blob.Offset(), blob.Length())
    // Blob ready to be transferred

    // Create transfer operation
    objInfo := putObjectInfo{
        blob:           *blob,
        channelBuilder: curWriteObj.ChannelBuilder,
        bucketName:     bucketName,
        jobId:          jobId,
    }

    transfer := producer.transferOperationBuilder(objInfo)

    // Increment wait group, and enqueue transfer operation
    producer.waitGroup.Add(1)
    *producer.queue <- transfer

    // Mark blob as processed
    producer.processedBlobTracker.MarkProcessed(*blob)

    return true, nil
}

// This initiates the production of the transfer operations which will be consumed by a consumer running in a separate go routine.
// Each transfer operation will put one blob of content to the BP.
// Once all blobs have been queued to be transferred, the producer will finish, even if all operations have not been consumed yet.
func (producer *putProducer) run() error {
    defer close(*producer.queue)

    // determine number of blobs to be processed
    totalBlobCount := producer.totalBlobCount()
    producer.Debugf("job status totalBlobs=%d processedBlobs=%d", totalBlobCount, producer.processedBlobTracker.NumberOfProcessedBlobs())

    // process all chunks and make sure all blobs are queued for transfer
    for producer.hasMoreToProcess(totalBlobCount) {
        processedCount, err := producer.queueBlobsReadyForTransfer(totalBlobCount)
        if err != nil {
            return err
        }

        // If the last operation processed blobs, then wait for something to finish
        if processedCount > 0 {
            // wait for a done signal to be received
            producer.doneNotifier.Wait()
        } else if producer.hasMoreToProcess(totalBlobCount) {
            // nothing could be processed, cache is probably full, wait a bit before trying again
            time.Sleep(producer.strategy.BlobStrategy.delay())
        }
    }
    return nil
}

func (producer *putProducer) hasMoreToProcess(totalBlobCount int64) bool {
    return producer.processedBlobTracker.NumberOfProcessedBlobs() < totalBlobCount || producer.deferredBlobQueue.Size() > 0
}

// Returns the number of items queued for work.
func (producer *putProducer) queueBlobsReadyForTransfer(totalBlobCount int64) (int, error) {
    // Attempt to transfer waiting blobs
    processedCount, err := producer.processWaitingBlobs(*producer.JobMasterObjectList.BucketName, producer.JobMasterObjectList.JobId)
    if err != nil {
        return 0, err
    }

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
            justProcessedCount, err := producer.processChunk(&curChunk, *chunksReadyResponse.MasterObjectList.BucketName, chunksReadyResponse.MasterObjectList.JobId)
            if err != nil {
                return 0, err
            }
            processedCount += justProcessedCount
        }
    }
    return processedCount, nil
}

// Determines the number of blobs to be transferred.
func (producer *putProducer) totalBlobCount() int64 {
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
