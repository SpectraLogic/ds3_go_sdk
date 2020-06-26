package helpers_integration

import (
    "bytes"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_integration/utils"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_utils/ds3Testing"
    "github.com/SpectraLogic/ds3_go_sdk/helpers"
    "github.com/SpectraLogic/ds3_go_sdk/helpers/channels"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "github.com/SpectraLogic/ds3_go_sdk/samples/utils"
    "io"
    "io/ioutil"
    "log"
    "os"
    "sync"
    "testing"
    "time"
)

var client *ds3.Client
var testBucket = "GoHelperTestBucket"
var envTestNameSpace = "GoHelperTest"
var defaultUser = "Administrator"

func TestMain(m *testing.M) {
    var err error
    var exitVal int
    var ids *testutils.TestIds
    client, ids, err = testutils.SetupTestEnv(testBucket, defaultUser, envTestNameSpace)
    if err != nil {
        log.Printf("Unable to setup test environment '%s'.", err.Error())
        exitVal = 1
    } else {
        exitVal = m.Run()
    }
    testutils.TeardownTestEnv(client, ids, testBucket)
    os.Exit(exitVal)
}

func TestPutBulk(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)
    helper := helpers.NewHelpers(client)

    strategy := newTestTransferStrategy(t)

    writeObjects, err := getTestBooksAsWriteObjects()
    ds3Testing.AssertNilError(t, err)

    jobId, err := helper.PutObjects(testBucket, *writeObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    if jobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }

    // verify all books are on BP
    getBucket, getBucketErr := client.GetBucket(ds3Models.NewGetBucketRequest(testBucket))
    ds3Testing.AssertNilError(t, getBucketErr)
    if len(getBucket.ListBucketResult.Objects) != len(*writeObjects) {
        t.Fatalf("Expected '%d' objects in bucket '%s', but found '%d'.", len(*writeObjects), testBucket, len(getBucket.ListBucketResult.Objects))
    }
    for i, obj := range getBucket.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "BookName", testutils.BookTitles[i], obj.Key)
    }

    testutils.VerifyFilesOnBP(t, testBucket, testutils.BookTitles, testutils.BookPath, client)
}

func TestPutBulkBlobSpanningChunksRandomAccess(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    helper := helpers.NewHelpers(client)

    strategy := newTestTransferStrategy(t)

    writeObj, err := getTestWriteObjectRandomAccess(LargeBookTitle, LargeBookPath + LargeBookTitle)

    var writeObjects []helperModels.PutObject
    writeObjects = append(writeObjects, *writeObj)

    ds3Testing.AssertNilError(t, err)

    jobId, err := helper.PutObjects(testBucket, writeObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    if jobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }


    testutils.VerifyFilesOnBP(t, testBucket, []string {LargeBookTitle}, LargeBookPath, client)
}

func TestPutBulkBlobSpanningChunksStreamAccess(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    helper := helpers.NewHelpers(client)

    strategy := newTestTransferStrategy(t)

    writeObj, err := getTestWriteObjectStreamAccess(LargeBookTitle, LargeBookPath + LargeBookTitle)

    var writeObjects []helperModels.PutObject
    writeObjects = append(writeObjects, *writeObj)

    ds3Testing.AssertNilError(t, err)

    jobId, err := helper.PutObjects(testBucket, writeObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    if jobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }

    testutils.VerifyFilesOnBP(t, testBucket, []string {LargeBookTitle}, LargeBookPath, client)
}

