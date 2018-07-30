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

package functions

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/buildclient"
    "log"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

// Demonstrates how to create a bucket in two ways.
// 1) how to create a bucket using the default data policy
// 2) how to create a bucket by specifying a specific data policy
//
// This assumes that the BP credentials are stored in environment variables.
func PutBucketSample() {
    fmt.Println("---- Put Bucket Sample ----")
    client, err := buildclient.FromEnv()
    if err != nil {
        log.Fatal(err)
    }

    // 1) Create a bucket using the default data policy. Note that there must already exist a
    //    default data policy on the specified BP for this to work.

    bucket1Name := "bucket1Sample"

    // Create the put bucket request
    bucket1Request := models.NewPutBucketRequest(bucket1Name)

    // Perform the put bucket call by using the client and invoking the command
    _, err = client.PutBucket(bucket1Request)
    if err != nil {
        log.Fatal(err)
    }

    // 2) Create a bucket while specifying the specific data policy to be assigned to the bucket
    // This assumes that the data policy "Dual Copy on Tape" already exists on the BP

    dataPolicyName := "Dual Copy on Tape"
    bucket2Name := "bucket2Sample"

    // Create the spectra S3 put bucket request and specify the desired data policy. Note that
    // you can use either the data policy's name or its UUID
    bucket2Request := models.NewPutBucketSpectraS3Request(bucket2Name).WithDataPolicyId(dataPolicyName)

    // Perform the spectra S3 put bucket call by using the client and invoking the command
    _, err = client.PutBucketSpectraS3(bucket2Request)
    if err != nil {
        log.Fatal(err)
    }
}
