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
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/buildclient"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "github.com/SpectraLogic/ds3_go_sdk/samples/utils"
    "strconv"
    "sync"
    "time"
)

func PerormancePutSample(numThreads int, filesPerThread int, fileSize int64, bucketName string) (*utils.PerformanceTest, error) {
    fmt.Println("---- Put Performance Data ----")

    ret := utils.NewPerformanceTest(numThreads, filesPerThread, fileSize)

    // Create a client from environment variables.
    client, err := buildclient.FromEnv()
    if err != nil {
        return nil, err
    }

    // Create a bucket where our files will be stored.
    _, err = client.PutBucket(models.NewPutBucketRequest(bucketName))
    if err != nil {
        return nil, fmt.Errorf("Bucket Exists?", err)
    }

    burstsChannel := make(chan utils.PerformanceBurst, numThreads)
    var wg sync.WaitGroup
    stopwatch := time.Now()
    for threadNum := 1; threadNum <= numThreads; threadNum++ {
        wg.Add(1)
        go putTestObjects(client, threadNum, ret, bucketName, burstsChannel, &wg)
    }
    wg.Wait()
    lap := time.Since(stopwatch)
    totalSeconds := lap.Seconds()

    // package all interval data
    bursts := []utils.PerformanceBurst{}
    for threadNum := 1; threadNum <= numThreads; threadNum++ {
        fmt.Println("Reading channel")
        bursts = append(bursts, <-burstsChannel)
    }
    ret.Bursts = bursts
    ret.Seconds = totalSeconds
    return ret, err
}

func putTestObjects(client *ds3.Client, threadNum int, test *utils.PerformanceTest, bucketName string, c chan utils.PerformanceBurst, wg *sync.WaitGroup) {

    defer wg.Done()
    ret := utils.NewPerformanceBurst(threadNum)

    // Create the list of Ds3PutObjects to be put to the Black Pearl via bulk put
    var ds3PutObjects []models.Ds3PutObject
    for fileNum := 0; fileNum < test.NumFiles; fileNum++ {
        fileName := utils.PerformanceFilePrefix + strconv.Itoa(threadNum) + "_" + strconv.Itoa(fileNum)
        curObj := models.Ds3PutObject{ Name:fileName, Size: test.FileSize}
        ds3PutObjects = append(ds3PutObjects, curObj)
    }

    // Create the bulk put request.
    putBulkRequest := models.NewPutBulkJobSpectraS3Request(bucketName, ds3PutObjects)

    // Send the bulk put request to the server. This creates the job, but does not
    // directly send the data.
    putBulkResponse, err := client.PutBulkJobSpectraS3(putBulkRequest)
    if err != nil {
        c <- *utils.NewPerformanceBurstError(fmt.Sprintf("Cound not create request", err))
        return
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
            c <- *utils.NewPerformanceBurstError(fmt.Sprintf("Cound not get job chuncks", err))
            return
        }

        // Check to see if any chunks can be processed
        numberOfChunks := len(chunksReadyResponse.MasterObjectList.Objects)
        if numberOfChunks > 0 {
            // Loop through all the chunks that are available for processing, and send
            // the files that are contained within them.
            for _, curChunk := range chunksReadyResponse.MasterObjectList.Objects  {

                for _, curObj := range curChunk.Objects {
                    stopwatch := time.Now()
                    reader := utils.BuildPerformanceReaderWithSizeDecorator(test.FileSize)

                    putObjRequest := models.NewPutObjectRequest(bucketName, *curObj.Name, reader).
                        WithJob(chunksReadyResponse.MasterObjectList.JobId).
                        WithOffset(curObj.Offset)

                    _, err = client.PutObject(putObjRequest)
                    if err != nil {
                        c <- *utils.NewPerformanceBurstError(fmt.Sprintf("Cound not put object", err))
                        return
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
    c <- *ret
}