// GOSDK-26: On archive helpers can hang indefinitely using streaming strategy if blob fails.
// Test validates the above jira is fixed.
func TestPutBulkBlobSpanningChunksStreamAccessDoesNotExist(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    helper := helpers.NewHelpers(client)

    errorCallbackCalled := false
    var mutex sync.Mutex
    errorCallback := func(objectName string, err error) {
        mutex.Lock()
        errorCallbackCalled = true
        mutex.Unlock()

        ds3Testing.AssertString(t, "object name", LargeBookTitle, objectName)
    }

    strategy := helpers.WriteTransferStrategy{
        BlobStrategy: newTestBlobStrategy(),
        Options:      helpers.WriteBulkJobOptions{MaxUploadSize: &helpers.MinUploadSize},
        Listeners:    helpers.ListenerStrategy{ErrorCallback:errorCallback},
    }

    // open a file but lie that its bigger than it is
    f, err := os.Open(testutils.BookPath + testutils.BookTitles[0])
    writeObj := helperModels.PutObject{
        PutObject: ds3Models.Ds3PutObject{Name: LargeBookTitle, Size: 20*1024*1024},
        ChannelBuilder: &testStreamAccessReadChannelBuilder{f: f},
    }

    var writeObjects []helperModels.PutObject
    writeObjects = append(writeObjects, writeObj)

    ds3Testing.AssertNilError(t, err)

    _, err = helper.PutObjects(testBucket, writeObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertBool(t, "error callback called", true, errorCallbackCalled)
}

func TestGetBulk(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)
    err := testutils.PutTestBooks(client, testBucket)
    ds3Testing.AssertNilError(t, err)

    helper := helpers.NewHelpers(client)

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
        Listeners: newErrorOnErrorListenerStrategy(t),
    }

    file0, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file0.Close()
    defer os.Remove(file0.Name())

    file1, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file1.Close()
    defer os.Remove(file1.Name())

    file2, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file2.Close()
    defer os.Remove(file2.Name())

    file3, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file3.Close()
    defer os.Remove(file3.Name())

    readObjects := []helperModels.GetObject {
        {Name: testutils.BookTitles[0], ChannelBuilder: channels.NewWriteChannelBuilder(file0.Name())},
        {Name: testutils.BookTitles[1], ChannelBuilder: channels.NewWriteChannelBuilder(file1.Name())},
        {Name: testutils.BookTitles[2], ChannelBuilder: channels.NewWriteChannelBuilder(file2.Name())},
        {Name: testutils.BookTitles[3], ChannelBuilder: channels.NewWriteChannelBuilder(file3.Name())},
    }

    jobId, err := helper.GetObjects(testBucket, readObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    if jobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }

    utils.VerifyBookContent(testutils.BookTitles[0], file0)
    utils.VerifyBookContent(testutils.BookTitles[1], file1)
    utils.VerifyBookContent(testutils.BookTitles[2], file2)
    utils.VerifyBookContent(testutils.BookTitles[3], file3)
}

func TestGetBulkBlobSpanningChunksRandomAccess(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    LoadLargeFile(t, testBucket, client)

    helper := helpers.NewHelpers(client)

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
        Listeners: newErrorOnErrorListenerStrategy(t),
    }

    file, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file.Close()
    defer os.Remove(file.Name())

    readObjects := []helperModels.GetObject{
        {Name: LargeBookTitle, ChannelBuilder: channels.NewWriteChannelBuilder(file.Name())},
    }

    jobId, err := helper.GetObjects(testBucket, readObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    if jobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }

    err = VerifyLargeBookContent(file)
    ds3Testing.AssertNilError(t, err)
}

func TestGetBulkBlobSpanningChunksStreaming(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    LoadLargeFile(t, testBucket, client)

    helper := helpers.NewHelpers(client)

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
        Listeners: newErrorOnErrorListenerStrategy(t),
    }

    file, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file.Close()
    defer os.Remove(file.Name())

    fileWriter, err := os.OpenFile(file.Name(), os.O_WRONLY | os.O_CREATE, os.ModePerm)
    defer fileWriter.Close()

    readObjects := []helperModels.GetObject{
        {Name: LargeBookTitle, ChannelBuilder: &testStreamAccessWriteChannelBuilder{f: fileWriter}},
    }

    jobId, err := helper.GetObjects(testBucket, readObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    if jobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }

    err = VerifyLargeBookContent(file)
    ds3Testing.AssertNilError(t, err)
}

func TestGetBulkBlobSpanningChunksStreamingFailBlob(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    LoadLargeFile(t, testBucket, client)

    helper := helpers.NewHelpers(client)

    errorCallbackCalled := false
    var mutex sync.Mutex
    errorCallback := func(objectName string, err error) {
        mutex.Lock()
        errorCallbackCalled = true
        mutex.Unlock()

        ds3Testing.AssertString(t, "object name", LargeBookTitle, objectName)
    }

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
        Listeners: helpers.ListenerStrategy{
            ErrorCallback: errorCallback,
        },
    }

    file, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file.Close()
    defer os.Remove(file.Name())

    readObjects := []helperModels.GetObject{
        {Name: LargeBookTitle, ChannelBuilder: &testStreamAccessWriteFailOnFirstBlob{}},
    }

    _, err = helper.GetObjects(testBucket, readObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertBool(t, "error callback called", true, errorCallbackCalled)
}

func TestGetBulkPartialObjectRandomAccess(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    LoadLargeFile(t, testBucket, client)

    helper := helpers.NewHelpers(client)

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
        Listeners: newErrorOnErrorListenerStrategy(t),
    }

    file, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file.Close()
    defer os.Remove(file.Name())

    ranges := []ds3Models.Range {
        {0, 100},
        {200, 300},
        {301, 400},
        {500, 600},
    }

    readObjects := []helperModels.GetObject{
        {Name: LargeBookTitle, ChannelBuilder: channels.NewPartialObjectChannelBuilder(file.Name(), ranges), Ranges: ranges},
    }

    jobId, err := helper.GetObjects(testBucket, readObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    if jobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }

    file.Seek(0, io.SeekStart)
    testutils.VerifyPartialFile(t, LargeBookPath + LargeBookTitle, 101, 0, file)
    testutils.VerifyPartialFile(t, LargeBookPath + LargeBookTitle, 201, 200, file)
    testutils.VerifyPartialFile(t, LargeBookPath + LargeBookTitle, 101, 500, file)
}

