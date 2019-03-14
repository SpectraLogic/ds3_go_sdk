package helpers

import (
    "log"
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "sync"
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
}

func newPutProducer(jobMasterObjectList *ds3Models.MasterObjectList, putObjects *[]helperModels.PutObject, queue *chan TransferOperation, strategy *WriteTransferStrategy, client *ds3.Client, waitGroup *sync.WaitGroup) *putProducer {
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
func (producer *putProducer) transferOperationBuilder(info putObjectInfo, aggErr *ds3Models.AggregateError) TransferOperation {
    return func() {
        reader, err := info.channelBuilder.GetChannel(info.blob.Offset())
        if err != nil {
            aggErr.Append(err)
            log.Printf("ERROR could not get reader for object with name='%s' offset=%d length=%d", info.blob.Name(), info.blob.Offset(), info.blob.Length())
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
            aggErr.Append(err)
            log.Printf("ERROR during transfer of %s: %s\n", info.blob.Name(), err.Error())
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
func (producer *putProducer) processChunk(curChunk *ds3Models.Objects, bucketName string, jobId string, aggErr *ds3Models.AggregateError) {
    log.Printf("DEBUG begin chunk processing %s", curChunk.ChunkId)

    // transfer blobs that are ready, and queue those that are waiting for channel
    for _, curObj := range curChunk.Objects {
        log.Printf("DEBUG queuing object in waiting to be processed %s offset=%d length=%d", *curObj.Name, curObj.Offset, curObj.Length)
        blob := helperModels.NewBlobDescription(*curObj.Name, curObj.Offset, curObj.Length)
        producer.queueBlobForTransfer(&blob, bucketName, jobId, aggErr)
    }
}

// Iterates through blobs that are waiting to be transferred and attempts to transfer.
// If successful, blob is removed from queue. Else, it is re-queued.
func (producer *putProducer) processWaitingBlobs(bucketName string, jobId string, aggErr *ds3Models.AggregateError) {
    // attempt to process all blobs in waiting to be transferred
    waitingBlobs := producer.deferredBlobQueue.Size()
    for i := 0; i < waitingBlobs; i++ {
        //attempt transfer
        curBlob, err := producer.deferredBlobQueue.Pop()
        log.Printf("DEBUG attempting to process %s offset=%d length=%d", curBlob.Name(), curBlob.Offset(), curBlob.Length())
        if err != nil {
            aggErr.Append(err)
            log.Printf("ERROR when attempting blob transfer: %s", err.Error())
        }
        producer.queueBlobForTransfer(curBlob, bucketName, jobId, aggErr)
    }
}

// Attempts to transfer a single blob. If the blob is not ready for transfer,
// it is added to the waiting to transfer queue.
func (producer *putProducer) queueBlobForTransfer(blob *helperModels.BlobDescription, bucketName string, jobId string, aggErr *ds3Models.AggregateError) {
    if producer.processedBlobTracker.IsProcessed(*blob) {
        return
    }

    curWriteObj := producer.writeObjectMap[blob.Name()]

    if !curWriteObj.ChannelBuilder.IsChannelAvailable(blob.Offset()) {
        log.Printf("DEBUG channel is NOT available for blob %s offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())
        // Not ready to be transferred
        producer.deferredBlobQueue.Push(blob)
        return
    }

    log.Printf("DEBUG channel is available for blob %s offset=%d length=%d", curWriteObj.PutObject.Name, blob.Offset(), blob.Length())
    // Blob ready to be transferred

    // Create transfer operation
    objInfo := putObjectInfo{
        blob:           *blob,
        channelBuilder: curWriteObj.ChannelBuilder,
        bucketName:     bucketName,
        jobId:          jobId,
    }

    var transfer TransferOperation = producer.transferOperationBuilder(objInfo, aggErr)

    // Increment wait group, and enqueue transfer operation
    producer.waitGroup.Add(1)
    *producer.queue <- transfer

    // Mark blob as processed
    producer.processedBlobTracker.MarkProcessed(*blob)
}

// This initiates the production of the transfer operations which will be consumed by a consumer running in a separate go routine.
// Each transfer operation will put one blob of content to the BP.
// Once all blobs have been queued to be transferred, the producer will finish, even if all operations have not been consumed yet.
func (producer *putProducer) run(aggErr *ds3Models.AggregateError) {
    defer producer.waitGroup.Done()
    defer close(*producer.queue)

    // determine number of blobs to be processed
    var totalBlobCount int64 = producer.totalBlobCount()
    log.Printf("DEBUG totalBlobs=%d processedBlobs=%d", totalBlobCount, producer.processedBlobTracker.NumberOfProcessedBlobs())

    // process all chunks and make sure all blobs are queued for transfer
    for producer.processedBlobTracker.NumberOfProcessedBlobs() < totalBlobCount || producer.deferredBlobQueue.Size() > 0 {
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
                producer.processChunk(&curChunk, *chunksReadyResponse.MasterObjectList.BucketName, chunksReadyResponse.MasterObjectList.JobId, aggErr)
            }

            // Attempt to transfer waiting blobs
            producer.processWaitingBlobs(*chunksReadyResponse.MasterObjectList.BucketName, chunksReadyResponse.MasterObjectList.JobId, aggErr)
        } else {
            // When no chunks are returned we need to sleep to allow for cache space to
            // be freed.
            producer.strategy.BlobStrategy.delay()
        }
    }
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
