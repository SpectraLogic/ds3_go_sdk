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

package main

import (
    "encoding/json"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/samples/functions"
    "github.com/SpectraLogic/ds3_go_sdk/samples/utils"
)

const (
    fileSize = 1024 * 1024
    filesPerThread = 10
    numThreads = 3
    jobName = "PerformanceTest"
)

// Runs the various sample code. Each sample is self contained and can be run separately.
func main() {
    // opject returned / printed tpo console
    output := utils.PerformanceOutput{}
    bucketName := "bucket_" + jobName

    // Put Random binary data
    putStats, err := functions.PerormancePutSample( numThreads, filesPerThread, fileSize, bucketName)
    if err != nil {
        output.AddError(fmt.Sprintf("Put operation failed", err))
    } else {
        putStats.Name = jobName + "_put"
        output.Put = putStats
    }

    // Retrieve same data to test Get.
    getStats, err := functions.PerformanceGetSample(numThreads, filesPerThread, fileSize, bucketName)
    if err != nil {
        output.AddError(fmt.Sprintf("Get operation failed", err))
    } else {
        getStats.Name = jobName + "_get"
        getStats.NumFiles = filesPerThread
        getStats.FileSize = fileSize
        output.Get = getStats
    }

    // export pretty JSON
    jsonOut, err := json.MarshalIndent(output, "", " ")
    if err != nil {
        jsonOut = []byte(fmt.Sprintf("{errors: [\"Could not marshall JSON object\"]}"))
    }
    fmt.Printf("%s\n", jsonOut)
}