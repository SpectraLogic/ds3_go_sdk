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

package commands

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_cli/commands/performance"
)

func performanceTest(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify bucket when doing a performance_test.")
    }
    if args.FileSize == 0 {
        return errors.New("Must specify file_size (MB) when doing a performance_test.")
    }
    if args.NumThreads == 0 {
        return errors.New("Must specify num_threads when doing a performance_test.")
    }
    if args.NumFiles == 0 {
        return errors.New("Must specify num_files when doing a performance_test.")
    }

    // set constants from perf args
    fileSize := int64(args.FileSize * 1024 * 1024)
    filesPerThread := args.NumFiles
    numThreads := args.NumThreads
    jobName := args.Bucket
    bucketName := "bucket_" + jobName

    // data container
    output := performance.PerformanceOutput{}

    // do a get service just to test connectivity
    _, err := client.GetService(models.NewGetServiceRequest())
    if err != nil {
        output.AddError(fmt.Sprintf("Could not connect to target", err))
    }

    // Put Random binary data
    putStats, err := performance.PerormancePutSample(client, numThreads, filesPerThread, fileSize, bucketName)
    if err != nil {
        output.AddError(fmt.Sprintf("Put operation failed", err))
    } else {
        putStats.Name = jobName + "_put"
        output.Put = putStats
    }

    // Retrieve same data to test Get.
    getStats, err := performance.PerformanceGetSample(client, numThreads, filesPerThread, fileSize, bucketName)
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
    return nil
}