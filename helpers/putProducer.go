package helpers

import (
    "log"
    ds3Models "spectra/ds3_go_sdk/ds3/models"
    helperModels "spectra/ds3_go_sdk/helpers/models"
    "spectra/ds3_go_sdk/ds3"
    "sync"
)

type putProducer struct {
    JobMasterObjectList *ds3Models.MasterObjectList //MOL from put bulk job creation
    WriteObjects *[]helperModels.PutObject
    queue *chan TransferOperation
    strategy *WriteTransferStrategy
    client *ds3.Client
    wg *sync.WaitGroup
    writeObjectMap map[string]helperModels.PutObject
    processedChunksTracker ProcessedChunksTracker
    waitingToBeTransferred BlobDescriptionQueue
}

func newPutProducer(jobMasterObjectList *ds3Models.MasterObjectList, writeObjects *[]helperModels.PutObject, queue *chan TransferOperation, strategy *WriteTransferStrategy, client *ds3.Client, wg *sync.WaitGroup) *putProducer {
    return &putProducer{
        JobMasterObjectList:jobMasterObjectList,
        WriteObjects:writeObjects,
        queue:queue,
        strategy:strategy,
        client:client,
        wg:wg,
        writeObjectMap:toWriteObjectMap(writeObjects),
        processedChunksTracker:NewProccessedChunksTracker(),
        waitingToBeTransferred:NewBlobDescriptionQueue(),
    }
}

func toWriteObjectMap(writeObjects *[]helperModels.PutObject) map[string]helperModels.PutObject {
    objectMap := make(map[string]helperModels.PutObject)

    if writeObjects == nil {
        return objectMap
    }

    for _, obj := range *writeObjects {
        objectMap[obj.PutObject.Name] = obj
    }

    return objectMap
}

//TODO clean up struct
type putObjectInfo struct {
    objName    string
    objOffset  int64
    objLength  int64
    channelBuilder helperModels.ReadChannelBuilder
    bucketName string
    jobId      string
}

func (pp *putProducer) transferOperationBuilder(info putObjectInfo) TransferOperation {
    return func() {
        reader, err := info.channelBuilder.GetChannel(info.objOffset)
        if err != nil {
            log.Printf("ERROR could not get reader for object with name='%s' offset=%d length=%d", info.objName, info.objOffset, info.objLength)
        }
        defer info.channelBuilder.OnDone(reader)

        sizedReader := NewIoReaderWithSizeDecorator(reader, info.objLength)
        putObjRequest := ds3Models.NewPutObjectRequest(info.bucketName, info.objName, sizedReader).
            WithJob(info.jobId).
            WithOffset(info.objOffset)

        _, err = pp.client.PutObject(putObjRequest)
        if err != nil {
            log.Printf("Error during transfer of %s: %s\n", info.objName, err.Error()) //todo handle error better
        }
    }
}

// Processes all the blobs in a chunk and attempts to add them to the transfer queue.
// If a blob is not ready for transfer, then it is added to the waiting to be transferred queue.
func (pp *putProducer) processChunk(curChunk *ds3Models.Objects, bucketName string, jobId string) {
    log.Printf("DEBUG begin chunk processing %s", curChunk.ChunkId) //todo delete
    if !pp.processedChunksTracker.IsProcessed(curChunk.ChunkId) {
        // transfer blobs that are ready, and queue those that are waiting for channel
        for _, curObj := range curChunk.Objects {
            log.Printf("DEBUG queuing object in waiting to be processed %s offset=%d length=%d", *curObj.Name, curObj.Offset, curObj.Length) //todo delete
            blob := helperModels.NewBlobDescription(*curObj.Name, curObj.Offset, curObj.Length)
            pp.transferBlob(&blob, bucketName, jobId)
        }

        //Mark chunk as processed
        pp.processedChunksTracker.SetProcessed(curChunk.ChunkId)
    }


}

