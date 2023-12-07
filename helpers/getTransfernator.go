package helpers

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "net/http"
    "sync"
)

type getTransceiver struct {
    BucketName string
    ReadObjects *[]helperModels.GetObject
    Strategy *ReadTransferStrategy
    Client *ds3.Client
}

func newGetTransceiver(bucketName string, readObjects *[]helperModels.GetObject, strategy *ReadTransferStrategy, client *ds3.Client) *getTransceiver {
    return &getTransceiver{
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
    if options.Priority != ds3Models.UNDEFINED {
        bulkGet.WithPriority(options.Priority)
    }
    if options.Name != nil {
        bulkGet.WithName(*options.Name)
    }

    return bulkGet
}

// Converts a GetObject into its corresponding Ds3GetObjects for use in bulk get request building.
func createPartialGetObjects(getObject helperModels.GetObject) []ds3Models.Ds3GetObject {
    // handle getting the entire object
    if len(getObject.Ranges) == 0 {
        return []ds3Models.Ds3GetObject { { Name:getObject.Name } }
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

func (transceiver *getTransceiver) createBulkGetJob() (*ds3Models.GetBulkJobSpectraS3Response, *[]helperModels.GetObject, error) {
    // attempt to create a bulk get of all objects
    bulkGet := newBulkGetRequest(transceiver.BucketName, transceiver.ReadObjects, transceiver.Strategy.Options)
    bulkGetResponse, err := transceiver.Client.GetBulkJobSpectraS3(bulkGet)
    if err == nil {
        return bulkGetResponse, transceiver.ReadObjects, nil
    }

    badStatusErr, ok := err.(*ds3Models.BadStatusCodeError)
    if !ok || badStatusErr.ActualStatusCode != http.StatusNotFound{
        // failed to create bulk get for reason other than 404
        return nil, nil, err
    }

    // head each item and try again for all objects that exist in the bucket
    var objectsThatExist []helperModels.GetObject
    for _, obj := range *transceiver.ReadObjects {
        _, err := transceiver.Client.HeadObject(ds3Models.NewHeadObjectRequest(transceiver.BucketName, obj.Name))
        if err != nil {
            // mark file as having a fatal error
            readableErr := fmt.Errorf("failed HeadObject call on %s: %v", obj.Name, err)
            obj.ChannelBuilder.SetFatalError(readableErr)
            transceiver.Strategy.Listeners.Errored(obj.Name, readableErr)
        } else {
            objectsThatExist = append(objectsThatExist, obj)
        }
    }
    if len(objectsThatExist) <= 0 {
        return nil, nil, err
    }

    // create bulk get job for all objects that exist in the bucket
    bulkGet = newBulkGetRequest(transceiver.BucketName, &objectsThatExist, transceiver.Strategy.Options)
    bulkGetResponse, err = transceiver.Client.GetBulkJobSpectraS3(bulkGet)
    if err != nil {
        return nil, nil, err
    } else {
        return bulkGetResponse, &objectsThatExist, nil
    }
}

func (transceiver *getTransceiver) transfer() (string, error) {
    // create bulk get job
    bulkGetResponse, objectsToRetrieve, err := transceiver.createBulkGetJob()
    if err != nil {
        return "", err
    }

    // init queue, producer and consumer
    var waitGroup sync.WaitGroup

    doneNotifier := NewConditionalBool()

    queue := newOperationQueue(transceiver.Strategy.BlobStrategy.maxWaitingTransfers(), transceiver.Client.Logger)
    producer := newGetProducer(&bulkGetResponse.MasterObjectList, objectsToRetrieve, &queue, transceiver.Strategy, transceiver.Client, &waitGroup, doneNotifier)
    consumer := newConsumer(&queue, &waitGroup, transceiver.Strategy.BlobStrategy.maxConcurrentTransfers(), doneNotifier)

    // Wait for completion of producer-consumer goroutines
    waitGroup.Add(1)  // adding producer and consumer goroutines to wait group
    go consumer.run()
    err = producer.run() // producer will add to waitGroup for every blob retrieval added to queue, and each transfer performed will decrement from waitGroup
    waitGroup.Wait()

    return bulkGetResponse.MasterObjectList.JobId, err
}
