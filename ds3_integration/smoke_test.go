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

package ds3_integration

import (
    "bytes"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_integration/utils"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_utils/ds3Testing"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "testing"
)

var client *ds3.Client
var testBucket = "GoIntegrationTestBucket"
var envTestNameSpace = "GoIntegrationSmokeTest"
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

func TestGetService(t *testing.T) {
    response, err := client.GetService(models.NewGetServiceRequest())
    ds3Testing.AssertNilError(t, err)
    if response == nil {
        t.Fatal("Received an unexpected nil response.")
    }
}

func TestBucket(t *testing.T) {
    //Create bucket
    bucketName := "GoTestPutBucket"
    putErr := testutils.PutBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, putErr)
    defer testutils.DeleteBucketLogError(t, client, bucketName)

    //Verify that bucket exists
    getBucketResponse, getErr := testutils.GetBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, getErr)
    ds3Testing.AssertNonNilStringPtr(t, "Name", bucketName, getBucketResponse.ListBucketResult.Name)
}

func TestObject(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)
    beowulf := "beowulf.txt"

    //Put object to BP
    book, bookErr := testutils.LoadBookLogError(t, beowulf)
    ds3Testing.AssertNilError(t, bookErr)
    defer testutils.DeleteObjectLogError(t, client, testBucket, beowulf)

    putObjErr := testutils.PutObjectLogError(t, client, testBucket, beowulf, book)
    ds3Testing.AssertNilError(t, putObjErr)

    //Verify that object exists
    getObjectResponse, getObjErr := testutils.GetObjectLogError(t, client, testBucket, beowulf)
    ds3Testing.AssertNilError(t, putObjErr)
    if getObjErr == nil {
        defer getObjectResponse.Content.Close()
        bs, readErr := ioutil.ReadAll(getObjectResponse.Content)
        ds3Testing.AssertNilError(t, readErr)
        if bytes.Compare(bs, book) != 0 {
            t.Error("Retrieved book does not match uploaded book.")
        }
    }
}

func TestPutGetFilePercentEncodingObjectNames(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    inputSymbols := []string {"-", ".", "_", "~", "!", "$", "'", "(", ")", "*", ",", "&", "=", "@", ":", "/", ";", "+", "?", " ", "%", "#", "'", "\t"}

    for i, input := range inputSymbols {
        objectName := fmt.Sprintf("test%ssymbol%d", input, i)

        func() {
            //Put object to BP
            content := []byte("hello world")
            defer testutils.DeleteObjectLogError(t, client, testBucket, objectName)

            putObjErr := testutils.PutObjectLogError(t, client, testBucket, objectName, content)
            ds3Testing.AssertNilError(t, putObjErr)

            //Verify that object exists
            getObjectResponse, getObjErr := testutils.GetObjectLogError(t, client, testBucket, objectName)
            ds3Testing.AssertNilError(t, getObjErr)
            if getObjErr == nil {
                defer getObjectResponse.Content.Close()
                retrievedContent, readErr := ioutil.ReadAll(getObjectResponse.Content)
                ds3Testing.AssertNilError(t, readErr)
                if bytes.Compare(retrievedContent, content) != 0 {
                    t.Error("Retrieved content does not match uploaded content.")
                }
            }
        }()
    }
}

func TestBadBucket(t *testing.T) {
    // Attempt to create a malformed bucket
    badBucketName := ""
    err := testutils.PutBucket(client, badBucketName)
    ds3Testing.AssertBadStatusCodeError(t, 400, err)
}

func TestBucketAlreadyExists(t *testing.T) {
    // Attempt to create a bucket that already exists
    err := testutils.PutBucket(client, testBucket)
    ds3Testing.AssertBadStatusCodeError(t, 409, err)
}

func TestDeleteEmptyBucket(t *testing.T) {
    bucketName := "GoDeleteEmptyBucket"
    putErr := testutils.PutBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, putErr)

    //Delete bucket
    deleteErr := testutils.DeleteBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, deleteErr)

    //Verify bucket does not exist
    _, getDeletedErr := testutils.GetBucket(client, bucketName)
    ds3Testing.AssertBadStatusCodeError(t, 404, getDeletedErr)
}

