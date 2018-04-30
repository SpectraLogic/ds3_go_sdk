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
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/buildclient"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/samples/utils"
)

// Demonstrates how to get a list of S3 objects in a bucket. Assumes that the target
// bucket already exists and has files, i.e. run putBulkSample.go first.
func GetBucketSample() {
    fmt.Println("---- Get Bucket Sample ----")

    // Create the client from environment variables.
    client, err := buildclient.FromEnv()
    if err != nil {
        log.Fatal(err)
    }

    // Create the get bucket request.
    bucketRequest := models.NewGetBucketRequest(utils.BucketName)

    // Perform the Get Bucket call by using the client and invoking the desired command.
    bucketResponse, err := client.GetBucket(bucketRequest)
    if err != nil {
        log.Fatal(err)
    }

    if len(bucketResponse.ListBucketResult.Objects) == 0 {
        fmt.Printf("There are no objects in bucket '%s'.\n", utils.BucketName)
        return
    }

    fmt.Printf("Objects in bucket '%s'\n", utils.BucketName)
    for i, obj := range bucketResponse.ListBucketResult.Objects {
        fmt.Printf("%d) Object: '%s' created on '%s'\n",
            i,
            utils.ToSafeString(obj.Key),
            utils.ToSafeString(obj.LastModified))
    }
}
