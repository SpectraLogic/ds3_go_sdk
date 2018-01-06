package ds3_integration

import (
    "testing"
    "sync"
    "log"
    "fmt"
    "runtime/debug"
    "spectra/ds3_go_sdk/ds3"
    "spectra/ds3_go_sdk/ds3_utils/ds3Testing"
    "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/ds3_integration/utils"
    "spectra/ds3_go_sdk/ds3/networking"
    "os"
    "spectra/ds3_go_sdk/ds3/buildclient"
    "net/url"
)

func putFileWithClient(name string, offset int64, jobId string, content *[]byte, group *sync.WaitGroup, t *testing.T) {
    defer group.Done()
    log.Printf("Putting file %s", name)

    _, err := client.PutObject(models.NewPutObjectRequest(testBucket, name, ds3.BuildByteReaderWithSizeDecorator(*content)).
        WithOffset(offset).
        WithJob(jobId))

    if err != nil {
        t.Errorf("Unexpected error when putting file '%s': '%s'.", name, err.Error())
        debug.PrintStack()
        t.Fail()
    }
}

func TestConcurrentClientUsage(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    content, err := testutils.LoadBook(testutils.BookTitles[0])
    ds3Testing.AssertNilError(t, err)

    size := int64(len(content))

    // create a job for putting n instances of file
    n := 30

    var ds3Objects []models.Ds3PutObject
    for i := 0; i < n; i++ {
        ds3Objects = append(ds3Objects, models.Ds3PutObject{ Name:fmt.Sprintf("file%d", i), Size:size })
    }

    bulkPut, err := client.PutBulkJobSpectraS3(models.NewPutBulkJobSpectraS3Request(testBucket, ds3Objects))
    ds3Testing.AssertNilError(t, err)

    // launch go routines to put files concurrently
    var group sync.WaitGroup
    group.Add(n)

    for _, chunk := range bulkPut.MasterObjectList.Objects {
        allocateChunk, allocateErr := client.AllocateJobChunkSpectraS3(models.NewAllocateJobChunkSpectraS3Request(chunk.ChunkId))
        ds3Testing.AssertNilError(t, allocateErr)
        for _, obj := range allocateChunk.Objects.Objects {
            go putFileWithClient(*obj.Name, obj.Offset, bulkPut.MasterObjectList.JobId, &content, &group, t)
        }
    }

    group.Wait()
}

func putFileWithSendNetwork(name string, offset int64, jobId string, content *[]byte, bucketName string, group *sync.WaitGroup, t *testing.T, network networking.Network, info *networking.ConnectionInfo) {
    defer group.Done()

    // manually create http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(ds3.HTTP_VERB_PUT).
        WithPath("/" + bucketName + "/" + name).
        WithOptionalQueryParam("job", &jobId).
        WithOptionalQueryParam("offset", networking.Int64PtrToStrPtr(&offset)).
        WithReader(ds3.BuildByteReaderWithSizeDecorator(*content)).
        WithChecksum(models.NewNoneChecksum()).
        WithHeaders(make(map[string]string)).
        Build(info)

    //networkRetryDecorator := networking.NewNetworkRetryDecorator(network, 5)
    t.Logf("Created request to for file '%s' to url '%s'", name, httpRequest.URL.String())
    response, err := network.Invoke(httpRequest)
    if err != nil {
        t.Errorf("Unexpected error when putting file '%s': '%s'.", name, err.Error())
        debug.PrintStack()
        t.Fail()
    }

    _, err = models.NewPutObjectResponse(response)
    if err != nil {
        t.Errorf("Unexpected error when putting file '%s': '%s'.", name, err.Error())
        debug.PrintStack()
        t.Fail()
    }
}

func TestConcurrentSendNetworkUsage(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    content, err := testutils.LoadBook(testutils.BookTitles[0])
    ds3Testing.AssertNilError(t, err)

    size := int64(len(content))

    // create a job for putting n instances of file
    n := 30

    var ds3Objects []models.Ds3PutObject
    for i := 0; i < n; i++ {
        ds3Objects = append(ds3Objects, models.Ds3PutObject{ Name:fmt.Sprintf("file%d", i), Size:size })
    }

    bulkPut, err := client.PutBulkJobSpectraS3(models.NewPutBulkJobSpectraS3Request(testBucket, ds3Objects))
    ds3Testing.AssertNilError(t, err)

    // launch go routines to put files concurrently
    var group sync.WaitGroup
    group.Add(n)

    //Retrieve the environment variables
    endpoint := os.Getenv(buildclient.EndpointEnv)
    accessKey := os.Getenv(buildclient.AccessKeyEnv)
    secretKey := os.Getenv(buildclient.SecretKeyEnv)

    endpointUrl, err := url.Parse(endpoint)
    ds3Testing.AssertNilError(t, err)
    info := networking.ConnectionInfo{
        Endpoint:    endpointUrl,
        Credentials: &networking.Credentials{AccessId: accessKey, Key: secretKey},
        Proxy:       nil}

    network := networking.NewSendNetwork(&info)

    for _, chunk := range bulkPut.MasterObjectList.Objects {
        //TODO Try get available job chunks ready for client processing instead
        allocateChunk, allocateErr := client.AllocateJobChunkSpectraS3(models.NewAllocateJobChunkSpectraS3Request(chunk.ChunkId))
        ds3Testing.AssertNilError(t, allocateErr)
        for _, obj := range allocateChunk.Objects.Objects {
            go putFileWithSendNetwork(*obj.Name, obj.Offset, bulkPut.MasterObjectList.JobId, &content, testBucket, &group, t, network, &info)
        }
    }

    group.Wait()
}