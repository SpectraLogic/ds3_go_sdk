package testutils

import (
    "testing"
    "ds3/models"
    "ds3"
    "io/ioutil"
    "ds3/networking"
    "errors"
    "ds3/buildclient"
    "log"
    "io"
    "ds3_utils/ds3Testing"
    "bytes"
)

var bookPath = "./resources/books/"
var nilResponse error = errors.New("Received an unexpected nil response.")
var BookTitles = []string{ "beowulf.txt", "sherlock_holmes.txt", "tale_of_two_cities.txt", "ulysses.txt"}

func VerifyBookContent(t *testing.T, bookName string, content io.ReadCloser) {
    actual, loadErr := LoadBook(bookName)
    ds3Testing.AssertNilError(t, loadErr)

    bs, readErr := ioutil.ReadAll(content)
    ds3Testing.AssertNilError(t, readErr)
    if bytes.Compare(bs, actual) != 0 {
        t.Error("Retrieved book does not match uploaded book.")
    }
}

func ConvertObjectsIntoObjectNameList(objects []models.Contents) []string {
    var objectNames []string
    for _, obj := range objects{
        objectNames = append(objectNames, *obj.Key)
    }
    return objectNames
}

//Puts the test books onto the BP
func PutTestBooks(client *ds3.Client, bucketName string) error {
    books, bookErr := GetResourceBooks()
    if bookErr != nil {
        return bookErr
    }

    ds3Objects, objErr := ConvertBooksToDs3Objects(books)
    if objErr != nil {
        return objErr
    }

    // Create bulk put job
    bulkPut, bulkPutErr := client.PutBulkJobSpectraS3(models.NewPutBulkJobSpectraS3Request(bucketName, ds3Objects))
    if bulkPutErr != nil {
        return bulkPutErr
    }

    for _, chunk := range bulkPut.MasterObjectList.Objects {
        allocateChunk, allocateErr := client.AllocateJobChunkSpectraS3(models.NewAllocateJobChunkSpectraS3Request(chunk.ChunkId))
        if allocateErr != nil {
            return allocateErr
        }
        for _, obj := range allocateChunk.Objects.Objects {
            content := books[*obj.Name]
            _, putObjErr := client.PutObject(models.NewPutObjectRequest(bucketName, *obj.Name, content).
                WithOffset(obj.Offset).
                WithJob(bulkPut.MasterObjectList.JobId))

            if putObjErr != nil {
                return putObjErr
            }
        }
    }
    return nil
}

func DeleteBucketContents(client *ds3.Client, bucketName string) {
    //Get the contents of bucket
    getBucket, getBucketErr := client.GetBucket(models.NewGetBucketRequest(bucketName))
    if getBucketErr != nil {
        log.Printf("WARNING: Cleanup error when attempting to get contents of bucket '%s': '%s'.", bucketName, getBucketErr.Error())
        return
    }
    for _, obj := range getBucket.ListBucketResult.Objects {
        deleteErr := DeleteObject(client, bucketName, *obj.Key)
        if deleteErr != nil {
            log.Printf("WARNING: Cleanup error when attempting to delete object '%s' from bucket '%s': '%s'.", *obj.Key, bucketName, deleteErr.Error())
        }
    }
}

func GetResourceBooks() (map[string]networking.ReaderWithSizeDecorator, error) {
    result := make(map[string]networking.ReaderWithSizeDecorator)
    for _, title := range BookTitles {
        curStream, err := GetResourceBookAsStream(title)
        if err != nil {
            return nil, err
        }
        result[title] = *curStream
    }
    return result, nil
}

func ConvertBooksToDs3Objects(books map[string]networking.ReaderWithSizeDecorator) ([]models.Ds3PutObject, error) {
    var ds3Objects []models.Ds3PutObject
    for title, stream := range books {
        size, err := stream.Size()
        if err != nil {
            return nil, err
        }
        var curObj models.Ds3PutObject = models.Ds3PutObject{
            Name:title,
            Size:size,
        }
        ds3Objects = append(ds3Objects, curObj)
    }
    return ds3Objects, nil
}

func GetResourceBookAsStream(book string) (*networking.ReaderWithSizeDecorator, error) {
    content, err := LoadBook(book)
    if err != nil {
        return nil, err
    }
    stream := networking.BuildByteReaderWithSizeDecorator(content)
    return &stream, nil
}

// Loads the specified book. If an error occurs, it is logged, and the calling
// test is marked as failed.
func LoadBookLogError(t *testing.T, book string) ([]byte, error) {
    data, err := LoadBook(book)
    if err != nil {
        t.Errorf("Unexpected error when loading test file: %s", err.Error())
        return nil, err
    }
    return data, nil
}

// Loads the specified book. Assumes that the book is located in the /resources/books/ folder.
func LoadBook(book string) ([]byte, error) {
    return ioutil.ReadFile(bookPath + book)
}

