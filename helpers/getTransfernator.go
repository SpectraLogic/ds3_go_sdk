package helpers

import (
    "spectra/ds3_go_sdk/ds3"
    "spectra/ds3_go_sdk/ds3/models"
    "sync"
)

type getTransfernator struct {
    BucketName string
    ReadObjects *[]GetObject
    Strategy *ReadTransferStrategy
    Client *ds3.Client
}

func newGetTransfernator(bucketName string, readObjects *[]GetObject, strategy *ReadTransferStrategy, client *ds3.Client) *getTransfernator {
    return &getTransfernator{
        BucketName:bucketName,
        ReadObjects:readObjects,
        Strategy:strategy,
        Client:client,
    }
}

// Creates the bulk get request from the list of write objects and get bulk job options
func newBulkGetRequest(bucketName string, readObjects *[]GetObject, options ReadBulkJobOptions) *models.GetBulkJobSpectraS3Request {
    var getObjects []models.Ds3GetObject
    for _, obj := range *readObjects {
        getObjects = append(getObjects, obj.GetObject)
    }

    bulkGet := models.NewGetBulkJobSpectraS3RequestWithPartialObjects(bucketName, getObjects)
    if options.Aggregating != nil {
        bulkGet.WithAggregating(*options.Aggregating)
    }
    if options.ChunkClientProcessingOrderGuarantee != models.UNDEFINED {
        bulkGet.WithChunkClientProcessingOrderGuarantee(options.ChunkClientProcessingOrderGuarantee)
    }
    if options.ImplicitJobIdResolution != nil {
        bulkGet.WithImplicitJobIdResolution(*options.ImplicitJobIdResolution)
    }
    if options.priority != models.UNDEFINED {
        bulkGet.WithPriority(options.priority)
    }

    return bulkGet
}

func (transfernator *getTransfernator) transfer() error {
    // create bulk get job
    bulkGet := newBulkGetRequest(transfernator.BucketName, transfernator.ReadObjects, transfernator.Strategy.Options)

    bulkGetResponse, err := transfernator.Client.GetBulkJobSpectraS3(bulkGet)
    // todo handle what is recoverable
    if err != nil {
        return err
    }

    // init queue, producer and consumer
    var wg sync.WaitGroup

    queue := newOperationQueue(MaxQueueSize) //todo make composable
    producer := newGetProducer(&bulkGetResponse.MasterObjectList, transfernator.ReadObjects, &queue, transfernator.Strategy, transfernator.Client, &wg)
    consumer := newPutConsumer(&queue, &wg, transfernator.Strategy.BlobStrategy.maxTransferGoroutines()) //todo rename??

    // Wait for completion of producer-consumer goroutines
    wg.Add(2) // adding producer and consumer goroutines to wait group
    go producer.run() // producer will add to wg for every blob retrieval added to queue, and each transfer performed will decrement from wg
    go consumer.run()
    wg.Wait()

    return nil
}