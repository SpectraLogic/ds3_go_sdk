package helpers

import (
    "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/ds3"
    "sync"
    "log"
    "io"
)

type getProducer struct {
    JobMasterObjectList *models.MasterObjectList  //MOL from put bulk job creation
    ReadObjects *[]GetObject
    queue *chan TransferOperation
    strategy *ReadTransferStrategy
    client *ds3.Client
    wg *sync.WaitGroup
    readObjectMap map[string]GetObject
    processedChunksTracker ProcessedChunksTracker
    waitingToBeTransferred BlobDescriptionQueue
}

func newGetProducer(jobMasterObjectList *models.MasterObjectList, readObjects *[]GetObject, queue *chan TransferOperation, strategy *ReadTransferStrategy, client *ds3.Client, wg *sync.WaitGroup) *getProducer {
    return &getProducer{
        JobMasterObjectList:jobMasterObjectList,
        ReadObjects:readObjects,
        queue:queue,
        strategy:strategy,
        client:client,
        wg:wg,
        readObjectMap:toReadObjectMap(readObjects),
        processedChunksTracker:NewProccessedChunksTracker(),
        waitingToBeTransferred:NewBlobDescriptionQueue(),
    }
}

func toReadObjectMap(readObjects *[]GetObject) map[string]GetObject {
    objectMap := make(map[string]GetObject)

    if readObjects == nil {
        return objectMap
    }

    for _, obj := range *readObjects {
        objectMap[obj.GetObject.Name] = obj
    }

    return objectMap
}

func (producer *getProducer) processChunk(curChunk *models.Objects, bucketName string, jobId string) {
    //todo
}

//TODO clean up struct
type getObjectInfo struct {
    objName string
    objOffset int64
    objLength int64
    writer io.Writer
    onDone func()
    bucketName string
    jobId string
}

func (producer *getProducer) transferOperationBuilder(info getObjectInfo) TransferOperation {
    return func() {
        defer info.onDone()

        log.Printf("TRANSFER: name='%s' offset=%d length=%d", info.objName, info.objOffset, info.objLength)

        getObjRequest := models.NewGetObjectRequest()
    }
}

// Attempts to transfer a single blob from the BP to the client. If the blob is not ready for transfer,
// then it is added to the waiting to transfer queue
func (producer *getProducer) transferBlob(blob *BlobDescription, bucketName string, jobId string) {
    curReadObj := producer.readObjectMap[blob.Name()]

    if !curReadObj.ChannelBuilder.IsChannelAvailable(blob.Offset()) {
        log.Printf("DEBUG channel is NOT available for getting blob %s offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())
        producer.waitingToBeTransferred.Push(blob)
        return
    }

    log.Printf("DEBUG channel is available for getting blob %s offset=%d length=%d", blob.Name(), blob.Offset(), blob.Length())
    // blob ready to be transferred
    writer, err := curReadObj.ChannelBuilder.GetChannel(blob.Offset())
    if err != nil {
        log.Printf("ERROR getting write channel for '%s': %s", blob.Name(), err.Error())
        return
    }

    // Create transfer operation
    objInfo := getObjectInfo {
        objName:        blob.Name(),
        objOffset:      blob.Offset(),
        objLength:      blob.Length(),
        writer:         writer,
        onDone:         curReadObj.ChannelBuilder.OnDone,
        bucketName:     bucketName,
        jobId:          jobId,
    }

    var transfer TransferOperation = producer.transferOperationBuilder(objInfo)

    // Increment wait group, and enqueue transfer operation
    producer.wg.Add(1)
    *producer.queue <- transfer
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
    totalChunkCount := len(producer.JobMasterObjectList.Objects)
    log.Printf("DEBUG totalChunks=%d processedChunks=%d", totalChunkCount, producer.processedChunksTracker.NumberOfProcessedChunks()) //todo delete

    // process all chunks and make sure all blobs are queued for transfer
    for producer.processedChunksTracker.NumberOfProcessedChunks() < totalChunkCount || producer.waitingToBeTransferred.Size() > 0 {
        // Get the list of available chunks that the server can receive. The server may
        // not be able to receive everything, so not all chunks will necessarily be
        // returned
        chunksReady := models.NewGetJobChunksReadyForClientProcessingSpectraS3Request(producer.JobMasterObjectList.JobId)
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
