package helpers

import (
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "sync"
)

type putTransceiver struct {
    BucketName string
    WriteObjects *[]helperModels.PutObject
    Strategy *WriteTransferStrategy
    Client *ds3.Client
}

func newPutTransceiver(bucketName string, writeObjects *[]helperModels.PutObject, strategy *WriteTransferStrategy, client *ds3.Client) *putTransceiver {
    return &putTransceiver{
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
    if options.Protected != nil {
        bulkPut = bulkPut.WithProtected(*options.Protected)
    }
    return bulkPut
}

func (transceiver *putTransceiver) transfer() (string, error) {
    // create bulk put job
    bulkPut := newBulkPutRequest(transceiver.BucketName, transceiver.WriteObjects, transceiver.Strategy.Options)

    bulkPutResponse, err := transceiver.Client.PutBulkJobSpectraS3(bulkPut)
    if err != nil {
        return "", err
    }

    // init queue, producer and consumer
    var waitGroup sync.WaitGroup

    doneNotifier := NewConditionalBool()

    queue := newOperationQueue(transceiver.Strategy.BlobStrategy.maxWaitingTransfers(), transceiver.Client.Logger)
    producer := newPutProducer(&bulkPutResponse.MasterObjectList, transceiver.WriteObjects, &queue, transceiver.Strategy, transceiver.Client, &waitGroup, doneNotifier)
    consumer := newConsumer(&queue, &waitGroup, transceiver.Strategy.BlobStrategy.maxConcurrentTransfers(), doneNotifier)

    // Wait for completion of producer-consumer goroutines
    waitGroup.Add(1)  // adding producer and consumer goroutines to wait group

    go consumer.run()
    err = producer.run() // producer will add to waitGroup for every blob added to queue, and each transfer performed will decrement from waitGroup
    waitGroup.Wait()

    return bulkPutResponse.MasterObjectList.JobId, err
}

/*
func (transfernator *putTransceiver) cancel() {
    panic("Not yet implemented")
}
*/