func TestPutObjectDoesNotExist(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)
    helper := helpers.NewHelpers(client)
    const testObjectName = "does-not-exist"

    errorCallbackCalled := false
    var mutex sync.Mutex
    errorCallback := func(objectName string, err error) {
        mutex.Lock()
        errorCallbackCalled = true
        mutex.Unlock()

        ds3Testing.AssertString(t, "errored object name", testObjectName, objectName)
        ds3Testing.AssertBool(t, "error expected to by of type NotExist", true, os.IsNotExist(err))
    }

    strategy := helpers.WriteTransferStrategy{
        BlobStrategy: newTestBlobStrategy(),
        Options:      helpers.WriteBulkJobOptions{MaxUploadSize: &helpers.MinUploadSize},
        Listeners:    helpers.ListenerStrategy{ErrorCallback:errorCallback},
    }

    channelBuilder := channels.NewReadChannelBuilder(testObjectName)
    nonExistentPutObj := helperModels.PutObject{
        PutObject:      ds3Models.Ds3PutObject{Name:testObjectName,Size:10},
        ChannelBuilder: channelBuilder,
    }

    writeObjects := []helperModels.PutObject { nonExistentPutObj }

    _, err := helper.PutObjects(testBucket, writeObjects, strategy)
    ds3Testing.AssertNilError(t, err)

    ds3Testing.AssertBool(t, "error callback was called", true, errorCallbackCalled)
}

type closeWrapper struct {
    io.Reader
}

func (wrapper *closeWrapper) Close() error {
    return nil
}

type emptyReadChannelBuilder struct {
    channels.FatalErrorHandler
}

func (builder *emptyReadChannelBuilder) GetChannel(_ int64) (io.ReadCloser, error) {
    reader := bytes.NewReader([]byte{})
    return &closeWrapper{Reader: reader}, nil
}

func (builder *emptyReadChannelBuilder) IsChannelAvailable(_ int64) bool {
    return true
}

func (builder *emptyReadChannelBuilder) OnDone(reader io.ReadCloser) {
    reader.Close()
}

type emptyWriteCloser struct {}

func (writer *emptyWriteCloser) Write(p []byte) (n int, err error) {
    return len(p), nil
}

func (writer *emptyWriteCloser) Close() error {
    return nil
}

type emptyWriteChannelBuilder struct {
    channels.FatalErrorHandler
}

func (builder *emptyWriteChannelBuilder) GetChannel(_ int64) (io.WriteCloser, error) {
    return &emptyWriteCloser{}, nil
}

func (builder *emptyWriteChannelBuilder) IsChannelAvailable(_ int64) bool {
    return true
}

func (builder *emptyWriteChannelBuilder) OnDone(writer io.WriteCloser) {
    writer.Close()
}

func TestBulkPutAndGetLotsOfFiles(t *testing.T) {
    defer func() {
        // force delete the bucket because its faster than deleting all the files
        testutils.DeleteBucket(client, testBucket)
        // re-creating the bucket after force delete because other tests expect it to exist
        testutils.PutBucket(client, testBucket)
    }()
    helper := helpers.NewHelpers(client)

    // put a bunch of files
    var writeObjects []helperModels.PutObject

    const numObjects = 18650
    for i := 0; i < numObjects; i++ {
        objName := fmt.Sprintf("file-%d.txt", i)
        curWriteObj := helperModels.PutObject{
            PutObject:      ds3Models.Ds3PutObject{Name:objName, Size:0},
            ChannelBuilder: &emptyReadChannelBuilder{},
        }
        writeObjects = append(writeObjects, curWriteObj)
    }

    writeStrategy := helpers.WriteTransferStrategy{
        BlobStrategy: &helpers.SimpleBlobStrategy{
            Delay:time.Second * 30,
            MaxConcurrentTransfers:10,
        },
        Options:   helpers.WriteBulkJobOptions{MaxUploadSize: &helpers.MinUploadSize},
        Listeners: newErrorOnErrorListenerStrategy(t),
    }

    jobId, err := helper.PutObjects(testBucket, writeObjects, writeStrategy)
    ds3Testing.AssertNilError(t, err)
    if jobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }

    // retrieve said files
    var readObjects []helperModels.GetObject
    for _, writeObj := range writeObjects {
        curGetObj := helperModels.GetObject{
            Name:           writeObj.PutObject.Name,
            ChannelBuilder: &emptyWriteChannelBuilder{},
        }
        readObjects = append(readObjects, curGetObj)
    }

    readStrategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: &helpers.SimpleBlobStrategy{
            Delay:time.Second * 30,
            MaxConcurrentTransfers:10,
        },
        Listeners: newErrorOnErrorListenerStrategy(t),
    }

    jobId, err = helper.GetObjects(testBucket, readObjects, readStrategy)
    ds3Testing.AssertNilError(t, err)

    if jobId == "" {
        t.Errorf("expected to get a BP job ID, but instead got nothing")
    }
}

