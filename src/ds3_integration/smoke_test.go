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
    "ds3/buildclient"
    "ds3"
    "ds3/models"
    "ds3_integration/utils"
    "io/ioutil"
    "bytes"
    "ds3_utils/ds3Testing"
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
