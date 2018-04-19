package helpers

import (
    "spectra/ds3_go_sdk/ds3"
    ds3Models "spectra/ds3_go_sdk/ds3/models"
    helperModels "spectra/ds3_go_sdk/helpers/models"
    "sync"
)

type getTransfernator struct {
    BucketName string
    ReadObjects *[]helperModels.GetObject
    Strategy *ReadTransferStrategy
    Client *ds3.Client
}

func newGetTransfernator(bucketName string, readObjects *[]helperModels.GetObject, strategy *ReadTransferStrategy, client *ds3.Client) *getTransfernator {
    return &getTransfernator{
        BucketName:bucketName,
        ReadObjects:readObjects,
        Strategy:strategy,
        Client:client,
    }
}

// Creates the bulk get request from the list of write objects and get bulk job options
func newBulkGetRequest(bucketName string, readObjects *[]helperModels.GetObject, options ReadBulkJobOptions) *ds3Models.GetBulkJobSpectraS3Request {
    var getObjects []ds3Models.Ds3GetObject
    for _, obj := range *readObjects {
        getObjects = append(getObjects, createPartialGetObjects(obj)...)
    }

    bulkGet := ds3Models.NewGetBulkJobSpectraS3RequestWithPartialObjects(bucketName, getObjects)
    if options.Aggregating != nil {
        bulkGet.WithAggregating(*options.Aggregating)
    }
    if options.ChunkClientProcessingOrderGuarantee != ds3Models.UNDEFINED {
        bulkGet.WithChunkClientProcessingOrderGuarantee(options.ChunkClientProcessingOrderGuarantee)
    }
    if options.ImplicitJobIdResolution != nil {
        bulkGet.WithImplicitJobIdResolution(*options.ImplicitJobIdResolution)
    }
    if options.priority != ds3Models.UNDEFINED {
        bulkGet.WithPriority(options.priority)
    }

    return bulkGet
}

// Converts a GetObject into its corresponding Ds3GetObjects for use in bulk get request building.
func createPartialGetObjects(getObject helperModels.GetObject) []ds3Models.Ds3GetObject {
    // handle getting the entire object
    if len(getObject.Ranges) == 0 {
        return []ds3Models.Ds3GetObject { { Name:getObject.Name }, }
    }
    // handle partial object retrieval
    var partialObjects []ds3Models.Ds3GetObject
    for _, r := range getObject.Ranges {
        offset := r.Start
        length := r.End - r.Start + 1
        partialObjects = append(partialObjects, ds3Models.Ds3GetObject{Name:getObject.Name, Offset:&offset, Length:&length})
    }
    return partialObjects
}

func (transfernator *getTransfernator) transfer() error {
    // create bulk get job
    bulkGet := newBulkGetRequest(transfernator.BucketName, transfernator.ReadObjects, transfernator.Strategy.Options)

    bulkGetResponse, err := transfernator.Client.GetBulkJobSpectraS3(bulkGet)
    if err != nil {
        return err
    }

    // init queue, producer and consumer
    var waitGroup sync.WaitGroup

    queue := newOperationQueue(transfernator.Strategy.BlobStrategy.maxWaitingTransfers())
    producer := newGetProducer(&bulkGetResponse.MasterObjectList, transfernator.ReadObjects, &queue, transfernator.Strategy, transfernator.Client, &waitGroup)
    consumer := newConsumer(&queue, &waitGroup, transfernator.Strategy.BlobStrategy.maxConcurrentTransfers())

    // Wait for completion of producer-consumer goroutines
    var aggErr ds3Models.AggregateError
    waitGroup.Add(2)  // adding producer and consumer goroutines to wait group
    go producer.run(&aggErr) // producer will add to waitGroup for every blob retrieval added to queue, and each transfer performed will decrement from waitGroup
    go consumer.run()
    waitGroup.Wait()

    return aggErr.GetErrors()
}