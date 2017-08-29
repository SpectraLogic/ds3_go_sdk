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

package main

import (
    "log"
    "ds3/models"
    "ds3/buildclient"
    "fmt"
    "samples/utils"
)

func main() {
    // Create the client from environment variables.
    client, err := buildclient.FromEnv()
    if err != nil {
        log.Fatal(err)
    }

    // Create the get service request. All requests to a DS3 appliance start this way.
    request := models.NewGetServiceRequest()

    // Perform the Get Service call by using the client and invoking the desired command.
    response, err := client.GetService(request)
    if err != nil {
        log.Fatal(err)
    }

    // Printing contents of get service.
    fmt.Println("Buckets:")
    for i, bucket := range response.ListAllMyBucketsResult.Buckets {
        fmt.Printf("%d) Name: %s; CreationDate: %s\n", i + 1, utils.ToSafeString(bucket.Name), utils.ToSafeString(bucket.CreationDate))
    }
}
