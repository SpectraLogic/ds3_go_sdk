package helpers

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "github.com/SpectraLogic/ds3_go_sdk/helpers/ranges"
    "net/http"
    "sort"
    "strings"
)

type HelperInterface interface {
    //ListObjectsFromBucket(bucketName string) []ds3Models.S3Object
    //ListObjectsFromDirectory(directoryName string) []helperModels.PutObject

    // Puts the specified list of objects on the Black Pearl in the specified bucket.
    // Returns the Black Pearl bulk put Job ID and any errors that may have occurred.
    // A job ID will be returned if a BP job was successfully created, regardless of
    // whether additional errors occur.
    PutObjects(bucketName string, objects []helperModels.PutObject, strategy WriteTransferStrategy) (string, error)

    // Retrieves the list of objects from the specified bucket on the Black Pearl.
    // Returns the Black Pearl bulk get Job ID and any errors that may have occurred.
    // A job ID will be returned if a BP job was successfully created, regardless of
    // whether additional errors occur.
    GetObjects(bucketName string, objects []helperModels.GetObject, strategy ReadTransferStrategy) (string, error)

    // Retrieves the list of objects from the specified bucket on the Black Pearl.
    // If a get job cannot be created due to insufficient cache space to fulfill an
    // IN_ORDER processing guarantee, then the job is split across multiple BP jobs.
    // This allows for the IN_ORDER retrieval of objects that exceed available cache space.
    GetObjectsSpanningJobs(bucketName string, objects []helperModels.GetObject, strategy ReadTransferStrategy) ([]string, error)
}

type HelperImpl struct {
     client *ds3.Client
}

func NewHelpers(client *ds3.Client) HelperInterface {
    return &HelperImpl{client:client}
}

/*
func (helper *HelperImpl) ListObjectsFromBucket(bucketName string) []ds3Models.S3Object {
    panic("not implemented yet")
}

func (helper *HelperImpl) ListObjectsFromDirectory(directoryName string) []helperModels.PutObject {
    panic("not implemented yet")
}
*/

func (helper *HelperImpl) PutObjects(bucketName string, objects []helperModels.PutObject, strategy WriteTransferStrategy) (string, error) {
    transceiver := newPutTransceiver(bucketName, &objects, &strategy, helper.client)
    return transceiver.transfer()
}

func (helper *HelperImpl) GetObjects(bucketName string, objects []helperModels.GetObject, strategy ReadTransferStrategy) (string, error) {
    transceiver := newGetTransceiver(bucketName, &objects, &strategy, helper.client)
    return transceiver.transfer()
}

func (helper *HelperImpl) GetObjectsSpanningJobs(bucketName string, objects []helperModels.GetObject, strategy ReadTransferStrategy) ([]string, error) {
    // Attempt to send the entire job at once
    jobId, err := helper.GetObjects(bucketName, objects, strategy)
    if err == nil {
        // success
        return []string{jobId}, nil
    } else if !helper.isCannotPreAllocateError(err) {
        // error not related to pre-allocation
        return nil, err
    }

    // Retrieve each file individually
    var jobIds []string
    for _, getObject := range objects {
        fileJobIds := helper.retrieveIndividualFile(bucketName, getObject, strategy)
        jobIds = append(jobIds, fileJobIds...)
    }
    return jobIds, nil
}

func (helper *HelperImpl) isCannotPreAllocateError(err error) bool {
    badStatusErr, ok := err.(*models.BadStatusCodeError)
    if !ok || badStatusErr.ActualStatusCode != http.StatusServiceUnavailable {
        // failed to create bulk get for reason other than 503
        return false
    }

    if strings.Contains(badStatusErr.Error(), "GET jobs that have a chunkClientProcessingOrderGuarantee of IN_ORDER must be entirely pre-allocated.  Cannot pre-allocate") {
        return true
    }
    return false
}

func (helper *HelperImpl) retrieveIndividualFile(bucketName string, getObject helperModels.GetObject, strategy ReadTransferStrategy) []string {
    // Get the blob offsets
    headObject, err := helper.client.HeadObject(models.NewHeadObjectRequest(bucketName, getObject.Name))
    if err != nil {
        getObject.ChannelBuilder.SetFatalError(err)
        return nil
    }
    var offsets []int64
    for offset := range headObject.BlobChecksums {
        offsets = append(offsets, offset)
    }

    sort.Slice(offsets, func(i, j int) bool {
        return offsets[i] < offsets[j]
    })

    // Get the object size
    objectsDetails, err := helper.client.GetObjectsWithFullDetailsSpectraS3(
        models.NewGetObjectsWithFullDetailsSpectraS3Request().
            WithBucketId(bucketName).WithName(getObject.Name).
            WithLatest(true))

    if err != nil {
        getObject.ChannelBuilder.SetFatalError(err)
        return nil
    } else if len(objectsDetails.DetailedS3ObjectList.DetailedS3Objects) < 1 {
        getObject.ChannelBuilder.SetFatalError(fmt.Errorf("failed to get object details"))
        return nil
    }

    // Retrieve the object one blob at a time in order
    objectCopy := getObject
    objectEnd := objectsDetails.DetailedS3ObjectList.DetailedS3Objects[0].Size - 1
    if len(objectCopy.Ranges) == 0 {
        // If the user didn't specify a range, add a range that covers the entire file
        // so that we can use the blobRangeFinder to tell us what ranges to specify.
        objectCopy.Ranges = append(objectCopy.Ranges, models.Range{Start: 0, End: objectEnd})
    }
    blobFinder := ranges.NewBlobRangeFinder(&[]helperModels.GetObject{objectCopy})

    var jobIds []string
    for i, offset := range offsets {
        var blobEnd int64
        if i+1 < len(offsets) {
            blobEnd = offsets[i+1]-1
        } else {
            blobEnd = objectEnd
        }
        length := blobEnd - offset + 1
        blobRanges := blobFinder.GetRanges(objectCopy.Name, offset, length)
        if len(blobRanges) == 0 {
            // This blob does not need to be retrieved
            continue
        }

        jobId, err := helper.retrieveBlob(bucketName, getObject, blobRanges, strategy)
        if err != nil {
            getObject.ChannelBuilder.SetFatalError(err)
            return nil
        }
        jobIds = append(jobIds, jobId)

        if objectCopy.ChannelBuilder.HasFatalError() {
            // Failed to retrieve a portion of the file, don't bother with the rest
            break
        }
    }
    return jobIds
}

func (helper *HelperImpl) retrieveBlob(bucketName string, getObject helperModels.GetObject, blobRanges []models.Range, strategy ReadTransferStrategy) (string, error) {
    // Since there is only one blob being retrieved, create the job with Order-Guarantee=None so that the
    // job will wait if cache needs to be reclaimed on the BP before the chunk can be allocated.
    getObjectBlob := getObject
    getObjectBlob.Ranges = blobRanges

    strategyCopy := strategy
    strategyCopy.Options.ChunkClientProcessingOrderGuarantee = models.JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_NONE

    return helper.GetObjects(bucketName, []helperModels.GetObject{getObjectBlob}, strategyCopy)
}
