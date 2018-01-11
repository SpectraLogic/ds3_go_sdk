package helpers

import (
    "spectra/ds3_go_sdk/ds3"
    "spectra/ds3_go_sdk/ds3/models"
    "sync"
)

type putTransfernator struct {
    BucketName string
    WriteObjects *[]PutObject
    Strategy *WriteTransferStrategy
    Client *ds3.Client
}

func newPutTransfernator(bucketName string, writeObjects *[]PutObject, strategy *WriteTransferStrategy, client *ds3.Client) *putTransfernator {
    return &putTransfernator{
        BucketName:bucketName,
        WriteObjects:writeObjects,
        Strategy:strategy,
        Client:client,
    }
}

// Creates the bulk put request from the list of write objects and put bulk job options
func newBulkPutRequest(bucketName string, writeObjects *[]PutObject, options WriteBulkJobOptions) *models.PutBulkJobSpectraS3Request {
    var putObjects []models.Ds3PutObject
    for _, obj := range *writeObjects {
        putObjects = append(putObjects, obj.PutObject)
    }

    bulkPut := models.NewPutBulkJobSpectraS3Request(bucketName, putObjects)
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

func (pt *putTransfernator) transfer() error {
    // create bulk put job
    bulkPut := newBulkPutRequest(pt.BucketName, pt.WriteObjects, pt.Strategy.Options)

    bulkPutResponse, err := pt.Client.PutBulkJobSpectraS3(bulkPut)
    //todo handle what is recoverable
    if err != nil {
        return err
    }

    // init queue, producer and consumer
    var wg sync.WaitGroup

    queue := newOperationQueue(MaxQueueSize) //todo make composable
    producer := newPutProducer(&bulkPutResponse.MasterObjectList, pt.WriteObjects, &queue, pt.Strategy, pt.Client, &wg)
    consumer := newPutConsumer(&queue, &wg, pt.Strategy.BlobStrategy.maxTransferGoroutines())

    // Wait for completion of producer-consumer goroutines
    wg.Add(2) // adding producer and consumer goroutines to wait group
    go producer.run() // producer will add to wg for every blob added to queue, and each transfer performed will decrement from wg
    go consumer.run()
    wg.Wait()

    return nil
}

func (pt *putTransfernator) cancel() {
    //todo
    panic("Not yet implemented")
}