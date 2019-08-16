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
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/buildclient"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "github.com/SpectraLogic/ds3_go_sdk/samples/utils"
    "log"
    "strconv"
    "time"
)

func PerormancePutSample(numFiles int, blockSize int64, bucketName string) (*utils.PerformanceTest, error) {
    ret := utils.NewPerformanceTest(numFiles, blockSize)
    fmt.Println("---- Put Performance Data ----")

    // Create a client from environment variables.
    client, err := buildclient.FromEnv()
    if err != nil {
        return nil, err
    }

    // Create a bucket where our files will be stored.
    _, err = client.PutBucket(models.NewPutBucketRequest(bucketName))
    if err != nil {
        // warn but don't fail on bucket exists
        ret.Error = fmt.Sprintf("Warning: Bucket Exits?", err)
    }

    // Create the list of Ds3PutObjects to be put to the Black Pearl via bulk put
    var ds3PutObjects []models.Ds3PutObject
    for fileNum := 0; fileNum < numFiles; fileNum++ {
        fileName := utils.PerformanceFilePrefix + strconv.Itoa(fileNum)
        curObj := models.Ds3PutObject{ Name:fileName, Size:blockSize}
        ds3PutObjects = append(ds3PutObjects, curObj)
    }

    // Create the bulk put request.
    putBulkRequest := models.NewPutBulkJobSpectraS3Request(bucketName, ds3PutObjects)

    // Send the bulk put request to the server. This creates the job, but does not
    // directly send the data.
    putBulkResponse, err := client.PutBulkJobSpectraS3(putBulkRequest)
    if err != nil {
        return nil, err
    }

    // The bulk request will split the files over several chunks if it needs to.
    // We then need to ask what chunks we can send, and then send them making sure
    // we don't resend the same chunks.

    // The bulk job is split into multiple chunks, which then need to be transferred.
    totalChunkCount := len(putBulkResponse.MasterObjectList.Objects)
    curChunkCount := 0

    for curChunkCount < totalChunkCount {
        // Get the list of available chunks that the server can receive. The server may
        // not be able to receive everything, so not all chunks will necessarily be
        // returned
        chunksReady := models.NewGetJobChunksReadyForClientProcessingSpectraS3Request(putBulkResponse.MasterObjectList.JobId)
        chunksReadyResponse, err := client.GetJobChunksReadyForClientProcessingSpectraS3(chunksReady)
        if err != nil {
            return nil, err
        }

        // Check to see if any chunks can be processed
        numberOfChunks := len(chunksReadyResponse.MasterObjectList.Objects)
        if numberOfChunks > 0 {
            // Loop through all the chunks that are available for processing, and send
            // the files that are contained within them.
            for _, curChunk := range chunksReadyResponse.MasterObjectList.Objects  {

                for _, curObj := range curChunk.Objects {
                    stopwatch := time.Now()
                    reader := utils.BuildPerformanceReaderWithSizeDecorator(blockSize)

                    putObjRequest := models.NewPutObjectRequest(bucketName, *curObj.Name, reader).
                        WithJob(chunksReadyResponse.MasterObjectList.JobId).
                        WithOffset(curObj.Offset)

                    _, err = client.PutObject(putObjRequest)
                    if err != nil {
                        log.Fatal(err)
                    }
                    lap := time.Since(stopwatch)
                    fmt.Printf("Sent: %d ", curObj.Length - curObj.Offset)
                    fmt.Printf(" Time (s): %f\n", lap.Seconds())
                    ret = ret.AddInterval(lap.Seconds(), curObj.Length - curObj.Offset)
                }
                curChunkCount++
            }
        } else {
            // When no chunks are returned we need to sleep to allow for cache space to
            // be freed.
            time.Sleep(time.Second * 5)
        }
    }
    return ret, nil
}
