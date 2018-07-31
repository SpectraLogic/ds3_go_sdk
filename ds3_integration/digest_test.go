// Copyright 2014-2018 Spectra Logic Corporation. All Rights Reserved.
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
    "github.com/SpectraLogic/ds3_go_sdk/ds3_integration/utils"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_utils/ds3Testing"
    "bytes"
    "io/ioutil"
)

func testPutBeowulfWithSpecifiedObjectName(objectName string, t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)
    fileName := "beowulf.txt"

    //Put object to BP
    book, err := testutils.LoadBookLogError(t, fileName)
    ds3Testing.AssertNilError(t, err)

    err = testutils.PutObjectLogError(t, client, testBucket, objectName, book)
    ds3Testing.AssertNilError(t, err)
    defer testutils.DeleteObjectLogError(t, client, testBucket, objectName)

    //Verify that object exists
    getObjectResponse, err := testutils.GetObjectLogError(t, client, testBucket, objectName)
    if err != nil {
        defer getObjectResponse.Content.Close()
        bs, err := ioutil.ReadAll(getObjectResponse.Content)
        ds3Testing.AssertNilError(t, err)
        if bytes.Compare(bs, book) != 0 {
            t.Error("Retrieved book does not match uploaded book.")
        }
    }
}

func TestObjectWithSpaces(t *testing.T) {
    objectName := "b e o w u l f"
    testPutBeowulfWithSpecifiedObjectName(objectName, t)
}

func TestObjectWithSlashes(t *testing.T) {
    objectName := "be/owu/lf"
    testPutBeowulfWithSpecifiedObjectName(objectName, t)
}

func TestObjectWithPlus(t *testing.T) {
    objectName := "beo+wulf"
    testPutBeowulfWithSpecifiedObjectName(objectName, t)
}

func TestObjectWithUtf8Chars(t *testing.T) {
    objectName := "aÀˠϠૐ༺ᎀ"
    testPutBeowulfWithSpecifiedObjectName(objectName, t)
}