func TestDeleteBucketMalformedName(t *testing.T) {
    err := testutils.DeleteBucket(client, "")
    ds3Testing.AssertBadStatusCodeError(t, 400, err)
}

func TestDeleteBucketNonexistent(t *testing.T) {
    err := testutils.DeleteBucket(client, "not-here")
    ds3Testing.AssertBadStatusCodeError(t, 404, err)
}

func TestDeleteBucketNonEmpty(t *testing.T) {
    beowulf := "beowulf.txt"
    bucketName := "GoTestDeleteBucketNonEmpty"

    //Create test bucket
    err := testutils.PutBucket(client, bucketName)
    ds3Testing.AssertNilError(t, err)
    defer testutils.DeleteBucketLogError(t, client, bucketName)

    //Put object to test bucket
    book, bookErr := testutils.LoadBookLogError(t, beowulf)
    ds3Testing.AssertNilError(t, bookErr)
    defer testutils.DeleteObjectLogError(t, client, bucketName, beowulf)

    putObjErr := testutils.PutObjectLogError(t, client, bucketName, beowulf, book)
    ds3Testing.AssertNilError(t, putObjErr)

    //Attempt to delete non-empty bucket
    deleteErr := testutils.DeleteBucket(client, bucketName)
    ds3Testing.AssertBadStatusCodeError(t, 409, deleteErr)
}

func TestGetBucket(t *testing.T) {
    response, err := testutils.GetBucket(client, testBucket)
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertNonNilStringPtr(t, "Name", testBucket, response.ListBucketResult.Name)
}

func TestGetBucketNonexistent(t *testing.T) {
    badBucketName := "not-there"
    _, err := testutils.GetBucket(client, badBucketName)
    ds3Testing.AssertBadStatusCodeError(t, 404, err)
}

func TestHeadBucket(t *testing.T) {
    _, err := client.HeadBucket(models.NewHeadBucketRequest(testBucket))
    ds3Testing.AssertNilError(t, err)
}

func TestHeadBucketNonExistent(t *testing.T) {
    bucketName := "not-here"
    _, err := client.HeadBucket(models.NewHeadBucketRequest(bucketName))
    ds3Testing.AssertBadStatusCodeError(t, 404, err)
}

func TestGetBucketPagination(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    var ds3PutObjects []models.Ds3PutObject
    for i := 0; i < 15; i++ {
        name := "file" + strconv.Itoa(i + 10) // Start at 10 to avoid alphabetical reordering 0, 1, 10, 11, etc
        curObj := models.Ds3PutObject{Name:name, Size:0}
        ds3PutObjects = append(ds3PutObjects, curObj)
    }

    err := testutils.PutEmptyObjects(client, testBucket, ds3PutObjects)
    ds3Testing.AssertNilError(t, err)

    //Test files indexed 0-4
    result1, err := client.GetBucket(models.NewGetBucketRequest(testBucket).WithMaxKeys(5))
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertInt(t, "Number of Objects", 5, len(result1.ListBucketResult.Objects))
    if result1.ListBucketResult.NextMarker == nil {
        t.Fatal("Expected NextMarker to be non-nil value.")
    }
    for i, obj := range result1.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "ObjectName", ds3PutObjects[i].Name, obj.Key)
    }

    // Test files indexed 5-9
    result2, err := client.GetBucket(
        models.NewGetBucketRequest(testBucket).
            WithMaxKeys(5).
            WithMarker(*result1.ListBucketResult.NextMarker))

        ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertInt(t, "Number of Objects", 5, len(result2.ListBucketResult.Objects))
    if result2.ListBucketResult.NextMarker == nil {
        t.Fatal("Expected NextMarker to be non-nil value.")
    }
    for i, obj := range result2.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "ObjectName", ds3PutObjects[i + 5].Name, obj.Key)
    }

    // Test files indexed 10-14
    result3, err := client.GetBucket(
        models.NewGetBucketRequest(testBucket).
            WithMaxKeys(5).
            WithMarker(*result2.ListBucketResult.NextMarker))

    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertInt(t, "Number of Objects", 5, len(result3.ListBucketResult.Objects))
    if result2.ListBucketResult.NextMarker == nil {
        t.Fatal("Expected NextMarker to be non-nil value.")
    }
    for i, obj := range result3.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "ObjectName", ds3PutObjects[i + 10].Name, obj.Key)
    }
}

