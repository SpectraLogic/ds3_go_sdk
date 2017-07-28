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
    "strconv"
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
    testutils.DeleteBucketContents(client, testBucket)
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

func TestGetBucketPagination(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    var ds3Objects []models.Ds3Object
    for i := 0; i < 15; i++ {
        name := "file" + strconv.Itoa(i + 10) // Start at 10 to avoid alphabetical reordering 0, 1, 10, 11, etc
        curObj := models.Ds3Object{Name:name, Size:0}
        ds3Objects = append(ds3Objects, curObj)
    }

    _, putBulkErr := client.PutBulkJobSpectraS3(models.NewPutBulkJobSpectraS3Request(testBucket, ds3Objects))
    ds3Testing.AssertNilError(t, putBulkErr)

    //Test files indexed 0-4
    result1, err := client.GetBucket(models.NewGetBucketRequest(testBucket).WithMaxKeys(5))
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertInt(t, "Number of Objects", 5, len(result1.ListBucketResult.Objects))
    if result1.ListBucketResult.NextMarker == nil {
        t.Fatal("Expected NextMarker to be non-nil value.")
    }
    for i, obj := range result1.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "ObjectName", ds3Objects[i].Name, obj.Key)
    }

    // Test files indexed 5-9
    result2, err := client.GetBucket(models.NewGetBucketRequest(testBucket).WithMaxKeys(5).WithMarker(result1.ListBucketResult.NextMarker))
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertInt(t, "Number of Objects", 5, len(result2.ListBucketResult.Objects))
    if result2.ListBucketResult.NextMarker == nil {
        t.Fatal("Expected NextMarker to be non-nil value.")
    }
    for i, obj := range result2.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "ObjectName", ds3Objects[i + 5].Name, obj.Key)
    }

    // Test files indexed 10-14
    result3, err := client.GetBucket(models.NewGetBucketRequest(testBucket).WithMaxKeys(5).WithMarker(result2.ListBucketResult.NextMarker))
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertInt(t, "Number of Objects", 5, len(result3.ListBucketResult.Objects))
    if result2.ListBucketResult.NextMarker == nil {
        t.Fatal("Expected NextMarker to be non-nil value.")
    }
    for i, obj := range result3.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "ObjectName", ds3Objects[i + 10].Name, obj.Key)
    }
}

func TestGetBucketDelimiter(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    var ds3Objects []models.Ds3Object
    for i := 0; i < 10; i++ {
        name := "file" + strconv.Itoa(i)
        curObj := models.Ds3Object{Name:name, Size:0}
        ds3Objects = append(ds3Objects, curObj)
    }
    for i := 0; i < 10; i++ {
        name := "dir/file" + strconv.Itoa(i)
        curObj := models.Ds3Object{Name:name, Size:0}
        ds3Objects = append(ds3Objects, curObj)
    }

    _, putBulkErr := client.PutBulkJobSpectraS3(models.NewPutBulkJobSpectraS3Request(testBucket, ds3Objects))
    ds3Testing.AssertNilError(t, putBulkErr)

    //Test files indexed 0-4
    delimiter := "/"
    result, err := client.GetBucket(models.NewGetBucketRequest(testBucket).WithDelimiter(&delimiter))
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertInt(t, "Number of Objects", 10, len(result.ListBucketResult.Objects))
    ds3Testing.AssertInt(t, "Number of Common Prefixes", 1, len(result.ListBucketResult.CommonPrefixes))
    ds3Testing.AssertString(t, "CommonPrefix", "dir/", result.ListBucketResult.CommonPrefixes[0])
}

func TestBulkPut(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    books, bookErr := testutils.GetResourceBooks()
    ds3Testing.AssertNilError(t, bookErr)

    ds3Objects, objErr := testutils.ConvertBooksToDs3Objects(books)
    ds3Testing.AssertNilError(t, objErr)

    // Create bulk put job
    bulkPut, bulkPutErr := client.PutBulkJobSpectraS3(models.NewPutBulkJobSpectraS3Request(testBucket, ds3Objects))
    ds3Testing.AssertNilError(t, bulkPutErr)

    for _, chunk := range bulkPut.MasterObjectList.Objects {
        allocateChunk, allocateErr := client.AllocateJobChunkSpectraS3(models.NewAllocateJobChunkSpectraS3Request(chunk.ChunkId))
        ds3Testing.AssertNilError(t, allocateErr)
        for _, obj := range allocateChunk.Objects.Objects {
            content := books[*obj.Name]
            _, putObjErr := client.PutObject(models.NewPutObjectRequest(testBucket, *obj.Name, content).
                WithOffset(obj.Offset).
                WithJob(bulkPut.MasterObjectList.JobId))

            ds3Testing.AssertNilError(t, putObjErr)
        }
    }

    // Verify books are in the bucket
    getBucket, getBucketErr := client.GetBucket(models.NewGetBucketRequest(testBucket))
    ds3Testing.AssertNilError(t, getBucketErr)
    if len(getBucket.ListBucketResult.Objects) != len(books) {
        t.Fatalf("Expected '%d' objects in bucket '%s', but found '%d'.", len(books), testBucket, len(getBucket.ListBucketResult.Objects))
    }
    for i, obj := range getBucket.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "BookName", testutils.BookTitles[i], obj.Key)
    }
}