// Puts the specified object. If an error occurs, it is logged, and the calling
// test is marked as failed.
func PutObjectLogError(t *testing.T, client *ds3.Client, bucketName string, objectName string, data []byte) (error) {
    err := PutObject(client, bucketName, objectName, data)
    if err != nil {
        t.Error(err)
    }
    return nil
}

// Puts the specified object. Returns an error if not successful.
func PutObject(client *ds3.Client, bucketName string, objectName string, data []byte) (error) {
    putObjectResponse, putErr := client.PutObject(models.NewPutObjectRequest(bucketName, objectName, networking.BuildByteReaderWithSizeDecorator(data)))
    if putErr != nil {
        return putErr
    }
    if putObjectResponse == nil {
        return nilResponse
    }
    return nil
}

// Deletes the specified object. If an error occurs, it is logged, and the calling
// test is marked as failed.
func DeleteObjectLogError(t *testing.T, client *ds3.Client, bucketName string, objectName string) (error) {
    err := DeleteObject(client, bucketName, objectName)
    if err != nil {
        t.Error(err)
    }
    return err
}

// Deletes the specified object. Returns an error if not successful.
func DeleteObject(client *ds3.Client, bucketName string, objectName string) (error) {
    deleteObjectResponse, deleteErr := client.DeleteObject(models.NewDeleteObjectRequest(bucketName, objectName))
    if deleteErr != nil {
        return deleteErr
    }
    if deleteObjectResponse == nil {
        return nilResponse
    }
    return nil
}

// Retrieves the specified object. If an error occurs, it is logged, and the calling
// test is marked as failed.
func GetObjectLogError(t *testing.T, client *ds3.Client, bucketName string, objectName string) (*models.GetObjectResponse, error) {
    response, err := GetObject(client, bucketName, objectName)
    if err != nil {
        t.Error(err.Error())
        return nil, err
    }
    return response, nil
}

// Retrieves the specified object. Returns an error if not successful.
func GetObject(client *ds3.Client, bucketName string, objectName string) (*models.GetObjectResponse, error) {
    response, err := client.GetObject(models.NewGetObjectRequest(bucketName, objectName))
    if err != nil {
        return nil, err
    } else if response == nil {
        return nil, nilResponse
    }
    return response, nil
}

// Puts the specified bucket. If an error occurs, it is logged, and the calling
// test is marked as failed.
func PutBucketLogError(t *testing.T, client *ds3.Client, bucketName string) (error) {
    err := PutBucket(client, bucketName)
    if err != nil {
        t.Error(err)
    }
    return err
}

// Puts the specified bucket. Returns an error if not successful.
func PutBucket(client *ds3.Client, bucketName string) (error) {
    putBucketResponse, putErr := client.PutBucket(models.NewPutBucketRequest(bucketName))
    if putErr != nil {
        return putErr
    }
    if putBucketResponse == nil {
        return nilResponse
    }
    return nil
}

// Deletes the specified bucket. If an error occurs, it is logged, and the calling
// test is marked as failed.
func DeleteBucketLogError(t *testing.T, client *ds3.Client, bucketName string) (error) {
    err := DeleteBucket(client, bucketName)
    if err != nil {
        t.Error(err)
    }
    return err
}

// Deletes the specified bucket and returns an error if one occurs
func DeleteBucket(client *ds3.Client, bucketName string) (error) {
    deleteBucket, deleteErr := client.DeleteBucket(models.NewDeleteBucketRequest(bucketName))
    if deleteErr != nil {
        return deleteErr
    }
    if deleteBucket == nil {
        return nilResponse
    }
    return nil
}

// Retrieves the specified bucket. If an error occurs, it is logged, and the calling
// test is marked as failed.
func GetBucketLogError(t *testing.T, client *ds3.Client, bucketName string) (*models.GetBucketResponse, error) {
    response, err := GetBucket(client, bucketName)
    if err != nil {
        t.Error(err)
        return nil, nilResponse
    }
    return response, nil
}

// Retrieves the specified bucket. Returns an error if unsuccessful.
func GetBucket(client *ds3.Client, bucketName string) (*models.GetBucketResponse, error) {
    response, err := client.GetBucket(models.NewGetBucketRequest(bucketName))
    if err != nil {
        return nil, err
    } else if response == nil {
        return nil, nilResponse
    }
    return response, nil
}

// Sets up the test environment by creating a client from environment
// variables and creating a test bucket
func SetupTestEnv(test_bucket string) (*ds3.Client, error) {
    // Build the client from environment variables
    client, clientErr := buildclient.FromEnv()
    if clientErr != nil {
        return nil, clientErr
    }

    // Create the test bucket
    putBucketErr := PutBucket(client, test_bucket)
    if putBucketErr != nil {
        return nil, putBucketErr
    }

    return client, nil
}

// Tears down the test environment by deleting all contents in the
// test bucket
func TeardownTestEnv(client *ds3.Client, testBucket string) {
    //Delete contents in test bucket
    DeleteBucketContents(client, testBucket)

    //Delete the test bucket
    err := DeleteBucket(client, testBucket)
    if err != nil {
        log.Print(err)
    }
}