func TestGetBucketDelimiter(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    var ds3PutObjects []models.Ds3PutObject
    for i := 0; i < 10; i++ {
        name := "file" + strconv.Itoa(i)
        curObj := models.Ds3PutObject{Name:name, Size:0}
        ds3PutObjects = append(ds3PutObjects, curObj)
    }
    for i := 0; i < 10; i++ {
        name := "dir/file" + strconv.Itoa(i)
        curObj := models.Ds3PutObject{Name:name, Size:0}
        ds3PutObjects = append(ds3PutObjects, curObj)
    }

    err := testutils.PutEmptyObjects(client, testBucket, ds3PutObjects)
    ds3Testing.AssertNilError(t, err)

    //Test files indexed 0-4
    result, err := client.GetBucket(models.NewGetBucketRequest(testBucket).WithDelimiter("/"))
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertInt(t, "Number of Objects", 10, len(result.ListBucketResult.Objects))
    ds3Testing.AssertInt(t, "Number of Common Prefixes", 1, len(result.ListBucketResult.CommonPrefixes))
    ds3Testing.AssertString(t, "CommonPrefix", "dir/", result.ListBucketResult.CommonPrefixes[0])
}

func TestBulkPut(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    books, bookErr := testutils.GetResourceBooks()
    ds3Testing.AssertNilError(t, bookErr)

    ds3Objects, objErr := testutils.ConvertBooksToDs3Objects(books)
    ds3Testing.AssertNilError(t, objErr)

    // Create bulk put job
    bulkPut, bulkPutErr := client.PutBulkJobSpectraS3(models.NewPutBulkJobSpectraS3Request(testBucket, ds3Objects))
    ds3Testing.AssertNilError(t, bulkPutErr)

    for _, chunk := range bulkPut.MasterObjectList.Objects {
        allocateChunk, allocateErr := client.AllocateJobChunkSpectraS3(models.NewAllocateJobChunkSpectraS3Request(chunk.ChunkId))
        ds3Testing.AssertNilError(t, allocateErr)
        for _, obj := range allocateChunk.Objects.Objects {
            content := books[*obj.Name]
            _, putObjErr := client.PutObject(models.NewPutObjectRequest(testBucket, *obj.Name, content).
                WithOffset(obj.Offset).
                WithJob(bulkPut.MasterObjectList.JobId))

            ds3Testing.AssertNilError(t, putObjErr)
        }
    }

    // Verify books are in the bucket
    getBucket, getBucketErr := client.GetBucket(models.NewGetBucketRequest(testBucket))
    ds3Testing.AssertNilError(t, getBucketErr)
    if len(getBucket.ListBucketResult.Objects) != len(books) {
        t.Fatalf("Expected '%d' objects in bucket '%s', but found '%d'.", len(books), testBucket, len(getBucket.ListBucketResult.Objects))
    }
    for i, obj := range getBucket.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "BookName", testutils.BookTitles[i], obj.Key)
    }
}

func TestBulkGet(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    bookErr := testutils.PutTestBooks(client, testBucket)
    ds3Testing.AssertNilError(t, bookErr)

    // Create bulk get job
    bucketContents, bucketErr := client.GetBucket(models.NewGetBucketRequest(testBucket))
    ds3Testing.AssertNilError(t, bucketErr)
    ds3Testing.AssertInt(t, "Number of Objects on BP", 4, len(bucketContents.ListBucketResult.Objects))

    objectNames := testutils.ConvertObjectsIntoObjectNameList(bucketContents.ListBucketResult.Objects)
    bulkGet, bulkGetErr := client.GetBulkJobSpectraS3(models.NewGetBulkJobSpectraS3Request(testBucket, objectNames))
    ds3Testing.AssertNilError(t, bulkGetErr)

    availableChunks, chunksErr := client.GetJobChunkSpectraS3(models.NewGetJobChunkSpectraS3Request(bulkGet.MasterObjectList.JobId))
    ds3Testing.AssertNilError(t, chunksErr)

    // Get all objects and verify content
    for _, obj := range availableChunks.Objects.Objects {
        func() {
            getObj, objErr := client.GetObject(models.NewGetObjectRequest(testBucket, *obj.Name))
            ds3Testing.AssertNilError(t, objErr)

            defer getObj.Content.Close()
            testutils.VerifyBookContent(t, *obj.Name, getObj.Content)
        }()
    }
}

