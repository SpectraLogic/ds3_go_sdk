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

package performance

import (
    "bufio"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "log"
    "strconv"
    "sync"
    "time"
)

// Demonstrates how to perform a bulk get. Assumes that the target bucket already exists
// and has files, i.e. run putBulkSample.go first.
func PerformanceGetSample(client *ds3.Client, numThreads int, filesPerThread int, fileSize int64, bucketName string) (*PerformanceTest, error) {
    ret := NewPerformanceTest(numThreads, filesPerThread, fileSize)

    defer deleteBucket(client, bucketName)

    burstsChannel := make(chan PerformanceBurst, numThreads)
    var wg sync.WaitGroup
    stopwatch := time.Now()
    for threadNum := 1; threadNum <= numThreads; threadNum++ {
        wg.Add(1)
        go getTestObjects(client, threadNum, ret, bucketName, burstsChannel, &wg)
    }
    wg.Wait()
    lap := time.Since(stopwatch)
    totalSeconds := lap.Seconds()


    // package all interval data
    bursts := []PerformanceBurst{}
    for threadNum := 1; threadNum <= numThreads; threadNum++ {
        bursts = append(bursts, <-burstsChannel)
    }
    ret.Bursts = bursts
    ret.Seconds = totalSeconds
    return ret, nil
}

func getTestObjects(client *ds3.Client, threadNum int, test *PerformanceTest, bucketName string, c chan PerformanceBurst, wg *sync.WaitGroup) {

    defer wg.Done()
    ret := NewPerformanceBurst(threadNum)

    // Create the list of Ds3GetObjects to be retrieved from the Black Pearl
    var getObjects []string
    for fileNum := 0; fileNum < test.NumFiles; fileNum++ {
        fileName := PerformanceFilePrefix + strconv.Itoa(threadNum) + "_" + strconv.Itoa(fileNum)
        getObjects = append(getObjects, fileName)
    }

    // Initialize the bulk get request.
    bulkGetRequest := models.NewGetBulkJobSpectraS3Request(bucketName, getObjects)

    // Send the bulk get request to the server. Note that this creates the bulk get job,
    // but does not retrieve the objects.
    bulkGetResponse, err := client.GetBulkJobSpectraS3(bulkGetRequest)
    if err != nil {
        c <- *NewPerformanceBurstError(fmt.Sprintf("Cound not get job chuncks", err))
        return
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
            c <- *NewPerformanceBurstError(fmt.Sprintf("Cound not process job chuncks", err))
           return
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
                    bufferredReader := bufio.NewReader(response.Content)
                    contentLen, err := bufferredReader.Discard(int(test.FileSize))
                    if err != nil {
                        c <- *NewPerformanceBurstError(fmt.Sprintf("Cound not read object data", err))
                        return
                    }
                    bufferredReader.Reset(response.Content)

                    lap := time.Since(stopwatch)
                    ret = ret.AddInterval(lap.Seconds(), int64(contentLen))
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


func deleteBucket(client *ds3.Client, bucketName string) (*models.DeleteBucketSpectraS3Response, error) {
    return client.DeleteBucketSpectraS3(models.NewDeleteBucketSpectraS3Request(bucketName).WithForce())
}
