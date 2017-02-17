package testutils

import (
    "testing"
    "ds3/models"
    "ds3"
    "io/ioutil"
    "ds3/network"
    "errors"
)

var bookPath = "./resources/books/"
var nilResponse error = errors.New("Received an unexpected nil response.")

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
    putObjectResponse, putErr := client.PutObject(models.NewPutObjectRequest(bucketName, objectName, net.BuildSizedReadCloser(data)))
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