func TestGetUsersSpectraS3WithName(t *testing.T) {
    user, err := testutils.GetUserByNameLogError(t, client, defaultUser)
    ds3Testing.AssertNilError(t, err)
    ds3Testing.AssertNonNilStringPtr(t, "Name", defaultUser, user.Name)
}

func TestDataPolicySpectraS3(t *testing.T) {
    dataPolicyName := "GoTestDataPolicy"

    // Create data policy and verify creation
    putResponse, putErr := testutils.PutDataPolicyLogError(t, client, dataPolicyName)
    ds3Testing.AssertNilError(t, putErr)

    defer testutils.DeleteDataPolicyLogError(t, client, dataPolicyName)

    getResponse, getErr := testutils.GetDataPolicyLogError(t, client, putResponse.DataPolicy.Id)
    ds3Testing.AssertNilError(t, getErr)
    ds3Testing.AssertNonNilStringPtr(t, "Data Policy Name", dataPolicyName, getResponse.DataPolicy.Name)
}

func TestSettingDefaultDataPolicy(t *testing.T) {
    dataPolicyName := "GoTestSettingDefaultDataPolicy"

    // Create data policy
    putDataPolicyResponse, putErr := testutils.PutDataPolicyLogError(t, client, dataPolicyName)
    ds3Testing.AssertNilError(t, putErr)

    defer testutils.DeleteDataPolicyLogError(t, client, dataPolicyName)

    // Get default user and make note of default data policy if one is set
    user, userErr := testutils.GetUserByNameLogError(t, client, defaultUser)
    ds3Testing.AssertNilError(t, userErr)
    curDataPolicy := user.DefaultDataPolicyId

    // Modify user's default data policy and verify
    getDataPolicyResponse, getErr := testutils.GetDataPolicyLogError(t, client, dataPolicyName)
    ds3Testing.AssertNilError(t, getErr)

    modifyResponse, modifyErr := client.ModifyUserSpectraS3(models.NewModifyUserSpectraS3Request(user.Id).
        WithDefaultDataPolicyId(getDataPolicyResponse.DataPolicy.Id))
    ds3Testing.AssertNilError(t, modifyErr)

    defer client.ModifyUserSpectraS3(models.NewModifyUserSpectraS3Request(user.Id).
        WithDefaultDataPolicyId(*curDataPolicy))

    ds3Testing.AssertNonNilStringPtr(t, "Name", defaultUser, modifyResponse.SpectraUser.Name)
    ds3Testing.AssertNonNilStringPtr(t, "DataPolicyId", putDataPolicyResponse.DataPolicy.Id, modifyResponse.SpectraUser.DefaultDataPolicyId)
}

func TestStorageDomain(t *testing.T) {
    storageDomainName := "GoTestStorageDomain"

    // Create storage domain
    putResponse, putErr := client.PutStorageDomainSpectraS3(models.NewPutStorageDomainSpectraS3Request(storageDomainName))
    ds3Testing.AssertNilError(t, putErr)

    defer client.DeleteStorageDomainSpectraS3(models.NewDeleteStorageDomainSpectraS3Request(storageDomainName))

    ds3Testing.AssertNonNilStringPtr(t, "Name", storageDomainName, putResponse.StorageDomain.Name)
}

func TestPoolPartition(t *testing.T) {
    poolPartitionName := "GoTestPoolPartition"

    // Create pool partition
    putResponse, putErr := client.PutPoolPartitionSpectraS3(models.NewPutPoolPartitionSpectraS3Request(poolPartitionName, models.POOL_TYPE_ONLINE))
    ds3Testing.AssertNilError(t, putErr)

    defer client.DeletePoolPartitionSpectraS3(models.NewDeletePoolPartitionSpectraS3Request(poolPartitionName))

    ds3Testing.AssertNonNilStringPtr(t, "Name", poolPartitionName, putResponse.PoolPartition.Name)
    if putResponse.PoolPartition.Type != models.POOL_TYPE_ONLINE {
        t.Fatalf("Expected pool partition type of ONLINE but was '%s'.", putResponse.PoolPartition.Type.String())
    }
}

