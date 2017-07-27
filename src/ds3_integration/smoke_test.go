// Copyright 2014-2017 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package ds3_integration

import (
    "testing"
    "log"
    "os"
    "ds3"
    "ds3/models"
    "ds3_integration/utils"
    "io/ioutil"
    "bytes"
    "ds3_utils/ds3Testing"
)

var client *ds3.Client
var testBucket = "GoIntegrationTestBucket"

func TestMain(m *testing.M) {
    var err error
    var exitVal int
    client, err = testutils.SetupTestEnv(testBucket)
    if err != nil {
        log.Printf("Unable to setup test environment '%s'.", err.Error())
        exitVal = 1
    } else {
        exitVal = m.Run()
    }
    testutils.TeardownTestEnv(client, testBucket)
    os.Exit(exitVal)
}

func TestGetService(t *testing.T) {
    response, err := client.GetService(models.NewGetServiceRequest())
    ds3Testing.AssertNilError(t, err)
    if response == nil {
        t.Fatal("Received an unexpected nil response.")
    }
}

func TestBucket(t *testing.T) {
    //Create bucket
    bucketName := "GoTestPutBucket"
    putErr := testutils.PutBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, putErr)
    defer testutils.DeleteBucketLogError(t, client, bucketName)

    //Verify that bucket exists
    getBucketResponse, getErr := testutils.GetBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, getErr)
    ds3Testing.AssertNonNilStringPtr(t, "Name", bucketName, getBucketResponse.ListBucketResult.Name)
}

func TestObject(t *testing.T) {
    beowulf := "beowulf.txt"

    //Put object to BP
    book, bookErr := testutils.LoadBookLogError(t, beowulf)
    ds3Testing.AssertNilError(t, bookErr)
    defer testutils.DeleteObjectLogError(t, client, testBucket, beowulf)

    putObjErr := testutils.PutObjectLogError(t, client, testBucket, beowulf, book)
    ds3Testing.AssertNilError(t, putObjErr)

    //Verify that object exists
    getObjectResponse, getObjErr := testutils.GetObjectLogError(t, client, testBucket, beowulf)
    if getObjErr != nil {
        defer getObjectResponse.Content.Close()
        bs, readErr := ioutil.ReadAll(getObjectResponse.Content)
        ds3Testing.AssertNilError(t, readErr)
        if bytes.Compare(bs, book) != 0 {
            t.Error("Retrieved book does not match uploaded book.")
        }
    }
}

func TestBadBucket(t *testing.T) {
    // Attempt to create a malformed bucket
    badBucketName := ""
    err := testutils.PutBucket(client, badBucketName)
    ds3Testing.AssertBadStatusCodeError(t, 400, err)
}

func TestBucketAlreadyExists(t *testing.T) {
    // Attempt to create a bucket that already exists
    err := testutils.PutBucket(client, testBucket)
    ds3Testing.AssertBadStatusCodeError(t, 409, err)
}

func TestDeleteEmptyBucket(t *testing.T) {
    bucketName := "GoDeleteEmptyBucket"
    putErr := testutils.PutBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, putErr)

    //Delete bucket
    deleteErr := testutils.DeleteBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, deleteErr)

    //Verify bucket does not exist
    _, getDeletedErr := testutils.GetBucket(client, bucketName)
    ds3Testing.AssertBadStatusCodeError(t, 404, getDeletedErr)
}

func TestDeleteBucketMalformedName(t *testing.T) {
    err := testutils.DeleteBucket(client, "")
    ds3Testing.AssertBadStatusCodeError(t, 400, err)
}

func TestDeleteBucketNonexistent(t *testing.T) {
    err := testutils.DeleteBucket(client, "not-here")
    ds3Testing.AssertBadStatusCodeError(t, 404, err)
}

func TestDeleteBucketNonEmpty(t *testing.T) {
    beowulf := "beowulf.txt"
    bucketName := "GoTestDeleteBucketNonEmpty"

    //Create test bucket
    err := testutils.PutBucket(client, bucketName)
    ds3Testing.AssertNilError(t, err)
    defer testutils.DeleteBucketLogError(t, client, bucketName)

    //Put object to test bucket
    book, bookErr := testutils.LoadBookLogError(t, beowulf)
    ds3Testing.AssertNilError(t, bookErr)
    defer testutils.DeleteObjectLogError(t, client, bucketName, beowulf)

    putObjErr := testutils.PutObjectLogError(t, client, bucketName, beowulf, book)
    ds3Testing.AssertNilError(t, putObjErr)

    //Attempt to delete non-empty bucket
    deleteErr := testutils.DeleteBucket(client, bucketName)
    ds3Testing.AssertBadStatusCodeError(t, 409, deleteErr)
}

func TestGetBucket(t *testing.T) {
    response, err := testutils.GetBucket(client, testBucket)
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertNonNilStringPtr(t, "Name", testBucket, response.ListBucketResult.Name)
}

func TestGetBucketNonexistent(t *testing.T) {
    badBucketName := "not-there"
    _, err := testutils.GetBucket(client, badBucketName)
    ds3Testing.AssertBadStatusCodeError(t, 404, err)
}

func TestHeadBucket(t *testing.T) {
    _, err := client.HeadBucket(models.NewHeadBucketRequest(testBucket))
    ds3Testing.AssertNilError(t, err)
}

func TestHeadBucketNonExistent(t *testing.T) {
    bucketName := "not-here"
    _, err := client.HeadBucket(models.NewHeadBucketRequest(bucketName))
    ds3Testing.AssertBadStatusCodeError(t, 404, err)
}