// Iterates through blobs that are waiting to be transferred and attempts to transfer.
// If successful, blob is removed from queue. Else, it is re-queued.
func (pp *putProducer) transferWaitingBlobs(bucketName string, jobId string) {
    // attempt to process all blobs in waiting to be transferred
    waitingChunks := pp.waitingToBeTransferred.Size()
    for i := 0; i < waitingChunks; i++ {
        //attempt transfer
        curBlob, err := pp.waitingToBeTransferred.Pop()
        log.Printf("DEBUG attempting to process %s offset=%d length=%d", curBlob.Name(), curBlob.Offset(), curBlob.Length()) //todo delete
        if err != nil {
            //todo should not be possible to get here
            log.Printf("ERROR when attempting blob transfer: %s", err.Error())
        }
        pp.transferBlob(curBlob, bucketName, jobId)
    }
}

// Attempts to transfer a single blob. If the blob is not ready for transfer,
// it is added to the waiting to transfer queue.
func (pp *putProducer) transferBlob(blob *helperModels.BlobDescription, bucketName string, jobId string) {
    curWriteObj := pp.writeObjectMap[blob.Name()]
    if curWriteObj.ChannelBuilder.IsChannelAvailable(blob.Offset()) {
        log.Printf("DEBUG channel is available for blob %s offset=%d length=%d", curWriteObj.PutObject.Name, blob.Offset(), blob.Length()) //todo delete
        // Blob ready to be transferred

        // Create transfer operation
        objInfo := putObjectInfo{
            objName:        blob.Name(),
            objOffset:      blob.Offset(),
            objLength:      blob.Length(),
            channelBuilder: curWriteObj.ChannelBuilder,
            bucketName:     bucketName,
            jobId:          jobId,
        }

        var transfer TransferOperation = pp.transferOperationBuilder(objInfo)

        // Increment wait group, and enqueue transfer operation
        pp.wg.Add(1)
        *pp.queue <- transfer
    } else {
        log.Printf("DEBUG channel is NOT available for blob %s offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())
        // Not ready to be transferred
        pp.waitingToBeTransferred.Push(blob)
    }
}

func (pp *putProducer) run() {
    defer pp.wg.Done()
    defer close(*pp.queue)

    // determine number of chunks to be processed
    totalChunkCount := len(pp.JobMasterObjectList.Objects)
    log.Printf("DEBUG totalChunks=%d processedChunks=%d", totalChunkCount, pp.processedChunksTracker.NumberOfProcessedChunks()) //todo delete

    // process all chunks and make sure all blobs are queued for transfer
    for pp.processedChunksTracker.NumberOfProcessedChunks() < totalChunkCount || pp.waitingToBeTransferred.Size() > 0 {
        // Get the list of available chunks that the server can receive. The server may
        // not be able to receive everything, so not all chunks will necessarily be
        // returned
        chunksReady := ds3Models.NewGetJobChunksReadyForClientProcessingSpectraS3Request(pp.JobMasterObjectList.JobId)
        chunksReadyResponse, err := pp.client.GetJobChunksReadyForClientProcessingSpectraS3(chunksReady)
        if err != nil {
            log.Fatal(err)
        }

        // Check to see if any chunks can be processed
        numberOfChunks := len(chunksReadyResponse.MasterObjectList.Objects)
        if numberOfChunks > 0 {
            // Loop through all the chunks that are available for processing, and send
            // the files that are contained within them.
            for _, curChunk := range chunksReadyResponse.MasterObjectList.Objects {
                pp.processChunk(&curChunk, *chunksReadyResponse.MasterObjectList.BucketName, chunksReadyResponse.MasterObjectList.JobId)
            }

            // Attempt to transfer waiting blobs
            pp.transferWaitingBlobs(*chunksReadyResponse.MasterObjectList.BucketName, chunksReadyResponse.MasterObjectList.JobId)
        } else {
            // When no chunks are returned we need to sleep to allow for cache space to
            // be freed.
            pp.strategy.BlobStrategy.delay()
        }
    }
}