func TestRetryGettingBlobRange(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    helper := helpers.NewHelpers(client)
    strategy := newTestTransferStrategy(t)

    // Put a blobbed object to BP
    const bigFilePath = LargeBookPath + LargeBookTitle
    writeObj, err := getTestWriteObjectRandomAccess(LargeBookTitle, bigFilePath)
    ds3Testing.AssertNilError(t, err)

    var writeObjects []helperModels.PutObject
    writeObjects = append(writeObjects, *writeObj)

    putJobId, err := helper.PutObjects(testBucket, writeObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    if putJobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }

    // Try to get some data from each blob
    getJob, err := client.GetJobSpectraS3(ds3Models.NewGetJobSpectraS3Request(putJobId))
    ds3Testing.AssertNilError(t, err)

    blobsChecked := 0
    for _, curObj := range getJob.MasterObjectList.Objects {
        for _, blob := range curObj.Objects {
            func() {
                // create a temp file for writing the blob to
                tempFile, err := ioutil.TempFile("", "go-sdk-test-")
                ds3Testing.AssertNilError(t, err)
                defer func() {
                    tempFile.Close()
                    os.Remove(tempFile.Name())
                }()

                // get a range of the blob
                startRange := blob.Offset+10 // retrieve subset of blob
                endRange := blob.Length+blob.Offset-1
                bytesWritten, err := helpers.RetryGettingBlobRange(client, testBucket, writeObj.PutObject.Name, blob.Offset, startRange, endRange, tempFile, client.Logger)
                ds3Testing.AssertNilError(t, err)
                ds3Testing.AssertInt64(t, "bytes written", endRange-startRange+1, bytesWritten)

                // verify that retrieved partial blob is correct
                err = tempFile.Sync()
                ds3Testing.AssertNilError(t, err)

                tempFile.Seek(0, 0)
                length := endRange-startRange
                testutils.VerifyPartialFile(t, bigFilePath, length, startRange, tempFile)
            }()
            blobsChecked++
        }
    }
    if blobsChecked == 0 {
        t.Fatalf("didn't verify any blobs")
    }
}

func TestGetRemainingBlob(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    helper := helpers.NewHelpers(client)
    strategy := newTestTransferStrategy(t)

    // Put a blobbed object to BP
    const bigFilePath = LargeBookPath + LargeBookTitle
    writeObj, err := getTestWriteObjectRandomAccess(LargeBookTitle, bigFilePath)
    ds3Testing.AssertNilError(t, err)

    var writeObjects []helperModels.PutObject
    writeObjects = append(writeObjects, *writeObj)

    putJobId, err := helper.PutObjects(testBucket, writeObjects, strategy)
    ds3Testing.AssertNilError(t, err)
    if putJobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }

    // Try to get some data from each blob
    getJob, err := client.GetJobSpectraS3(ds3Models.NewGetJobSpectraS3Request(putJobId))
    ds3Testing.AssertNilError(t, err)

    blobsChecked := 0
    for _, curObj := range getJob.MasterObjectList.Objects {
        for _, blob := range curObj.Objects {
            func() {
                // create a temp file for writing the blob to
                tempFile, err := ioutil.TempFile("", "go-sdk-test-")
                ds3Testing.AssertNilError(t, err)
                defer func() {
                    tempFile.Close()
                    os.Remove(tempFile.Name())
                }()

                // get the remainder of the blob after skipping some bytes
                blob := helperModels.NewBlobDescription(*blob.Name, blob.Offset, blob.Length)
                var amountToSkip int64 = 10
                err = helpers.GetRemainingBlob(client, testBucket, &blob, amountToSkip, tempFile, client.Logger)
                ds3Testing.AssertNilError(t, err)

                // verify that retrieved partial blob is correct
                err = tempFile.Sync()
                ds3Testing.AssertNilError(t, err)

                tempFile.Seek(0, 0)
                length := blob.Length() - amountToSkip
                testutils.VerifyPartialFile(t, bigFilePath, length, blob.Offset()+amountToSkip, tempFile)
            }()
            blobsChecked++
        }
    }
    if blobsChecked == 0 {
        t.Fatalf("didn't verify any blobs")
    }
}