func TestStorageDomainMember(t *testing.T) {
    varTestName := "GoTestStorageDomainMember"

    // Create storage domain
    storageDomainResponse, storageDomainErr := client.PutStorageDomainSpectraS3(models.NewPutStorageDomainSpectraS3Request(varTestName))
    ds3Testing.AssertNilError(t, storageDomainErr)

    defer client.DeleteStorageDomainSpectraS3(models.NewDeleteStorageDomainSpectraS3Request(varTestName))

    // Create pool partition
    poolPartitionResponse, poolPartitionErr := client.PutPoolPartitionSpectraS3(models.NewPutPoolPartitionSpectraS3Request(varTestName, models.POOL_TYPE_ONLINE))
    ds3Testing.AssertNilError(t, poolPartitionErr)

    defer client.DeletePoolPartitionSpectraS3(models.NewDeletePoolPartitionSpectraS3Request(varTestName))

    // Create storage domain member linking pool partition to storage domain
    response, err := client.PutPoolStorageDomainMemberSpectraS3(models.NewPutPoolStorageDomainMemberSpectraS3Request(
        poolPartitionResponse.PoolPartition.Id,
        storageDomainResponse.StorageDomain.Id))
    ds3Testing.AssertNilError(t, err)

    defer client.DeleteStorageDomainMemberSpectraS3(models.NewDeleteStorageDomainMemberSpectraS3Request(response.StorageDomainMember.Id))

    ds3Testing.AssertNonNilStringPtr(t, "PoolPartitionId", poolPartitionResponse.PoolPartition.Id, response.StorageDomainMember.PoolPartitionId)
    ds3Testing.AssertString(t, "StorageDomainId", storageDomainResponse.StorageDomain.Id, response.StorageDomainMember.StorageDomainId)
}

func TestDataPersistenceRule(t *testing.T) {
    varTestName := "GoTestDataPersistenceRule"
    dataPersistenceRuleType := models.DATA_PERSISTENCE_RULE_TYPE_PERMANENT
    dataIsolationLevel := models.DATA_ISOLATION_LEVEL_STANDARD

    // Create data policy
    putDataPolicyResponse, putErr := testutils.PutDataPolicyLogError(t, client, varTestName)
    ds3Testing.AssertNilError(t, putErr)

    defer testutils.DeleteDataPolicyLogError(t, client, varTestName)

    // Create storage domain
    storageDomainResponse, storageDomainErr := client.PutStorageDomainSpectraS3(models.NewPutStorageDomainSpectraS3Request(varTestName))
    ds3Testing.AssertNilError(t, storageDomainErr)

    defer client.DeleteStorageDomainSpectraS3(models.NewDeleteStorageDomainSpectraS3Request(varTestName))

    // Create pool partition
    poolPartitionResponse, poolPartitionErr := client.PutPoolPartitionSpectraS3(models.NewPutPoolPartitionSpectraS3Request(varTestName, models.POOL_TYPE_ONLINE))
    ds3Testing.AssertNilError(t, poolPartitionErr)

    defer client.DeletePoolPartitionSpectraS3(models.NewDeletePoolPartitionSpectraS3Request(varTestName))

    // Create storage domain member linking pool partition to storage domain
    memberResponse, memberErr := client.PutPoolStorageDomainMemberSpectraS3(models.NewPutPoolStorageDomainMemberSpectraS3Request(
        poolPartitionResponse.PoolPartition.Id,
        storageDomainResponse.StorageDomain.Id))
    ds3Testing.AssertNilError(t, memberErr)

    defer client.DeleteStorageDomainMemberSpectraS3(models.NewDeleteStorageDomainMemberSpectraS3Request(memberResponse.StorageDomainMember.Id))

    // Create data persistence rule linking data policy and storage domain
    response, err := client.PutDataPersistenceRuleSpectraS3(models.NewPutDataPersistenceRuleSpectraS3Request(
        dataPersistenceRuleType,
        putDataPolicyResponse.DataPolicy.Id,
        dataIsolationLevel,
        storageDomainResponse.StorageDomain.Id))
    ds3Testing.AssertNilError(t, err)

    defer client.DeleteDataPersistenceRuleSpectraS3(models.NewDeleteDataPersistenceRuleSpectraS3Request(response.DataPersistenceRule.Id))

    ds3Testing.AssertString(t, "DataPolicyId", putDataPolicyResponse.DataPolicy.Id, response.DataPersistenceRule.DataPolicyId)
    ds3Testing.AssertString(t, "StorageDomainId", storageDomainResponse.StorageDomain.Id, response.DataPersistenceRule.StorageDomainId)
    if response.DataPersistenceRule.Type != dataPersistenceRuleType {
        t.Fatalf("Expected DataPersistenceRuleType to be '%s' but was '%s'.", dataPersistenceRuleType.String(), response.DataPersistenceRule.Type.String())
    }
    if response.DataPersistenceRule.IsolationLevel != dataIsolationLevel {
        t.Fatalf("Expected DataIsolationLevel to be '%s' but was '%s'.", dataIsolationLevel.String(), response.DataPersistenceRule.IsolationLevel.String())
    }
}

