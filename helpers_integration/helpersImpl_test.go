package helpers_integration

import (
    "testing"
    "log"
    "os"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_utils/ds3Testing"
    "github.com/SpectraLogic/ds3_go_sdk/helpers"
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_integration/utils"
    "io/ioutil"
    "github.com/SpectraLogic/ds3_go_sdk/samples/utils"
    "github.com/SpectraLogic/ds3_go_sdk/helpers/channels"
    "io"
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

    strategy := newTestTransferStrategy()

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

    strategy := newTestTransferStrategy()

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

    strategy := newTestTransferStrategy()

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


func TestGetBulk(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)
    err := testutils.PutTestBooks(client, testBucket)
    ds3Testing.AssertNilError(t, err)

    helper := helpers.NewHelpers(client)

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
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

    LoadLargeFile(testBucket, client)

    helper := helpers.NewHelpers(client)

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
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

func TestGetBulkPartialObjectRandomAccess(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    LoadLargeFile(testBucket, client)

    helper := helpers.NewHelpers(client)

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
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

    strategy := newTestTransferStrategy()

    channelBuilder := channels.NewReadChannelBuilder("does-not-exist")
    nonExistentPutObj := helperModels.PutObject{
        PutObject:      ds3Models.Ds3PutObject{Name:"does-not-exist",Size:10},
        ChannelBuilder: channelBuilder,
    }

    writeObjects := []helperModels.PutObject { nonExistentPutObj }

    jobId, err := helper.PutObjects(testBucket, writeObjects, strategy)

    // Verify that the expected error occurred instead of a panic
    if err == nil {
        t.Fatal("expected to get an error due to file not existing")
    }
    aggErr, ok := err.(*ds3Models.AggregateError)
    if !ok {
        t.Fatal("expected error to be of type AggregateError")
    }
    if len(aggErr.Errors) != 1 {
        t.Fatalf("expected 1 aggregate error, but got %d", len(aggErr.Errors))
    }
    expected := "open does-not-exist: no such file or directory"
    ds3Testing.AssertString(t, "expected error", expected, aggErr.Errors[0].Error())


    if jobId == "" {
        t.Error("expected to get a BP job ID, but instead got nothing")
    }

    // verify nothing is on the BP
    getBucket, getBucketErr := client.GetBucket(ds3Models.NewGetBucketRequest(testBucket))
    ds3Testing.AssertNilError(t, getBucketErr)
    if len(getBucket.ListBucketResult.Objects) != 0 {
        t.Fatalf("Expected '%d' objects in bucket '%s', but found '%d'.", 0, testBucket, len(getBucket.ListBucketResult.Objects))
    }
}