package ds3_integration

import (
    "testing"
    "log"
    "os"
    "ds3/buildclient"
    "ds3"
    "ds3/models"
    "ds3_integration/utils"
    "io/ioutil"
    "bytes"
)

var client *ds3.Client
var testBucket = "GoSmokeTestBucket"

func setup() (*ds3.Client, error) {
    // Build the client from environment variables
    client, clientErr := buildclient.FromEnv()
    if clientErr != nil {
        return nil, clientErr
    }

    // Create the test bucket
    putBucketErr := testutils.PutBucket(client, testBucket)
    if putBucketErr != nil {
        return nil, putBucketErr
    }

    return client, nil
}

func teardown() {
    //Delete the test bucket
    err := testutils.DeleteBucket(client, testBucket)
    if err != nil {
        log.Print(err)
    }
}

func TestMain(m *testing.M) {
    var err error
    var exitVal int
    client, err = setup()
    if err != nil {
        log.Printf("Unable to setup test environment '%s'.", err.Error())
        exitVal = 1
    } else {
        exitVal = m.Run()
    }
    teardown()
    os.Exit(exitVal)
}

func TestGetService(t *testing.T) {
    response, err := client.GetService(models.NewGetServiceRequest())
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }
    if response == nil {
        t.Fatal("Received an unexpected nil response.")
    }
}

func TestBucket(t *testing.T) {
    //Create bucket
    bucketName := "GoTestPutBucket"
    putErr := testutils.PutBucketLogError(t, client, bucketName)
    if putErr != nil {
        t.FailNow()
    }

    //Verify that bucket exists
    getBucketResponse, getErr := testutils.GetBucketLogError(t, client, bucketName)
    if getErr == nil && *getBucketResponse.ListBucketResult.Name != bucketName {
        t.Errorf("Unexpected bucket name: expected `%s` but got `%s`.", bucketName, getBucketResponse.ListBucketResult.Name)
    }

    //Delete bucket
    testutils.DeleteBucketLogError(t, client, bucketName)
}

func TestObject(t *testing.T) {
    beowulf := "beowulf.txt"

    //Put object to BP
    book, bookErr := testutils.LoadBookLogError(t, beowulf)
    if bookErr != nil {
        t.FailNow()
    }
    putObjErr := testutils.PutObjectLogError(t, client, testBucket, beowulf, book)
    if putObjErr != nil {
        t.FailNow()
    }

    //Verify that object exists
    getObjectResponse, getObjErr := testutils.GetObjectLogError(t, client, testBucket, beowulf)
    if getObjErr != nil {
        defer getObjectResponse.Content.Close()
        bs, readErr := ioutil.ReadAll(getObjectResponse.Content)
        if readErr != nil {
            t.Errorf("Unexpected error '%s'.", readErr.Error())
        } else if bytes.Compare(bs, book) != 0 {
            t.Error("Retrieved book does not match uploaded book.")
        }
    }

    //Delete object
    testutils.DeleteObjectLogError(t, client, testBucket, beowulf)
}
