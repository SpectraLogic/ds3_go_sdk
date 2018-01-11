package helpers_integration

import (
    "testing"
    "log"
    "os"
    "spectra/ds3_go_sdk/ds3"
    "spectra/ds3_go_sdk/ds3_utils/ds3Testing"
    "spectra/ds3_go_sdk/helpers"
    "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/ds3_integration/utils"
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

    err = helper.PutObjects(*writeObjects, testBucket, strategy)
    ds3Testing.AssertNilError(t, err)

    // verify all books are on BP
    getBucket, getBucketErr := client.GetBucket(models.NewGetBucketRequest(testBucket))
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
    objName := "lesmis-copies.txt"
    path := "./resources/bigfiles/"
    defer testutils.DeleteBucketContents(client, testBucket)

    helper := helpers.NewHelpers(client)

    strategy := newTestTransferStrategy()

    writeObj, err := getTestWriteObjectRandomAccess(objName, path + objName)

    var writeObjects []helpers.PutObject
    writeObjects = append(writeObjects, *writeObj)

    ds3Testing.AssertNilError(t, err)

    err = helper.PutObjects(writeObjects, testBucket, strategy)
    ds3Testing.AssertNilError(t, err)


    testutils.VerifyFilesOnBP(t, testBucket, []string {objName}, path, client)
}

func TestPutBulkBlobSpanningChunksStreamAccess(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)
    objName := "lesmis-copies.txt"
    path := "./resources/bigfiles/"

    helper := helpers.NewHelpers(client)

    strategy := newTestTransferStrategy()

    writeObj, err := getTestWriteObjectStreamAccess(objName, path + objName)

    var writeObjects []helpers.PutObject
    writeObjects = append(writeObjects, *writeObj)

    ds3Testing.AssertNilError(t, err)

    err = helper.PutObjects(writeObjects, testBucket, strategy)
    ds3Testing.AssertNilError(t, err)

    testutils.VerifyFilesOnBP(t, testBucket, []string {objName}, path, client)
}
