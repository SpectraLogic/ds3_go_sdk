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
    "spectra/ds3_go_sdk/ds3/buildclient"
    "log"
    "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/samples/utils"
    "time"
)

// Demonstrates how to perform a bulk get. Assumes that the target bucket already exists
// and has files, i.e. run putBulkSample.go first.
func main() {
    // Create a client from environment variables.
    client, err := buildclient.FromEnv()
    if err != nil {
        log.Fatal(err)
    }

    // Get a list of all objects in the target bucket.
    bucketResponse, err := client.GetBucket(models.NewGetBucketRequest(utils.BucketName))
    if err != nil {
        log.Fatal(err)
    }

    objectList := bucketResponse.ListBucketResult.Objects

    // Convert object list returned from get bucket into a list of object names.
    var objectNames []string
    for _, obj := range objectList {
        objectNames = append(objectNames, *obj.Key)
    }

    // Initialize the bulk get request.
    bulkGetRequest := models.NewGetBulkJobSpectraS3Request(utils.BucketName, objectNames)

    // Send the bulk get request to the server. Note that this creates the bulk get job,
    // but does not retrieve the objects.
    bulkGetResponse, err := client.GetBulkJobSpectraS3(bulkGetRequest)
    if err != nil {
        log.Fatal(err)
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
            log.Fatal(err)
        }

        // Check to see if any chunks can be processed.
        numberOfChunks := len(chunksReadyResponse.MasterObjectList.Objects)
        if numberOfChunks > 0 {
            // Loop through all the chunks that are available for processing, and get
            // the files that are contained within them.
            for _, curChunk := range chunksReadyResponse.MasterObjectList.Objects {

                for _, curObj := range curChunk.Objects {

                    getObjRequest := models.NewGetObjectRequest(utils.BucketName, *curObj.Name).
                        WithJob(bulkGetResponse.MasterObjectList.JobId).
                        WithOffset(curObj.Offset)

                    getObjResponse, err := client.GetObject(getObjRequest)
                    if err != nil {
                        log.Fatal(err)
                    }

                    err = utils.VerifyBookContent(*curObj.Name, getObjResponse.Content)
                    if err != nil {
                        log.Fatal(err)
                    }
                }
                curChunkCount++
            }
        } else {
            // When no chunks are returned we need to sleep to allow for cache space to
            // be freed.
            time.Sleep(time.Second * 5)
        }
    }
}
