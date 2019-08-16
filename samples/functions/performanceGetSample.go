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
    "github.com/SpectraLogic/ds3_go_sdk/ds3/buildclient"
    "log"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "github.com/SpectraLogic/ds3_go_sdk/samples/utils"
    "time"
    "fmt"
)

// Demonstrates how to perform a bulk get. Assumes that the target bucket already exists
// and has files, i.e. run putBulkSample.go first.
func PerformanceGetSample(bucketName string) (*utils.PerformanceTest, error) {
    fmt.Println("---- Get Bulk Sample ----")
    ret := utils.NewPerformanceTest(0,0)

    // Create a client from environment variables.
    client, err := buildclient.FromEnv()
    if err != nil {
        return nil, err
    }

    // Get a list of all objects in the target bucket.
    bucketResponse, err := client.GetBucket(models.NewGetBucketRequest(bucketName))
    if err != nil {
        return nil, err
    }

    objectList := bucketResponse.ListBucketResult.Objects

    // Convert object list returned from get bucket into a list of object names.
    var objectNames []string
    for _, obj := range objectList {
        objectNames = append(objectNames, *obj.Key)
    }

    // Initialize the bulk get request.
    bulkGetRequest := models.NewGetBulkJobSpectraS3Request(bucketName, objectNames)

    // Send the bulk get request to the server. Note that this creates the bulk get job,
    // but does not retrieve the objects.
    bulkGetResponse, err := client.GetBulkJobSpectraS3(bulkGetRequest)
    if err != nil {
        return nil, err
    }

    // Bulk jobs are split into multiple chunks which then need to be transferred.
    totalChunkCount := len(bulkGetResponse.MasterObjectList.Objects)
    curChunkCount := 0

    for curChunkCount < totalChunkCount {
        // Get the chunks that the server can send. The server may need to retrieve
        // objects into cache from the tape.
        chunksReady := models.NewGetJobChunksReadyForClientProcessingSpectraS3Request(bulkGetResponse.MasterObjectList.JobId)
        chunksReadyResponse, err := client.GetJobChunksReadyForClientProcessingSpectraS3(chunksReady)
        if err != nil {
            return nil, err
        }

        // Check to see if any chunks can be processed.
        numberOfChunks := len(chunksReadyResponse.MasterObjectList.Objects)
        if numberOfChunks > 0 {
            // Loop through all the chunks that are available for processing, and get
            // the files that are contained within them.
            for _, curChunk := range chunksReadyResponse.MasterObjectList.Objects {

                for _, curObj := range curChunk.Objects {
                    stopwatch := time.Now()

                    getObjRequest := models.NewGetObjectRequest(bucketName, *curObj.Name).
                        WithJob(bulkGetResponse.MasterObjectList.JobId).
                        WithOffset(curObj.Offset)

                    response, err := client.GetObject(getObjRequest)
                    if err != nil {
                        log.Fatal(err)
                    }
                    buf := make([]byte, curObj.Length)
                    contentLen, err := response.Content.Read(buf)
                    if err != nil {
                        return nil, err
                    }

                    lap := time.Since(stopwatch)
                    fmt.Printf("Got: %d ", contentLen)
                    fmt.Printf(" Time (s): %f\n", lap.Seconds())
                    ret = ret.AddInterval(lap.Seconds(), int64(contentLen))
                }
                curChunkCount++
            }
        } else {
            // When no chunks are returned we need to sleep to allow for cache space to
            // be freed.
            time.Sleep(time.Second * 5)
        }
        _, err = client.DeleteBucketSpectraS3(models.NewDeleteBucketSpectraS3Request(bucketName).WithForce())
        if err != nil {
            return nil, err
        }
    }
    return ret, nil
}