func TestPuttingFolder(t *testing.T) {
    bucketName := "GoTestPuttingFolder"
    err := testutils.PutBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, err)

    defer deleteBucketAndContent(t, bucketName)

    const folderPath = "Gracie/Eskimo/"

    readSizer := newNilReadSizer(nil, 0)
    putObjectRequest := models.NewPutObjectRequest(bucketName, folderPath, &readSizer)
    _, err = client.PutObject(putObjectRequest)
    ds3Testing.AssertNilError(t, err)

    getBucketRequest := models.NewGetBucketRequest(bucketName)
    getBucketResponse, err := client.GetBucket(getBucketRequest)
    ds3Testing.AssertNilError(t, err)

    ds3Testing.AssertInt(t, "Number of objects in bucket", 1, len(getBucketResponse.ListBucketResult.Objects))
    ds3Testing.AssertString(t, "Folder names equal", folderPath, *getBucketResponse.ListBucketResult.Objects[0].Key)
}

func deleteBucketAndContent(t *testing.T, bucketName string) {
    testutils.DeleteBucketContents(client, bucketName)
    testutils.DeleteBucketLogError(t, client, bucketName)
}

func TestPuttingZeroLengthObject(t *testing.T) {
    bucketName := "GoTestPuttingZeroLengthObject"
    err := testutils.PutBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, err)

    defer deleteBucketAndContent(t, bucketName)

    const objectName = "Gracie"

    zeroBytes := make([]byte, 0)

    putObjectRequest := models.NewPutObjectRequest(bucketName, objectName, ds3.BuildByteReaderWithSizeDecorator(zeroBytes))

    _, err = client.PutObject(putObjectRequest)

    ds3Testing.AssertNilError(t, err)

    getBucketRequest := models.NewGetBucketRequest(bucketName)
    getBucketResponse, err := client.GetBucket(getBucketRequest)
    ds3Testing.AssertNilError(t, err)

    ds3Testing.AssertInt(t, "Number of objects in bucket", 1, len(getBucketResponse.ListBucketResult.Objects))
    ds3Testing.AssertString(t, "Object names equal", objectName, *getBucketResponse.ListBucketResult.Objects[0].Key)
    ds3Testing.AssertInt64(t, "Object size equal", 0, getBucketResponse.ListBucketResult.Objects[0].Size)
}

func TestDeleteObjects(t *testing.T) {
    bucketName := "GoTestDeleteObjects"

    // create bucket
    err := testutils.PutBucketLogError(t, client, bucketName)
    ds3Testing.AssertNilError(t, err)
    defer deleteBucketAndContent(t, bucketName)

    // put the test books into the bucket
    testutils.PutTestBooks(client, bucketName)

    // list all objects in the bucket
    getBucket, err := client.GetBucket(models.NewGetBucketRequest(bucketName))

    ds3Testing.AssertInt(t, "number of objects to delete", 4, len(getBucket.ListBucketResult.Objects))

    names := make([]string, 0)
    for _, obj := range getBucket.ListBucketResult.Objects {
        names = append(names, *obj.Key)
    }

    deleteObjs := models.NewDeleteObjectsRequest(bucketName, names)
    _, err = client.DeleteObjects(deleteObjs)
    ds3Testing.AssertNilError(t, err)
}