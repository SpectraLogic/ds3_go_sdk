package helpers

import (
    "spectra/ds3_go_sdk/ds3"
    ds3Models "spectra/ds3_go_sdk/ds3/models"
    helperModels "spectra/ds3_go_sdk/helpers/models"
    "sync"
)

type putTransfernator struct {
    BucketName string
    WriteObjects *[]helperModels.PutObject
    Strategy *WriteTransferStrategy
    Client *ds3.Client
}

func newPutTransfernator(bucketName string, writeObjects *[]helperModels.PutObject, strategy *WriteTransferStrategy, client *ds3.Client) *putTransfernator {
    return &putTransfernator{
        BucketName:bucketName,
        WriteObjects:writeObjects,
        Strategy:strategy,
        Client:client,
    }
}

// Creates the bulk put request from the list of write objects and put bulk job options
func newBulkPutRequest(bucketName string, writeObjects *[]helperModels.PutObject, options WriteBulkJobOptions) *ds3Models.PutBulkJobSpectraS3Request {
    var putObjects []ds3Models.Ds3PutObject
    for _, obj := range *writeObjects {
        putObjects = append(putObjects, obj.PutObject)
    }

    bulkPut := ds3Models.NewPutBulkJobSpectraS3Request(bucketName, putObjects)
    if options.Aggregating != nil {
        bulkPut = bulkPut.WithAggregating(*options.Aggregating)
    }
    if options.ImplicitJobIdResolution != nil {
        bulkPut = bulkPut.WithImplicitJobIdResolution(*options.ImplicitJobIdResolution)
    }
    if options.MaxUploadSize != nil {
        bulkPut = bulkPut.WithMaxUploadSize(*options.MaxUploadSize)
    }
    if options.MinimizeSpanningAcrossMedia != nil {
        bulkPut = bulkPut.WithMinimizeSpanningAcrossMedia(*options.MinimizeSpanningAcrossMedia)
    }
    if options.Priority != nil {
        bulkPut = bulkPut.WithPriority(*options.Priority)
    }
    if options.VerifyAfterWrite != nil {
        bulkPut = bulkPut.WithVerifyAfterWrite(*options.VerifyAfterWrite)
    }
    if options.Name != nil {
        bulkPut = bulkPut.WithName(*options.Name)
    }
    if options.Force != nil && *options.Force {
        bulkPut = bulkPut.WithForce()
    }
    if options.IgnoreNamingConflicts != nil && *options.IgnoreNamingConflicts {
        bulkPut = bulkPut.WithIgnoreNamingConflicts()
    }
    return bulkPut
}

func (transfernator *putTransfernator) transfer() error {
    // create bulk put job
    bulkPut := newBulkPutRequest(transfernator.BucketName, transfernator.WriteObjects, transfernator.Strategy.Options)

    bulkPutResponse, err := transfernator.Client.PutBulkJobSpectraS3(bulkPut)
    //todo handle what is recoverable
    if err != nil {
        return err
    }

    // init queue, producer and consumer
    var waitGroup sync.WaitGroup

    queue := newOperationQueue(MaxQueueSize) //todo make composable
    producer := newPutProducer(&bulkPutResponse.MasterObjectList, transfernator.WriteObjects, &queue, transfernator.Strategy, transfernator.Client, &waitGroup)
    consumer := newConsumer(&queue, &waitGroup, transfernator.Strategy.BlobStrategy.maxConcurrentTransfers())

    // Wait for completion of producer-consumer goroutines
    waitGroup.Add(2)  // adding producer and consumer goroutines to wait group
    go producer.run() // producer will add to waitGroup for every blob added to queue, and each transfer performed will decrement from waitGroup
    go consumer.run()
    waitGroup.Wait()

    return nil
}

func (transfernator *putTransfernator) cancel() {
    //todo
    panic("Not yet implemented")
}