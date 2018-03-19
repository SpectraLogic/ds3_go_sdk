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

package functions

import (
    "log"
    "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/ds3/buildclient"
    "spectra/ds3_go_sdk/samples/utils"
    "fmt"
)

// Demonstrates how to get an object within a bucket. Assumes that the
// target bucket already exists and has the specified file.
// i.e. run putBulkSample.go first.
func GetObjectSample() {
    fmt.Println("---- Get Object Sample ----")

    // Create the client from environment variables.
    client, err := buildclient.FromEnv()
    if err != nil {
        log.Fatal(err)
    }

    // Grab the first book defined in utils.BookNames
    getObjRequest := models.NewGetObjectRequest(utils.BucketName, utils.BookNames[0])
    getObjResponse, err := client.GetObject(getObjRequest)
    if err != nil {
        log.Fatal(err)
    }

    // Verify that the contents of the book were retrieved correctly.
    utils.VerifyBookContent(utils.BookNames[0], getObjResponse.Content)
    fmt.Printf("Retrieved: %s\n", utils.BookNames[0])
}
