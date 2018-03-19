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

package ds3

import (
    "testing"
    "strconv"
    "net/url"
    "net/http"
    "io/ioutil"
    "spectra/ds3_go_sdk/ds3/models"
    "reflect"
    "spectra/ds3_go_sdk/ds3_utils/ds3Testing"
)

func TestGetService(t *testing.T) {
    stringResponse := "<ListAllMyBucketsResult xmlns=\"http://doc.s3.amazonaws.com/2006-03-01\"><Owner><ID>ryan</ID><DisplayName>ryan</DisplayName></Owner><Buckets><Bucket><Name>testBucket2</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest1</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest2</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest3</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest4</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest5</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest6</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>testBucket3</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>testBucket1</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>testbucket</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket></Buckets></ListAllMyBucketsResult>"
    expectedBucketNames := []string {
        "testBucket2",
        "bulkTest",
        "bulkTest1",
        "bulkTest2",
        "bulkTest3",
        "bulkTest4",
        "bulkTest5",
        "bulkTest6",
        "testBucket3",
        "testBucket1",
        "testbucket",
    }
    expectedCreationDates := []string {
        "2013-12-11T23:20:09",
        "2013-12-11T23:20:09",
        "2013-12-11T23:20:09",
        "2013-12-11T23:20:09",
        "2013-12-11T23:20:09",
        "2013-12-11T23:20:09",
        "2013-12-11T23:20:09",
        "2013-12-11T23:20:09",
        "2013-12-11T23:20:09",
        "2013-12-11T23:20:09",
        "2013-12-11T23:20:09",
    }

    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/", &url.Values{}, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        GetService(models.NewGetServiceRequest())

    // Validate the error contents.
    ds3Testing.AssertNilError(t, err)

    // Check the response.
    if response == nil {
        t.Fatalf("Received an unexpected nil response.")
    }

    ds3Testing.AssertString(t, "Id", "ryan", response.ListAllMyBucketsResult.Owner.Id)
    ds3Testing.AssertNonNilStringPtr(t, "DisplayName", "ryan", response.ListAllMyBucketsResult.Owner.DisplayName)
    if len(response.ListAllMyBucketsResult.Buckets) != len(expectedBucketNames) {
        t.Fatalf("Parsed an unexpected number (%d) of buckets.", len(response.ListAllMyBucketsResult.Buckets))
    }
    for i, bucket := range response.ListAllMyBucketsResult.Buckets {
        ds3Testing.AssertNonNilStringPtr(t, "Name", expectedBucketNames[i], bucket.Name)
        ds3Testing.AssertNonNilStringPtr(t, "CreationDate", expectedCreationDates[i], bucket.CreationDate)
    }
}

func TestGetBadService(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/", &url.Values{}, &http.Header{}, nil).
        Returning(400, "", nil).
        GetService(models.NewGetServiceRequest())

    // Check the response.
    if response != nil {
        t.Fatal("Expected a nil response but received a value.")
    }

    // Validate the error contents.
    if err == nil {
        t.Fatal("Expected an error but didn't get one.")
    }

    expectedError := "Received a status code of 400 when [200] was expected. Could not parse the response for additional information."
    ds3Testing.AssertString(t, "Error", expectedError, err.Error())
}

func TestGetBucket(t *testing.T) {
    stringResponse := "<ListBucketResult xmlns=\"http://s3.amazonaws.com/doc/2006-03-01/\"><Name>remoteTest16</Name><Prefix/><Marker/><MaxKeys>1000</MaxKeys><IsTruncated>false</IsTruncated><Contents><Key>user/hduser/gutenberg/20417.txt.utf-8</Key><LastModified>2014-01-03T13:26:47.000Z</LastModified><ETag>NOTRETURNED</ETag><Size>674570</Size><StorageClass>STANDARD</StorageClass><Owner><ID>ryan</ID><DisplayName>ryan</DisplayName></Owner></Contents><Contents><Key>user/hduser/gutenberg/5000.txt.utf-8</Key><LastModified>2014-01-03T13:26:47.000Z</LastModified><ETag>NOTRETURNED</ETag><Size>1423803</Size><StorageClass>STANDARD</StorageClass><Owner><ID>ryan</ID><DisplayName>ryan</DisplayName></Owner></Contents><Contents><Key>user/hduser/gutenberg/4300.txt.utf-8</Key><LastModified>2014-01-03T13:26:47.000Z</LastModified><ETag>NOTRETURNED</ETag><Size>1573150</Size><StorageClass>STANDARD</StorageClass><Owner><ID>ryan</ID><DisplayName>ryan</DisplayName></Owner></Contents></ListBucketResult>"
    keys := []string {
        "user/hduser/gutenberg/20417.txt.utf-8",
        "user/hduser/gutenberg/5000.txt.utf-8",
        "user/hduser/gutenberg/4300.txt.utf-8",
    }
    lastModifieds := []string {
        "2014-01-03T13:26:47.000Z",
        "2014-01-03T13:26:47.000Z",
        "2014-01-03T13:26:47.000Z",
    }
    etags := []string { "NOTRETURNED", "NOTRETURNED", "NOTRETURNED" }
    sizes := []int64 { 674570, 1423803, 1573150 }
    storageClasses := []string { "STANDARD", "STANDARD", "STANDARD" }
    ids := []string { "ryan", "ryan", "ryan" }
    displayNames := ids

    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/remoteTest16", &url.Values{}, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        GetBucket(models.NewGetBucketRequest("remoteTest16"))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    if len(response.ListBucketResult.Objects) != len(keys) {
        t.Fatalf("Expected %d objects but got %d.", len(keys), len(response.ListBucketResult.Objects))
    }
    ds3Testing.AssertNonNilStringPtr(t, "Name", "remoteTest16", response.ListBucketResult.Name)
    ds3Testing.AssertStringPtrIsNil(t, "Prefix", response.ListBucketResult.Prefix)
    ds3Testing.AssertStringPtrIsNil(t, "Marker", response.ListBucketResult.Marker)
    ds3Testing.AssertInt(t, "MaxKeys", 1000, response.ListBucketResult.MaxKeys)
    ds3Testing.AssertBool(t, "Truncated", false, response.ListBucketResult.Truncated)
    for i, object := range response.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "Key", keys[i], object.Key)
        ds3Testing.AssertNonNilStringPtr(t, "LastModified", lastModifieds[i], object.LastModified)
        ds3Testing.AssertNonNilStringPtr(t, "ETag", etags[i], object.ETag)
        ds3Testing.AssertInt64(t, "Size", sizes[i], object.Size)
        ds3Testing.AssertNonNilStringPtr(t, "StorageClass", storageClasses[i], object.StorageClass)
        ds3Testing.AssertString(t, "Id", ids[i], object.Owner.Id)
        ds3Testing.AssertNonNilStringPtr(t, "DisplayName", displayNames[i], object.Owner.DisplayName)
    }
}

func TestGetBucketWithCommonPrefixes(t *testing.T) {
    stringResponse := "<ListBucketResult><CommonPrefixes><Prefix>movies/</Prefix></CommonPrefixes><CommonPrefixes><Prefix>scores/</Prefix></CommonPrefixes><Contents><ETag/><Key>music.mp3</Key><LastModified/><Owner><DisplayName>jason</DisplayName><ID>a0b789ef-dcbd-4004-82af-6679a26329bf</ID></Owner><Size>0</Size><StorageClass/></Contents><Contents><ETag/><Key>song.mp3</Key><LastModified/><Owner><DisplayName>jason</DisplayName><ID>a0b789ef-dcbd-4004-82af-6679a26329bf</ID></Owner><Size>0</Size><StorageClass/></Contents><CreationDate>2017-03-23T23:26:34.000Z</CreationDate><Delimiter>/</Delimiter><IsTruncated>false</IsTruncated><Marker/><MaxKeys>1000</MaxKeys><Name>bucket1</Name><NextMarker/><Prefix/></ListBucketResult>"
    keys := []string {
        "music.mp3",
        "song.mp3",
    }

    ids := []string {
        "a0b789ef-dcbd-4004-82af-6679a26329bf",
        "a0b789ef-dcbd-4004-82af-6679a26329bf",
    }
    expectedPrefixes := []string {
        "movies/",
        "scores/",
    }

    displayName := "jason"

    // Create and run the mocked client.
    bucketName := "bucket1"
    delimiter := "/"
    queryParams := &url.Values{"delimiter": []string{delimiter}}
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/" + bucketName, queryParams, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        GetBucket(models.NewGetBucketRequest(bucketName).WithDelimiter(delimiter))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    if len(response.ListBucketResult.CommonPrefixes) != len(expectedPrefixes) {
        t.Fatalf("Expected %d common prefixes but got %d.", len(expectedPrefixes), len(response.ListBucketResult.CommonPrefixes))
    }
    for i, prefix := range response.ListBucketResult.CommonPrefixes {
        ds3Testing.AssertString(t, "CommonPrefix", expectedPrefixes[i], prefix)
    }

    if len(response.ListBucketResult.Objects) != len(keys) {
        t.Fatalf("Expected %d objects but got %d.", len(keys), len(response.ListBucketResult.Objects))
    }
    ds3Testing.AssertNonNilStringPtr(t, "Name", bucketName, response.ListBucketResult.Name)
    ds3Testing.AssertStringPtrIsNil(t, "Prefix", response.ListBucketResult.Prefix)
    ds3Testing.AssertStringPtrIsNil(t, "Marker", response.ListBucketResult.Marker)
    ds3Testing.AssertInt(t, "MaxKeys", 1000, response.ListBucketResult.MaxKeys)
    ds3Testing.AssertBool(t, "Truncated", false, response.ListBucketResult.Truncated)
    for i, object := range response.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "Key", keys[i], object.Key)
        ds3Testing.AssertStringPtrIsNil(t, "LastModified", object.LastModified)
        ds3Testing.AssertStringPtrIsNil(t, "ETag", object.ETag)
        ds3Testing.AssertInt64(t, "Size", 0, object.Size)
        ds3Testing.AssertStringPtrIsNil(t, "StorageClass", object.StorageClass)
        ds3Testing.AssertString(t, "Id", ids[i], object.Owner.Id)
        ds3Testing.AssertNonNilStringPtr(t, "DisplayName", displayName, object.Owner.DisplayName)
    }
}

func TestPutBucket(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_PUT, "/bucketName", &url.Values{}, &http.Header{}, nil).
        Returning(200, "", nil).
        PutBucket(models.NewPutBucketRequest("bucketName"))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestDeleteBucket(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_DELETE, "/bucketName", &url.Values{}, &http.Header{}, nil).
        Returning(204, "", nil).
        DeleteBucket(models.NewDeleteBucketRequest("bucketName"))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestDeleteFolderRecursivelySpectraS3(t *testing.T) {
    bucketId := "a24d14f3-e2f0-4bfb-ab71-f99d5ef43745"
    folderName := "FolderName"
    queryParams := &url.Values{"bucket_id": []string{bucketId}, "recursive": []string{""}}

    // Create and run the mocked client.
    response, err := mockedClient(t).
            Expecting(HTTP_VERB_DELETE, "/_rest_/folder/FolderName", queryParams, &http.Header{}, nil).
            Returning(204, "", nil).
            DeleteFolderRecursivelySpectraS3(models.NewDeleteFolderRecursivelySpectraS3Request(bucketId, folderName))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestDeleteObject(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_DELETE, "/bucketName/my/file.txt", &url.Values{}, &http.Header{}, nil).
        Returning(204, "", nil).
        DeleteObject(models.NewDeleteObjectRequest("bucketName", "my/file.txt"))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestGetBadBucket(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/remoteTest16", &url.Values{}, &http.Header{}, nil).
        Returning(400, "", nil).
        GetBucket(models.NewGetBucketRequest("remoteTest16"))

    // Check the error result.
    if err == nil {
        t.Fatal("Expected an error but got nil.")
    }

    expectedError := "Received a status code of 400 when [200] was expected. Could not parse the response for additional information."
    ds3Testing.AssertString(t, "Error", expectedError, err.Error())

    // Check the response value.
    if response != nil {
        t.Fatal("Response was supposed to be nil.")
    }
}

func TestGetObject(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/bucketName/object", &url.Values{}, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        GetObject(models.NewGetObjectRequest("bucketName", "object"))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    if response.Content == nil {
        t.Fatal("Response content was unexpectedly nil.")
    }
    defer response.Content.Close()
    bs, readErr := ioutil.ReadAll(response.Content)
    ds3Testing.AssertNilError(t, readErr)
    ds3Testing.AssertString(t, "Response", stringResponse, string(bs))
}

func TestGetPartialObject(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    requestHeaders := &http.Header{}
    requestHeaders.Add("Range", "bytes=1-5,7-8,20-500")

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/bucketName/object", &url.Values{}, requestHeaders, nil).
        Returning(206, stringResponse, nil).
        GetObject(models.NewGetObjectRequest("bucketName", "object").
            WithRanges(
                models.Range{Start: 1, End: 5},
                models.Range{Start: 7, End: 8},
                models.Range{Start: 20, End: 500}))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    if response.Content == nil {
        t.Fatal("Response content was unexpectedly nil.")
    }
    defer response.Content.Close()
    bs, readErr := ioutil.ReadAll(response.Content)
    ds3Testing.AssertNilError(t, readErr)
    ds3Testing.AssertString(t, "Response", stringResponse, string(bs))
}

func TestGetObjectRange(t *testing.T) {
    stringResponse := "object contents"
    requestHeaders := &http.Header{"Range": []string{"bytes=20-179"}}

    // Create and run the mocked client.
    request := models.NewGetObjectRequest("bucketName", "object").
            WithRanges(models.Range{Start: 20, End: 179})

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/bucketName/object", &url.Values{}, requestHeaders, nil).
        Returning(200, stringResponse, &http.Header{"Range": []string{"bytes=20-179"}}).
        GetObject(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    if response.Content == nil {
        t.Fatal("Response content was unexpectedly nil.")
    }
    defer response.Content.Close()
    bs, readErr := ioutil.ReadAll(response.Content)
    ds3Testing.AssertNilError(t, readErr)
    ds3Testing.AssertString(t, "Response", stringResponse, string(bs))
}

func TestGetObjectsDetailsSpectraS3(t *testing.T) {
    bucketId := "a24d14f3-e2f0-4bfb-ab71-f99d5ef43745"
    stringResponse := "<Data><S3Object><BucketId>a24d14f3-e2f0-4bfb-ab71-f99d5ef43745</BucketId><CreationDate>2015-09-21T20:06:47.694Z</CreationDate><Id>e37c3ce0-12aa-4f54-87e3-42532aca0e5e</Id><Name>beowulf.txt</Name><Type>DATA</Type><Version>1</Version></S3Object>" +
            "<S3Object><BucketId>a24d14f3-e2f0-4bfb-ab71-f99d5ef43745</BucketId><CreationDate>2015-09-21T20:06:47.779Z</CreationDate><Id>dc628815-c723-4c4e-b68b-5f5d10f38af5</Id><Name>sherlock_holmes.txt</Name><Type>DATA</Type><Version>1</Version></S3Object>" +
            "<S3Object><BucketId>a24d14f3-e2f0-4bfb-ab71-f99d5ef43745</BucketId><CreationDate>2015-09-21T20:06:47.772Z</CreationDate><Id>4f6985fd-fbae-4421-ba27-66fdb96187c5</Id><Name>tale_of_two_cities.txt</Name><Type>DATA</Type><Version>1</Version></S3Object>" +
            "<S3Object><BucketId>a24d14f3-e2f0-4bfb-ab71-f99d5ef43745</BucketId><CreationDate>2015-09-21T20:06:47.696Z</CreationDate><Id>82c18910-fadb-4461-a152-bf714ae91b55</Id><Name>ulysses.txt</Name><Type>DATA</Type><Version>1</Version></S3Object></Data>"

    request := models.NewGetObjectsDetailsSpectraS3Request().WithBucketId(bucketId)

    responseHeaders := &http.Header{"page-truncated": []string{"0"}, "total-result-count": []string{"4"}}
    queryParams := &url.Values{"bucket_id": []string{bucketId}}

    response, err := mockedClient(t).
            Expecting(HTTP_VERB_GET, "/_rest_/object", queryParams, &http.Header{}, nil).
            Returning(200, stringResponse, responseHeaders).
            GetObjectsDetailsSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    if response.S3ObjectList.S3Objects == nil {
        t.Fatal("S3Object list was unexpectedly nil.")
    }

    expectedNames := []string{"beowulf.txt", "sherlock_holmes.txt", "tale_of_two_cities.txt", "ulysses.txt"}
    expectedCreationDates := []string{"2015-09-21T20:06:47.694Z", "2015-09-21T20:06:47.779Z", "2015-09-21T20:06:47.772Z", "2015-09-21T20:06:47.696Z"}
    expectedIds := []string{"e37c3ce0-12aa-4f54-87e3-42532aca0e5e", "dc628815-c723-4c4e-b68b-5f5d10f38af5", "4f6985fd-fbae-4421-ba27-66fdb96187c5", "82c18910-fadb-4461-a152-bf714ae91b55"}

    if len(response.S3ObjectList.S3Objects) != len(expectedNames) {
        t.Fatalf("Expected '%d' S3Objects, but got '%d'", len(expectedNames), len(response.S3ObjectList.S3Objects))
    }

    for i, object := range response.S3ObjectList.S3Objects {
        ds3Testing.AssertString(t, "BucketId", bucketId, object.BucketId)
        ds3Testing.AssertNonNilStringPtr(t, "CreationDate", expectedCreationDates[i], object.CreationDate)
        ds3Testing.AssertString(t, "Id", expectedIds[i], object.Id)
        ds3Testing.AssertNonNilStringPtr(t, "Name", expectedNames[i], object.Name)
        if object.Type != models.S3_OBJECT_TYPE_DATA {
            t.Fatalf("Expected type of '%d' but was '%d'.", models.S3_OBJECT_TYPE_DATA, object.Type)
        }
        ds3Testing.AssertInt64(t, "Version", 1, object.Version)
    }
}

func TestGetJobToReplicateSpectraS3(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    request := models.NewGetJobToReplicateSpectraS3Request("23a876ec-2fac-4dc8-b8e6-98d6026e7f4a")

    expectedParams := &url.Values{"replicate": []string{""}}
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/_rest_/job/23a876ec-2fac-4dc8-b8e6-98d6026e7f4a", expectedParams, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        GetJobToReplicateSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    ds3Testing.AssertString(t, "Content", stringResponse, response.Content)
}

func TestPutObject(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_PUT, "/bucketName/object", &url.Values{}, &http.Header{}, nil).
        Returning(200, "", nil).
        PutObject(models.NewPutObjectRequest(
            "bucketName",
            "object",
            BuildByteReaderWithSizeDecorator([]byte(stringResponse)),
        ))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestPutObjectWithMetaData(t *testing.T) {
    stringResponse := "object contents"

    expectedMetaData := &http.Header{}
    expectedMetaData.Add("x-amz-meta-test1", "test1value")
    expectedMetaData.Add("x-amz-meta-test2", "test2value,test2value2")

    request := models.NewPutObjectRequest(
        "bucketName",
        "object",
        BuildByteReaderWithSizeDecorator([]byte(stringResponse))).
            WithMetaData("x-amz-meta-test1", "test1value").
            WithMetaData("test2", "test2value", "test2value2")

    if len(request.Metadata) < 1 {
        t.Fatal("Expected there to be metadata")
    }

    // Create and run the mocked client.
    response, err := mockedClient(t).
            Expecting(HTTP_VERB_PUT, "/bucketName/object", &url.Values{}, expectedMetaData, nil).
            Returning(200, "", nil).
            PutObject(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestBulkPut(t *testing.T) {
    runBulkPutTest(
        t,
        "start_bulk_put",
        func(client *Client, objects []models.Ds3PutObject) ([]models.Objects, error) {
            request, err := client.PutBulkJobSpectraS3(models.NewPutBulkJobSpectraS3Request("bucketName", objects))
            return request.MasterObjectList.Objects, err
        },
    )
}

type bulkPutTest func(*Client, []models.Ds3PutObject) ([]models.Objects, error)

func runBulkPutTest(t *testing.T, operation string, callToTest bulkPutTest) {
    keys := []string { "file2", "file1", "file3" }
    sizes := []int64 { 1202, 256, 2523 }

    stringRequest := "<Objects><Object Name=\"file1\" Size=\"256\"></Object><Object Name=\"file2\" Size=\"1202\"></Object><Object Name=\"file3\" Size=\"2523\"></Object></Objects>"
    stringResponse := "<MasterObjectList><Objects><Object Name='file2' Length='1202'/><Object Name='file1' Length='256'/><Object Name='file3' Length='2523'/></Objects></MasterObjectList>"

    inputObjects := []models.Ds3PutObject {
        {Name: "file1", Size: 256 },
        {Name: "file2", Size: 1202 },
        {Name: "file3", Size: 2523 },
    }

    // Create and run the mocked client.
    client := mockedClient(t).
        Expecting(
            HTTP_VERB_PUT,
            "/_rest_/bucket/bucketName",
            &url.Values{"operation": []string{operation}},
            &http.Header{},
            &stringRequest,
        ).
        Returning(200, stringResponse, nil)
    response, err := callToTest(client, inputObjects)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
    if len(response) != 1 {
        t.Fatalf("Expected 1 object list but got %d.", len(response))
    }
    if len(response[0].Objects) != len(keys) {
        t.Fatalf("Expected %d objects but got %d.", len(keys), len(response[0].Objects))
    }
    for i, obj := range response[0].Objects {
        ds3Testing.AssertNonNilStringPtr(t, "Name", keys[i], obj.Name)
        ds3Testing.AssertInt64(t, "Length", sizes[i], obj.Length)
    }
}

func TestBulkGetWithSimpleDs3GetObjets(t *testing.T) {
    objects := []models.Ds3GetObject {
        models.NewDs3GetObject("file1"),
        models.NewDs3GetObject("file2"),
        models.NewDs3GetObject("file3"),
    }

    stringRequest := "<Objects><Object Name=\"file1\"></Object><Object Name=\"file2\"></Object><Object Name=\"file3\"></Object></Objects>"

    runBulkGetTest(
        t,
        "start_bulk_get",
        &stringRequest,
        func(client *Client) ([]models.Objects, error) {
            request, err := client.GetBulkJobSpectraS3(models.NewGetBulkJobSpectraS3RequestWithPartialObjects("bucketName", objects))
            return request.MasterObjectList.Objects, err
        },
    )
}

func TestBulkGetWithPartialDs3GetObjects(t *testing.T) {
    objects := []models.Ds3GetObject {
        models.NewPartialDs3GetObject("file1", 10, 100),
        models.NewPartialDs3GetObject("file2", 20, 200),
        models.NewPartialDs3GetObject("file3", 30, 300),
    }

    stringRequest := "<Objects><Object Name=\"file1\" Length=\"10\" Offset=\"100\"></Object><Object Name=\"file2\" Length=\"20\" Offset=\"200\"></Object><Object Name=\"file3\" Length=\"30\" Offset=\"300\"></Object></Objects>"

    runBulkGetTest(
        t,
        "start_bulk_get",
        &stringRequest,
        func(client *Client) ([]models.Objects, error) {
            request, err := client.GetBulkJobSpectraS3(models.NewGetBulkJobSpectraS3RequestWithPartialObjects("bucketName", objects))
            return request.MasterObjectList.Objects, err
        },
    )
}

func TestBulkGetWithObjectNames(t *testing.T) {
    objects := []string {"file1", "file2", "file3"}

    stringRequest := "<Objects><Object Name=\"file1\"></Object><Object Name=\"file2\"></Object><Object Name=\"file3\"></Object></Objects>"

    runBulkGetTest(
        t,
        "start_bulk_get",
        &stringRequest,
        func(client *Client) ([]models.Objects, error) {
            request, err := client.GetBulkJobSpectraS3(models.NewGetBulkJobSpectraS3Request("bucketName", objects))
            return request.MasterObjectList.Objects, err
        },
    )
}

func TestBulkVerifyWithPartialDs3GetObjects(t *testing.T) {
    objects := []models.Ds3GetObject {
        models.NewPartialDs3GetObject("file1", 10, 100),
        models.NewPartialDs3GetObject("file2", 20, 200),
        models.NewPartialDs3GetObject("file3", 30, 300),
    }

    stringRequest := "<Objects><Object Name=\"file1\" Length=\"10\" Offset=\"100\"></Object><Object Name=\"file2\" Length=\"20\" Offset=\"200\"></Object><Object Name=\"file3\" Length=\"30\" Offset=\"300\"></Object></Objects>"

    runBulkGetTest(
        t,
        "start_bulk_verify",
        &stringRequest,
        func(client *Client) ([]models.Objects, error) {
            request, err := client.VerifyBulkJobSpectraS3(models.NewVerifyBulkJobSpectraS3RequestWithPartialObjects("bucketName", objects))
            return request.MasterObjectList.Objects, err
        },
    )
}

func TestBulkVerifyWithObjectNames(t *testing.T) {
    objects := []string {"file1", "file2", "file3"}

    stringRequest := "<Objects><Object Name=\"file1\"></Object><Object Name=\"file2\"></Object><Object Name=\"file3\"></Object></Objects>"

    runBulkGetTest(
        t,
        "start_bulk_verify",
        &stringRequest,
        func(client *Client) ([]models.Objects, error) {
            request, err := client.VerifyBulkJobSpectraS3(models.NewVerifyBulkJobSpectraS3Request("bucketName", objects))
            return request.MasterObjectList.Objects, err
        },
    )
}

type bulkGetTest func(*Client) ([]models.Objects, error)

func runBulkGetTest(t *testing.T, operation string, stringRequest *string, callToTest bulkGetTest) {
    keys := []string { "file2", "file1", "file3" }
    sizes := []int64 { 1202, 256, 2523 }

    stringResponse := "<MasterObjectList><Objects><Object Name='file2' Length='1202'/><Object Name='file1' Length='256'/><Object Name='file3' Length='2523'/></Objects></MasterObjectList>"

    // Create and run the mocked client.
    client := mockedClient(t).
        Expecting(
        HTTP_VERB_PUT,
        "/_rest_/bucket/bucketName",
        &url.Values{"operation": []string{operation}},
        &http.Header{},
        stringRequest,
    ).
        Returning(200, stringResponse, nil)
    response, err := callToTest(client)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
    if len(response) != 1 {
        t.Fatalf("Expected 1 object list but got %d.", len(response))
    }
    if len(response[0].Objects) != len(keys) {
        t.Fatalf("Expected %d objects but got %d.", len(keys), len(response[0].Objects))
    }
    for i, obj := range response[0].Objects {
        ds3Testing.AssertNonNilStringPtr(t, "Name", keys[i], obj.Name)
        ds3Testing.AssertInt64(t, "Length", sizes[i], obj.Length)
    }
}

func TestInitiateMultipart(t *testing.T) {
    stringResponse := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><InitiateMultipartUploadResult xmlns=\"http://s3.amazonaws.com/doc/2006-03-01/\"><Bucket>example-bucket</Bucket><Key>example-object</Key><UploadId>VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA</UploadId></InitiateMultipartUploadResult>"

    // Create and run the mocked client.
    qs := &url.Values{"uploads": []string{""}}
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_POST, "/bucketName/object", qs, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        InitiateMultiPartUpload(models.NewInitiateMultiPartUploadRequest(
            "bucketName",
            "object",
        ))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
    ds3Testing.AssertNonNilStringPtr(t, "Bucket", "example-bucket", response.InitiateMultipartUploadResult.Bucket)
    ds3Testing.AssertNonNilStringPtr(t, "Key", "example-object", response.InitiateMultipartUploadResult.Key)
    ds3Testing.AssertNonNilStringPtr(t, "UploadId", "VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA", response.InitiateMultipartUploadResult.UploadId)
}

func TestPutPart(t *testing.T) {
    content := "this is the part content"
    partNumber := 2
    uploadId := "VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA"
    eTag := "b54357faf0632cce46e942fa68356b38"

    // Create and run the mocked client.
    qs := &url.Values{
        "part_number": []string{strconv.Itoa(partNumber)},
        "upload_id": []string{uploadId},
    }
    responseHeaders := &http.Header{}
    responseHeaders.Add("ETag", eTag)

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_PUT, "/bucketName/object", qs, &http.Header{}, &content).
        Returning(200, "", responseHeaders).
        PutMultiPartUploadPart(models.NewPutMultiPartUploadPartRequest(
            "bucketName",
            "object",
            BuildByteReaderWithSizeDecorator([]byte(content)),
            partNumber,
            uploadId,
        ))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }

    ds3Testing.AssertString(t, "etag header", eTag, response.Headers.Get("etag"))
}

func TestCompleteMultipart(t *testing.T) {
    bucket := "bucketName"
    key := "object"
    location := "http://my-server/bucketName/object"
    etag := "b54357faf0632cce46e942fa68356b38"
    uploadId := "VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA"
    expectedRequest := "<CompleteMultipartUpload><Part><PartNumber>1</PartNumber><ETag>7a112844c1a2327e617f530cb06dccf8</ETag></Part><Part><PartNumber>2</PartNumber><ETag>7162e29f4e40da7f521d0794b57770ba</ETag></Part></CompleteMultipartUpload>"
    expectedResponse := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><CompleteMultipartUploadResult xmlns=\"http://s3.amazonaws.com/doc/2006-03-01/\"><Location>http://my-server/bucketName/object</Location><Bucket>bucketName</Bucket><Key>object</Key><ETag>b54357faf0632cce46e942fa68356b38</ETag></CompleteMultipartUploadResult>"

    // Create and run the mocked client.
    qs := &url.Values{
        "upload_id": []string{uploadId},
    }
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_POST, "/bucketName/object", qs, &http.Header{}, &expectedRequest).
        Returning(200, expectedResponse, &http.Header{"etag": []string{etag}}).
        CompleteMultiPartUpload(models.NewCompleteMultiPartUploadRequest(
            bucket,
            key,
            []models.Part{
                {PartNumber: 1, ETag: "7a112844c1a2327e617f530cb06dccf8"},
                {PartNumber: 2, ETag: "7162e29f4e40da7f521d0794b57770ba"},
            },
            uploadId,
        ))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }

    ds3Testing.AssertNonNilStringPtr(t, "Location", location, response.CompleteMultipartUploadResult.Location)
    ds3Testing.AssertNonNilStringPtr(t, "Bucket", bucket, response.CompleteMultipartUploadResult.Bucket)
    ds3Testing.AssertNonNilStringPtr(t, "Key", key, response.CompleteMultipartUploadResult.Key)
    ds3Testing.AssertNonNilStringPtr(t, "ETag", etag, response.CompleteMultipartUploadResult.ETag)
}

func TestDeleteObjects(t *testing.T) {
    bucket := "bucketName"
    objectNames := []string {"obj1", "obj2", "obj3"}
    expectedRequest := "<Delete><Object><Key>obj1</Key></Object><Object><Key>obj2</Key></Object><Object><Key>obj3</Key></Object></Delete>"
    expectedResponse := "<DeleteResult><Deleted><Key>obj1</Key></Deleted><Deleted><Key>obj2</Key></Deleted><Error><Code>ObjectNotFound</Code><Key>obj3</Key><Message>Object not found</Message></Error></DeleteResult>"

    expectedDeleted := []models.S3ObjectToDelete{{&objectNames[0]}, {&objectNames[1]}}

    code := "ObjectNotFound"
    message := "Object not found"
    expectedErrors := []models.DeleteObjectError {
        {
            Code:    &code,
            Key:     &objectNames[2],
            Message: &message,
        },
    }

    // Create and run the mocked client.
    qs := &url.Values{
        "delete": []string{""},
    }
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_POST, "/bucketName", qs, &http.Header{}, &expectedRequest).
        Returning(200, expectedResponse, &http.Header{}).
        DeleteObjects(models.NewDeleteObjectsRequest(bucket, objectNames))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    if response.DeleteResult.DeletedObjects == nil {
        t.Fatalf("Expected '%s' but got 'nil'", expectedDeleted)
    }

    if !reflect.DeepEqual(response.DeleteResult.DeletedObjects, expectedDeleted) {
        t.Fatalf("Expected '%s' but got '%s'", expectedDeleted, response.DeleteResult.DeletedObjects)
    }

    if response.DeleteResult.Errors == nil {
        t.Fatalf("Expected '%s' but got 'nil'", expectedErrors)
    }

    if !reflect.DeepEqual(response.DeleteResult.Errors, expectedErrors) {
        t.Fatalf("Expected '%s' but got '%s'", expectedErrors, response.DeleteResult.Errors)
    }
}

func TestAllocateJobChunkSpectraS3(t *testing.T) {
    responseString := "<Objects ChunkId=\"203f6886-b058-4f7c-a012-8779176453b1\" ChunkNumber=\"3\" NodeId=\"a02053b9-0147-11e4-8d6a-002590c1177c\">" +
            "<Object Name=\"client00obj000004-8000000\" InCache=\"true\" Length=\"5368709120\" Offset=\"0\"/>" +
            "<Object Name=\"client00obj000004-8000000\" InCache=\"true\" Length=\"2823290880\" Offset=\"5368709120\"/>" +
            "<Object Name=\"client00obj000003-8000000\" InCache=\"true\" Length=\"2823290880\" Offset=\"5368709120\"/>" +
            "<Object Name=\"client00obj000003-8000000\" InCache=\"true\" Length=\"5368709120\" Offset=\"0\"/>" +
            "</Objects>"

    chunkId := "203f6886-b058-4f7c-a012-8779176453b1"
    nodeId := "a02053b9-0147-11e4-8d6a-002590c1177c"

    qp := &url.Values{"operation": []string{"allocate"}}

    request := models.NewAllocateJobChunkSpectraS3Request(chunkId)

    response, err := mockedClient(t).
            Expecting(HTTP_VERB_PUT, "/_rest_/job_chunk/203f6886-b058-4f7c-a012-8779176453b1", qp, &http.Header{}, nil).
            Returning(200, responseString, nil).
            AllocateJobChunkSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    ds3Testing.AssertString(t, "ChunkId", chunkId, response.Objects.ChunkId)
    ds3Testing.AssertInt(t, "ChunkNumber", 3, response.Objects.ChunkNumber)
    ds3Testing.AssertNonNilStringPtr(t, "NodeId", nodeId, response.Objects.NodeId)

    expectedNames := []string{"client00obj000004-8000000", "client00obj000004-8000000", "client00obj000003-8000000", "client00obj000003-8000000"}
    expectedOffsets := []int64{0, 5368709120, 5368709120, 0}
    expectedLengths := []int64{5368709120, 2823290880, 2823290880, 5368709120}

    if len(response.Objects.Objects) != len(expectedNames) {
        t.Fatalf("Expected number of objects '%d' but got '%d'.", len(response.Objects.Objects), len(expectedNames))
    }
    for i, obj := range response.Objects.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "Name", expectedNames[i], obj.Name)
        ds3Testing.AssertNonNilBoolPtr(t, "InCache", true, obj.InCache)
        ds3Testing.AssertInt64(t, "Offset", expectedOffsets[i], obj.Offset)
        ds3Testing.AssertInt64(t, "Length", expectedLengths[i], obj.Length)
    }
}

const (
    test_master_object_list_id  = "1a85e743-ec8f-4789-afec-97e587a26936"
    test_master_object_list_xml = "<MasterObjectList BucketName=\"bucket8192000000\" JobId=\"1a85e743-ec8f-4789-afec-97e587a26936\" Priority=\"NORMAL\" RequestType=\"GET\" StartDate=\"2014-07-01T20:12:52.000Z\">" +
            "  <Nodes>" +
            "    <Node EndPoint=\"10.1.18.12\" HttpPort=\"80\" HttpsPort=\"443\" Id=\"a02053b9-0147-11e4-8d6a-002590c1177c\"/>" +
            "    <Node EndPoint=\"10.1.18.13\" HttpsPort=\"443\" Id=\"95e97010-8e70-4733-926c-aeeb21796848\"/>" +
            "  </Nodes>" +
            "  <Objects ChunkId=\"f58370c2-2538-4e78-a9f8-e4d2676bdf44\" ChunkNumber=\"0\" NodeId=\"a02053b9-0147-11e4-8d6a-002590c1177c\">" +
            "    <Object Name=\"client00obj000004-8000000\" InCache=\"true\" Length=\"5368709120\" Offset=\"0\"/>" +
            "    <Object Name=\"client00obj000004-8000000\" InCache=\"true\" Length=\"2823290880\" Offset=\"5368709120\"/>" +
            "  </Objects>" +
            "  <Objects ChunkId=\"4137d768-25bb-4942-9d36-b92dfbe75e01\" ChunkNumber=\"1\" NodeId=\"95e97010-8e70-4733-926c-aeeb21796848\">" +
            "    <Object Name=\"client00obj000008-8000000\" InCache=\"true\" Length=\"2823290880\" Offset=\"5368709120\"/>" +
            "    <Object Name=\"client00obj000008-8000000\" InCache=\"true\" Length=\"5368709120\" Offset=\"0\"/>" +
            "  </Objects>" +
            "</MasterObjectList>"
)

func verifyMasterObjectList(t *testing.T, masterObjectList models.MasterObjectList) {
    bucketName := "bucket8192000000"
    startDate := "2014-07-01T20:12:52.000Z"
    ds3Testing.AssertNonNilStringPtr(t, "BucketName", bucketName, masterObjectList.BucketName)
    ds3Testing.AssertString(t, "JobId", test_master_object_list_id, masterObjectList.JobId)
    if masterObjectList.Priority != models.PRIORITY_NORMAL {
        t.Fatalf("Expected priority '%d' but got '%d'.", models.PRIORITY_NORMAL, masterObjectList.Priority)
    }
    if masterObjectList.RequestType != models.JOB_REQUEST_TYPE_GET {
        t.Fatalf("Expected request type '%d' but got '%d'.", models.JOB_REQUEST_TYPE_GET, masterObjectList.RequestType)
    }
    ds3Testing.AssertString(t, "StartDate", startDate, masterObjectList.StartDate)

    // verify nodes
    if len(masterObjectList.Nodes) != 2 {
        t.Fatalf("Expected there to be 2 nodes but got '%d'.", len(masterObjectList.Nodes))
    }

    node1 := masterObjectList.Nodes[0]
    ds3Testing.AssertNonNilStringPtr(t, "EndPoint", "10.1.18.12", node1.EndPoint)
    ds3Testing.AssertNonNilIntPtr(t, "HttpPort", 80, node1.HttpPort)
    ds3Testing.AssertNonNilIntPtr(t, "HttpsPort", 443, node1.HttpsPort)
    ds3Testing.AssertString(t, "Id", "a02053b9-0147-11e4-8d6a-002590c1177c", node1.Id)

    node2 := masterObjectList.Nodes[1]
    ds3Testing.AssertNonNilStringPtr(t, "EndPoint", "10.1.18.13", node2.EndPoint)
    ds3Testing.AssertIntPtrIsNil(t, "HttpPort", node2.HttpPort)
    ds3Testing.AssertNonNilIntPtr(t, "HttpsPort", 443, node2.HttpsPort)
    ds3Testing.AssertString(t, "Id", "95e97010-8e70-4733-926c-aeeb21796848", node2.Id)

    // verify objects
    if len(masterObjectList.Objects) != 2 {
        t.Fatalf("Expected 2 object but got '%d'.", len(masterObjectList.Objects))
    }

    obj1 := masterObjectList.Objects[0]
    ds3Testing.AssertString(t, "ChunkId", "f58370c2-2538-4e78-a9f8-e4d2676bdf44", obj1.ChunkId)
    ds3Testing.AssertInt(t, "ChunkNumber", 0, obj1.ChunkNumber)
    ds3Testing.AssertNonNilStringPtr(t, "NodeId", "a02053b9-0147-11e4-8d6a-002590c1177c", obj1.NodeId)

    if len(obj1.Objects) != 2 {
        t.Fatalf("Expected 2 bulk objects but got '%d'.", len(obj1.Objects))
    }

    obj1BulkObj1 := obj1.Objects[0]
    ds3Testing.AssertNonNilStringPtr(t, "Name", "client00obj000004-8000000", obj1BulkObj1.Name)
    ds3Testing.AssertNonNilBoolPtr(t, "InCache", true, obj1BulkObj1.InCache)
    ds3Testing.AssertInt64(t, "Length", 5368709120, obj1BulkObj1.Length)
    ds3Testing.AssertInt64(t, "Offset", 0, obj1BulkObj1.Offset)

    obj1BulkObj2 := obj1.Objects[1]
    ds3Testing.AssertNonNilStringPtr(t, "Name", "client00obj000004-8000000", obj1BulkObj2.Name)
    ds3Testing.AssertNonNilBoolPtr(t, "InCache", true, obj1BulkObj2.InCache)
    ds3Testing.AssertInt64(t, "Length", 2823290880, obj1BulkObj2.Length)
    ds3Testing.AssertInt64(t, "Offset", 5368709120, obj1BulkObj2.Offset)

    obj2 := masterObjectList.Objects[1]
    ds3Testing.AssertString(t, "ChunkId", "4137d768-25bb-4942-9d36-b92dfbe75e01", obj2.ChunkId)
    ds3Testing.AssertInt(t, "ChunkNumber", 1, obj2.ChunkNumber)
    ds3Testing.AssertNonNilStringPtr(t, "NodeId", "95e97010-8e70-4733-926c-aeeb21796848", obj2.NodeId)
    if len(obj2.Objects) != 2 {
        t.Fatalf("Expected 2 bulk objects but got '%d'.", len(obj2.Objects))
    }

    obj2BulkObj1 := obj2.Objects[0]
    ds3Testing.AssertNonNilStringPtr(t, "Name", "client00obj000008-8000000", obj2BulkObj1.Name)
    ds3Testing.AssertNonNilBoolPtr(t, "InCache", true, obj2BulkObj1.InCache)
    ds3Testing.AssertInt64(t, "Length", 2823290880, obj2BulkObj1.Length)
    ds3Testing.AssertInt64(t, "Offset", 5368709120, obj2BulkObj1.Offset)

    obj2BulkObj2 := obj2.Objects[1]
    ds3Testing.AssertNonNilStringPtr(t, "Name", "client00obj000008-8000000", obj2BulkObj2.Name)
    ds3Testing.AssertNonNilBoolPtr(t, "InCache", true, obj2BulkObj2.InCache)
    ds3Testing.AssertInt64(t, "Length", 5368709120, obj2BulkObj2.Length)
    ds3Testing.AssertInt64(t, "Offset", 0, obj2BulkObj2.Offset)
}

func TestGetJobChunksReadyForClientProcessingSpectraS3(t *testing.T) {
    qp := &url.Values{"job": []string{test_master_object_list_id}}

    request := models.NewGetJobChunksReadyForClientProcessingSpectraS3Request(test_master_object_list_id)
    response, err := mockedClient(t).
            Expecting(HTTP_VERB_GET, "/_rest_/job_chunk", qp, &http.Header{}, nil).
            Returning(200, test_master_object_list_xml, nil).
            GetJobChunksReadyForClientProcessingSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    verifyMasterObjectList(t, response.MasterObjectList)
}

func TestGetJobSpectraS3(t *testing.T) {
    request := models.NewGetJobSpectraS3Request(test_master_object_list_id)
    response, err := mockedClient(t).
            Expecting(HTTP_VERB_GET, "/_rest_/job/1a85e743-ec8f-4789-afec-97e587a26936", &url.Values{}, &http.Header{}, nil).
            Returning(200, test_master_object_list_xml, nil).
            GetJobSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    verifyMasterObjectList(t, response.MasterObjectList)
}

func TestModifyJobSpectraS3(t *testing.T) {
    request := models.NewModifyJobSpectraS3Request(test_master_object_list_id)
    response, err := mockedClient(t).
            Expecting(HTTP_VERB_PUT, "/_rest_/job/1a85e743-ec8f-4789-afec-97e587a26936", &url.Values{}, &http.Header{}, nil).
            Returning(200, test_master_object_list_xml, nil).
            ModifyJobSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    verifyMasterObjectList(t, response.MasterObjectList)
}

func TestGetJobsSpectraS3(t *testing.T) {
    responseString := "<Jobs>" +
            "  <Job BucketName=\"bucket_1\" CachedSizeInBytes=\"69880\" ChunkClientProcessingOrderGuarantee=\"IN_ORDER\" CompletedSizeInBytes=\"0\" JobId=\"0807ff11-a9f6-4d55-bb92-b452c1bb00c7\" OriginalSizeInBytes=\"69880\" Priority=\"NORMAL\" RequestType=\"PUT\" StartDate=\"2014-09-04T17:23:45.000Z\" UserId=\"a7d3eff9-e6d2-4e37-8a0b-84e76211a18a\" UserName=\"spectra\">" +
            "    <Nodes>" +
            "      <Node EndPoint=\"10.10.10.10\" HttpPort=\"80\" HttpsPort=\"443\" Id=\"edb8cc38-32f2-11e4-bce1-080027ecf0d4\"/>" +
            "    </Nodes>" +
            "  </Job>" +
            "  <Job BucketName=\"bucket_2\" CachedSizeInBytes=\"0\" ChunkClientProcessingOrderGuarantee=\"IN_ORDER\" CompletedSizeInBytes=\"0\" JobId=\"c18554ba-e3a8-4905-91fd-3e6eec71bf45\" OriginalSizeInBytes=\"69880\" Priority=\"HIGH\" RequestType=\"GET\" StartDate=\"2014-09-04T17:24:04.000Z\" UserId=\"a7d3eff9-e6d2-4e37-8a0b-84e76211a18a\" UserName=\"spectra\">" +
            "    <Nodes>" +
            "      <Node EndPoint=\"10.10.10.10\" HttpPort=\"80\" HttpsPort=\"443\" Id=\"edb8cc38-32f2-11e4-bce1-080027ecf0d4\"/>" +
            "    </Nodes>" +
            "  </Job>" +
            "</Jobs>"

    response, err := mockedClient(t).
            Expecting(HTTP_VERB_GET, "/_rest_/job", &url.Values{}, &http.Header{}, nil).
            Returning(200, responseString, nil).
            GetJobsSpectraS3(models.NewGetJobsSpectraS3Request())

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    if len(response.JobList.Jobs) != 2 {
        t.Fatalf("Expected 2 jobs but got '%d'.", len(response.JobList.Jobs))
    }

    job1 := response.JobList.Jobs[0]
    ds3Testing.AssertNonNilStringPtr(t, "BucketName", "bucket_1", job1.BucketName)
    ds3Testing.AssertInt64(t, "CachedSizeInBytes", 69880, job1.CachedSizeInBytes)
    if job1.ChunkClientProcessingOrderGuarantee != models.JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER {
        t.Fatalf("Expected chunk client processing order guarantee '%d' but got '%d'.",
            models.JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER,
            job1.ChunkClientProcessingOrderGuarantee,
        )
    }
    ds3Testing.AssertInt64(t, "CompletedSizeInBytes", 0, job1.CompletedSizeInBytes)
    ds3Testing.AssertString(t, "JobId", "0807ff11-a9f6-4d55-bb92-b452c1bb00c7", job1.JobId)
    ds3Testing.AssertInt64(t, "OriginalSizeInBytes", 69880, job1.OriginalSizeInBytes)
    if job1.Priority != models.PRIORITY_NORMAL {
        t.Fatalf("Expected priority '%d' but got '%d'.", models.PRIORITY_NORMAL, job1.Priority)
    }
    if job1.RequestType != models.JOB_REQUEST_TYPE_PUT {
        t.Fatalf("Expected request type '%d' but got '%d'.", models.JOB_REQUEST_TYPE_PUT, job1.RequestType)
    }
    ds3Testing.AssertString(t, "StartDate", "2014-09-04T17:23:45.000Z", job1.StartDate)
    ds3Testing.AssertString(t, "UserId", "a7d3eff9-e6d2-4e37-8a0b-84e76211a18a", job1.UserId)
    ds3Testing.AssertNonNilStringPtr(t, "UserName", "spectra", job1.UserName)
    if len(job1.Nodes) != 1 {
        t.Fatalf("Expected '1' nodes but got '%d'.", len(job1.Nodes))
    }

    job1node := job1.Nodes[0]
    ds3Testing.AssertNonNilStringPtr(t, "EndPoint", "10.10.10.10", job1node.EndPoint)
    ds3Testing.AssertNonNilIntPtr(t, "HttpPort", 80, job1node.HttpPort)
    ds3Testing.AssertNonNilIntPtr(t, "HttpsPort", 443, job1node.HttpsPort)
    ds3Testing.AssertString(t, "Id", "edb8cc38-32f2-11e4-bce1-080027ecf0d4", job1node.Id)

    job2 := response.JobList.Jobs[1]
    ds3Testing.AssertNonNilStringPtr(t, "BucketName", "bucket_2", job2.BucketName)
    ds3Testing.AssertInt64(t, "CachedSizeInBytes", 0, job2.CachedSizeInBytes)
    if job2.ChunkClientProcessingOrderGuarantee != models.JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER {
        t.Fatalf("Expected chunk client processing order guarantee '%d' but got '%d'.",
            models.JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER,
            job2.ChunkClientProcessingOrderGuarantee,
        )
    }
    ds3Testing.AssertInt64(t, "CompletedSizeInBytes", 0, job2.CompletedSizeInBytes)
    ds3Testing.AssertString(t, "JobId", "c18554ba-e3a8-4905-91fd-3e6eec71bf45", job2.JobId)
    ds3Testing.AssertInt64(t, "OriginalSizeInBytes", 69880, job2.OriginalSizeInBytes)
    if job2.Priority != models.PRIORITY_HIGH {
        t.Fatalf("Expected priority '%d' but got '%d'.", models.PRIORITY_HIGH, job2.Priority)
    }
    if job2.RequestType != models.JOB_REQUEST_TYPE_GET {
        t.Fatalf("Expected request type '%d' but got '%d'.", models.JOB_REQUEST_TYPE_GET, job2.RequestType)
    }
    ds3Testing.AssertString(t, "StartDate", "2014-09-04T17:24:04.000Z", job2.StartDate)
    ds3Testing.AssertString(t, "UserId", "a7d3eff9-e6d2-4e37-8a0b-84e76211a18a", job2.UserId)
    ds3Testing.AssertNonNilStringPtr(t, "UserName", "spectra", job2.UserName)
    if len(job2.Nodes) != 1 {
        t.Fatalf("Expected '1' nodes but got '%d'.", len(job2.Nodes))
    }

    job2node := job2.Nodes[0]
    ds3Testing.AssertNonNilStringPtr(t, "EndPoint", "10.10.10.10", job2node.EndPoint)
    ds3Testing.AssertNonNilIntPtr(t, "HttpPort", 80, job2node.HttpPort)
    ds3Testing.AssertNonNilIntPtr(t, "HttpsPort", 443, job2node.HttpsPort)
    ds3Testing.AssertString(t, "Id", "edb8cc38-32f2-11e4-bce1-080027ecf0d4", job2node.Id)
}

func TestCancelJobSpectraS3(t *testing.T) {
    jobId := "1a85e743-ec8f-4789-afec-97e587a26936"
    qp := &url.Values{"force": []string{""}}
    request := models.NewCancelJobSpectraS3Request(jobId)
    response, err := mockedClient(t).
            Expecting(HTTP_VERB_DELETE, "/_rest_/job/1a85e743-ec8f-4789-afec-97e587a26936", qp, &http.Header{}, nil).
            Returning(204, "", nil).
            CancelJobSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestDeleteTapeDriveSpectraS3(t *testing.T) {
    jobId := "1a85e743-ec8f-4789-afec-97e587a26936"

    request := models.NewDeleteTapeDriveSpectraS3Request(jobId)
    response, err := mockedClient(t).
            Expecting(HTTP_VERB_DELETE, "/_rest_/tape_drive/1a85e743-ec8f-4789-afec-97e587a26936", &url.Values{}, &http.Header{}, nil).
            Returning(204, "", nil).
            DeleteTapeDriveSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestDeleteTapePartitionSpectraS3(t *testing.T) {
    id := "1a85e743-ec8f-4789-afec-97e587a26936"

    request := models.NewDeleteTapePartitionSpectraS3Request(id)
    response, err := mockedClient(t).
            Expecting(HTTP_VERB_DELETE, "/_rest_/tape_partition/1a85e743-ec8f-4789-afec-97e587a26936", &url.Values{}, &http.Header{}, nil).
            Returning(204, "", nil).
            DeleteTapePartitionSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestVerifySystemHealthSpectraS3(t *testing.T) {
    responsePayload := "<Data><MsRequiredToVerifyDataPlannerHealth>10</MsRequiredToVerifyDataPlannerHealth></Data>"
    response, err := mockedClient(t).
            Expecting(HTTP_VERB_GET, "/_rest_/system_health", &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, nil).
            VerifySystemHealthSpectraS3(models.NewVerifySystemHealthSpectraS3Request())

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    ds3Testing.AssertInt64(t, "MsRequiredToVerifyDataPlannerHealth", 10, response.HealthVerificationResult.MsRequiredToVerifyDataPlannerHealth)
}

func TestGetSystemInformationSpectraS3(t *testing.T) {
    responsePayload := "<Data>" +
            "<ApiVersion>518B3F2A95B71AC7325EFB12B2937376.15F3CC0489CBCD4648ECFF0FBF371B8A</ApiVersion>" +
            "<BuildInformation><Branch/><Revision/><Version/></BuildInformation>" +
            "<SerialNumber>UNKNOWN</SerialNumber>" +
            "</Data>"
    response, err := mockedClient(t).
            Expecting(HTTP_VERB_GET, "/_rest_/system_information", &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, nil).
            GetSystemInformationSpectraS3(models.NewGetSystemInformationSpectraS3Request())

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    ds3Testing.AssertNonNilStringPtr(t, "ApiVersion", "518B3F2A95B71AC7325EFB12B2937376.15F3CC0489CBCD4648ECFF0FBF371B8A", response.SystemInformation.ApiVersion)
    ds3Testing.AssertNonNilStringPtr(t, "SerialNumber", "UNKNOWN", response.SystemInformation.SerialNumber)
}

func TestGetTapeLibrariesSpectraS3(t *testing.T) {
    responsePayload := "<Data>" +
            "<TapeLibrary><Id>f4dae25d-e52a-4430-82bd-525e4f15493c</Id><ManagementUrl>a</ManagementUrl><Name>test library</Name><SerialNumber>test library</SerialNumber></TapeLibrary>" +
            "<TapeLibrary><Id>82bdab72-d79a-4b43-95d7-f2c16cd9aa45</Id><ManagementUrl>a</ManagementUrl><Name>test library 2</Name><SerialNumber>test library 2</SerialNumber></TapeLibrary>" +
            "</Data>"

    responseHeaders := &http.Header{}
    responseHeaders.Add("Page-Truncated", "2")
    responseHeaders.Add("Total-Result-Count", "3")

    response, err := mockedClient(t).
            Expecting(HTTP_VERB_GET, "/_rest_/tape_library", &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, responseHeaders).
            GetTapeLibrariesSpectraS3(models.NewGetTapeLibrariesSpectraS3Request())

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    if len(response.TapeLibraryList.TapeLibraries) != 2 {
        t.Fatalf("Expected '2' tape libraries but got '%d'.", len(response.TapeLibraryList.TapeLibraries))
    }

    lib1 := response.TapeLibraryList.TapeLibraries[0]
    ds3Testing.AssertString(t, "Id", "f4dae25d-e52a-4430-82bd-525e4f15493c", lib1.Id)
    ds3Testing.AssertNonNilStringPtr(t, "ManagementUrl", "a", lib1.ManagementUrl)
    ds3Testing.AssertNonNilStringPtr(t, "Name", "test library", lib1.Name)
    ds3Testing.AssertNonNilStringPtr(t, "SerialNumber", "test library", lib1.SerialNumber)

    lib2 := response.TapeLibraryList.TapeLibraries[1]
    ds3Testing.AssertString(t, "Id", "82bdab72-d79a-4b43-95d7-f2c16cd9aa45", lib2.Id)
    ds3Testing.AssertNonNilStringPtr(t, "ManagementUrl", "a", lib2.ManagementUrl)
    ds3Testing.AssertNonNilStringPtr(t, "Name", "test library 2", lib2.Name)
    ds3Testing.AssertNonNilStringPtr(t, "SerialNumber", "test library 2", lib2.SerialNumber)

    ds3Testing.AssertString(t, "Page-Truncated header", "2", response.Headers.Get("Page-Truncated"))
    ds3Testing.AssertString(t, "Total-Result-Count header", "3", response.Headers.Get("Total-Result-Count"))
}

func TestGetTapeLibrarySpectraS3(t *testing.T) {
    responsePayload := "<Data><Id>e23030e5-9b8d-4594-bdd1-15d3c45abb9f</Id><ManagementUrl>a</ManagementUrl><Name>125ca16e-60e3-43b2-a26f-0bc81843745f</Name><SerialNumber>test library</SerialNumber></Data>"

    id := "e23030e5-9b8d-4594-bdd1-15d3c45abb9f"

    response, err := mockedClient(t).
            Expecting(HTTP_VERB_GET, "/_rest_/tape_library/" + id, &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, nil).
            GetTapeLibrarySpectraS3(models.NewGetTapeLibrarySpectraS3Request(id))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    library := response.TapeLibrary
    ds3Testing.AssertString(t, "Id", id, library.Id)
    ds3Testing.AssertNonNilStringPtr(t, "ManagementUrl", "a", library.ManagementUrl)
    ds3Testing.AssertNonNilStringPtr(t, "Name", "125ca16e-60e3-43b2-a26f-0bc81843745f", library.Name)
    ds3Testing.AssertNonNilStringPtr(t, "SerialNumber", "test library", library.SerialNumber)
}

func TestGetTapeDriveSpectraS3(t *testing.T) {
    responsePayload := "<Data><ErrorMessage/><ForceTapeRemoval>false</ForceTapeRemoval><Id>ff5df6c8-7e24-4e4f-815d-a8a1a4cddc98</Id><PartitionId>ca69b187-47cf-425e-b92f-c09bacc7d3b3</PartitionId><SerialNumber>test tape drive</SerialNumber><State>NORMAL</State><TapeId>0ea07c32-8ff6-443f-b7c8-420667b0df84</TapeId><Type>UNKNOWN</Type></Data>"

    id := "ff5df6c8-7e24-4e4f-815d-a8a1a4cddc98"

    response, err := mockedClient(t).
            Expecting(HTTP_VERB_GET, "/_rest_/tape_drive/" + id, &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, nil).
            GetTapeDriveSpectraS3(models.NewGetTapeDriveSpectraS3Request(id))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    drive := response.TapeDrive
    ds3Testing.AssertStringPtrIsNil(t, "ErrorMessage", drive.ErrorMessage)
    ds3Testing.AssertBool(t, "ForceTapeRemoval", false, drive.ForceTapeRemoval)
    ds3Testing.AssertString(t, "Id", id, drive.Id)
    ds3Testing.AssertString(t, "PartitionId", "ca69b187-47cf-425e-b92f-c09bacc7d3b3", drive.PartitionId)
    ds3Testing.AssertNonNilStringPtr(t, "Serial Number", "test tape drive", drive.SerialNumber)
    if drive.State != models.TAPE_DRIVE_STATE_NORMAL {
        t.Fatalf("Expected drive stat '%d' but was '%d'.", models.TAPE_DRIVE_STATE_NORMAL, drive.State)
    }
    ds3Testing.AssertNonNilStringPtr(t, "TapeId", "0ea07c32-8ff6-443f-b7c8-420667b0df84", drive.TapeId)
    if drive.Type != models.TAPE_DRIVE_TYPE_UNKNOWN {
        t.Fatalf("Expected tape drive type '%d' but was '%d'.", models.TAPE_DRIVE_TYPE_UNKNOWN, drive.Type)
    }
}

func TestGetTapesSpectraS3(t *testing.T) {
    responsePayload := "<Data><Tape>" +
            "<AssignedToStorageDomain>false</AssignedToStorageDomain>" +
            "<AvailableRawCapacity>2408082046976</AvailableRawCapacity>" +
            "<BarCode>101000L6</BarCode>" +
            "<BucketId/><DescriptionForIdentification/><EjectDate/><EjectLabel/><EjectLocation/><EjectPending/>" +
            "<FullOfData>false</FullOfData>" +
            "<Id>c7c431df-f95d-4533-b350-ffd7a8a5caac</Id>" +
            "<LastAccessed>2015-09-04T06:53:08.236</LastAccessed>" +
            "<LastCheckpoint>eb77ea67-3c83-47ec-8714-cd46a97dc392:2</LastCheckpoint>" +
            "<LastModified>2015-08-21T16:14:30.714</LastModified>" +
            "<LastVerified/>" +
            "<PartitionId>4f8a5cbb-9837-41d9-afd1-cebed41f18f7</PartitionId>" +
            "<PreviousState/>" +
            "<SerialNumber>HP-W130501213</SerialNumber>" +
            "<State>NORMAL</State>" +
            "<TotalRawCapacity>2408088338432</TotalRawCapacity>" +
            "<Type>LTO6</Type>" +
            "<WriteProtected>false</WriteProtected>" +
            "</Tape></Data>"

    responseHeaders := &http.Header{}
    responseHeaders.Add("Page-Truncated", "2")
    responseHeaders.Add("Total-Result-Count", "3")

    response, err := mockedClient(t).
            Expecting(HTTP_VERB_GET, "/_rest_/tape", &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, responseHeaders).
            GetTapesSpectraS3(models.NewGetTapesSpectraS3Request())

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    if len(response.TapeList.Tapes) != 1 {
        t.Fatalf("Expected '1' tapes but got '%d'.", len(response.TapeList.Tapes))
    }

    tape := response.TapeList.Tapes[0]
    ds3Testing.AssertBool(t, "AssignedToStorageDomain", false, tape.AssignedToStorageDomain)
    ds3Testing.AssertNonNilInt64Ptr(t, "AvailableRawCapacity", 2408082046976, tape.AvailableRawCapacity)
    ds3Testing.AssertNonNilStringPtr(t, "BarCode", "101000L6", tape.BarCode)
    ds3Testing.AssertStringPtrIsNil(t, "BucketId", tape.BucketId)
    ds3Testing.AssertStringPtrIsNil(t, "DescriptionForIdentification", tape.DescriptionForIdentification)
    ds3Testing.AssertStringPtrIsNil(t, "EjectDate", tape.EjectDate)
    ds3Testing.AssertStringPtrIsNil(t, "EjectLabel", tape.EjectLabel)
    ds3Testing.AssertStringPtrIsNil(t, "EjectLocation", tape.EjectLocation)
    ds3Testing.AssertStringPtrIsNil(t, "EjectPending", tape.EjectPending)
    ds3Testing.AssertBool(t, "FullOfData", false, tape.FullOfData)
    ds3Testing.AssertString(t, "Id", "c7c431df-f95d-4533-b350-ffd7a8a5caac", tape.Id)
    ds3Testing.AssertNonNilStringPtr(t, "LastAccessed", "2015-09-04T06:53:08.236", tape.LastAccessed)
    ds3Testing.AssertNonNilStringPtr(t, "LastCheckpoint", "eb77ea67-3c83-47ec-8714-cd46a97dc392:2", tape.LastCheckpoint)
    ds3Testing.AssertNonNilStringPtr(t, "LastModified", "2015-08-21T16:14:30.714", tape.LastModified)
    ds3Testing.AssertStringPtrIsNil(t, "LastVerified", tape.LastVerified)
    ds3Testing.AssertNonNilStringPtr(t, "PartitionId", "4f8a5cbb-9837-41d9-afd1-cebed41f18f7", tape.PartitionId)
    if tape.PreviousState != nil {
        t.Fatalf("Expected previous state '%d' but was '%d'.", "nil", *tape.PreviousState)
    }
    ds3Testing.AssertNonNilStringPtr(t, "SerialNumber", "HP-W130501213", tape.SerialNumber)
    ds3Testing.AssertString(t, "State", models.TAPE_STATE_NORMAL.String(), tape.State.String())
    ds3Testing.AssertNonNilInt64Ptr(t, "TotalRawCapacity", 2408088338432, tape.TotalRawCapacity)
    ds3Testing.AssertString(t, "Type", "LTO6", tape.Type)
    ds3Testing.AssertBool(t, "WriteProtected", false, tape.WriteProtected)

    ds3Testing.AssertString(t, "Page-Truncated header", "2", response.Headers.Get("Page-Truncated"))
    ds3Testing.AssertString(t, "Total-Result-Count header", "3", response.Headers.Get("Total-Result-Count"))
}

func TestDeletePermanentlyLostTapeSpectraS3(t *testing.T) {
    id := "1a85e743-ec8f-4789-afec-97e587a26936"

    request := models.NewDeletePermanentlyLostTapeSpectraS3Request(id)
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_DELETE, "/_rest_/tape/" + id, &url.Values{}, &http.Header{}, nil).
        Returning(204, "", nil).
        DeletePermanentlyLostTapeSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestGetTapeSpectraS3(t *testing.T) {
    id := "c7c431df-f95d-4533-b350-ffd7a8a5caac"
    responsePayload := "<Data><AssignedToStorageDomain>false</AssignedToStorageDomain><AvailableRawCapacity>2408082046976</AvailableRawCapacity><BarCode>101000L6</BarCode><BucketId/><DescriptionForIdentification/><EjectDate/><EjectLabel/><EjectLocation/><EjectPending/><FullOfData>false</FullOfData><Id>c7c431df-f95d-4533-b350-ffd7a8a5caac</Id><LastAccessed>2015-09-04T06:53:08.236</LastAccessed><LastCheckpoint>eb77ea67-3c83-47ec-8714-cd46a97dc392:2</LastCheckpoint><LastModified>2015-08-21T16:14:30.714</LastModified><LastVerified/><PartitionId>4f8a5cbb-9837-41d9-afd1-cebed41f18f7</PartitionId><PreviousState/><SerialNumber>HP-W130501213</SerialNumber><State>NORMAL</State><TotalRawCapacity>2408088338432</TotalRawCapacity><Type>LTO6</Type><WriteProtected>false</WriteProtected></Data>"

    request := models.NewGetTapeSpectraS3Request(id)
    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/_rest_/tape/" + id, &url.Values{}, &http.Header{}, nil).
        Returning(200, responsePayload, nil).
        GetTapeSpectraS3(request)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    tape := response.Tape
    ds3Testing.AssertBool(t, "AssignedToStorageDomain", false, tape.AssignedToStorageDomain)
    ds3Testing.AssertNonNilInt64Ptr(t, "AvailableRawCapacity", 2408082046976, tape.AvailableRawCapacity)
    ds3Testing.AssertNonNilStringPtr(t, "BarCode", "101000L6", tape.BarCode)
    ds3Testing.AssertStringPtrIsNil(t, "BucketId", tape.BucketId)
    ds3Testing.AssertStringPtrIsNil(t, "DescriptionForIdentification", tape.DescriptionForIdentification)
    ds3Testing.AssertStringPtrIsNil(t, "EjectDate", tape.EjectDate)
    ds3Testing.AssertStringPtrIsNil(t, "EjectLabel", tape.EjectLabel)
    ds3Testing.AssertStringPtrIsNil(t, "EjectLocation", tape.EjectLocation)
    ds3Testing.AssertStringPtrIsNil(t, "EjectPending", tape.EjectPending)
    ds3Testing.AssertBool(t, "FullOfData", false, tape.FullOfData)
    ds3Testing.AssertString(t, "Id", "c7c431df-f95d-4533-b350-ffd7a8a5caac", tape.Id)
    ds3Testing.AssertNonNilStringPtr(t, "LastAccessed", "2015-09-04T06:53:08.236", tape.LastAccessed)
    ds3Testing.AssertNonNilStringPtr(t, "LastCheckpoint", "eb77ea67-3c83-47ec-8714-cd46a97dc392:2", tape.LastCheckpoint)
    ds3Testing.AssertNonNilStringPtr(t, "LastModified", "2015-08-21T16:14:30.714", tape.LastModified)
    ds3Testing.AssertStringPtrIsNil(t, "LastVerified", tape.LastVerified)
    ds3Testing.AssertNonNilStringPtr(t, "PartitionId", "4f8a5cbb-9837-41d9-afd1-cebed41f18f7", tape.PartitionId)
    if tape.PreviousState != nil {
        t.Fatalf("Expected previous state '%d' but was '%d'.", "nil", *tape.PreviousState)
    }
    ds3Testing.AssertNonNilStringPtr(t, "SerialNumber", "HP-W130501213", tape.SerialNumber)
    ds3Testing.AssertString(t, "State", models.TAPE_STATE_NORMAL.String(), tape.State.String())
    ds3Testing.AssertNonNilInt64Ptr(t, "TotalRawCapacity", 2408088338432, tape.TotalRawCapacity)
    ds3Testing.AssertString(t, "Type", "LTO6", tape.Type)
    ds3Testing.AssertBool(t, "WriteProtected", false, tape.WriteProtected)
}

func TestClearSuspectBlobAzureTargetsSpectraS3(t *testing.T) {
    expectedRequest := "<Ids><Id>id1</Id><Id>id2</Id><Id>id3</Id></Ids>"

    // Create and run the mocked client.
    ids := []string {"id1", "id2", "id3"}

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_DELETE, "/_rest_/suspect_blob_azure_target", &url.Values{}, &http.Header{}, &expectedRequest).
        Returning(204, "", nil).
        ClearSuspectBlobAzureTargetsSpectraS3(models.NewClearSuspectBlobAzureTargetsSpectraS3Request(ids))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
}

func TestClearSuspectBlobPoolsSpectraS3(t *testing.T) {
    expectedRequest := "<Ids><Id>id1</Id><Id>id2</Id><Id>id3</Id></Ids>"

    // Create and run the mocked client.
    ids := []string {"id1", "id2", "id3"}

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_DELETE, "/_rest_/suspect_blob_pool", &url.Values{}, &http.Header{}, &expectedRequest).
        Returning(204, "", nil).
        ClearSuspectBlobPoolsSpectraS3(models.NewClearSuspectBlobPoolsSpectraS3Request(ids))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
}

func TestClearSuspectBlobS3TargetsSpectraS3(t *testing.T) {
    expectedRequest := "<Ids><Id>id1</Id><Id>id2</Id><Id>id3</Id></Ids>"

    // Create and run the mocked client.
    ids := []string {"id1", "id2", "id3"}

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_DELETE, "/_rest_/suspect_blob_s3_target", &url.Values{}, &http.Header{}, &expectedRequest).
        Returning(204, "", nil).
        ClearSuspectBlobS3TargetsSpectraS3(models.NewClearSuspectBlobS3TargetsSpectraS3Request(ids))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
}

func TestClearSuspectBlobTapesSpectraS3(t *testing.T) {
    expectedRequest := "<Ids><Id>id1</Id><Id>id2</Id><Id>id3</Id></Ids>"

    // Create and run the mocked client.
    ids := []string {"id1", "id2", "id3"}

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_DELETE, "/_rest_/suspect_blob_tape", &url.Values{}, &http.Header{}, &expectedRequest).
        Returning(204, "", nil).
        ClearSuspectBlobTapesSpectraS3(models.NewClearSuspectBlobTapesSpectraS3Request(ids))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
}

func TestEjectStorageDomainBlobsSpectraS3(t *testing.T) {
    expectedRequest := "<Objects><Object Name=\"obj1\"></Object><Object Name=\"obj2\"></Object><Object Name=\"obj3\"></Object></Objects>"

    // Create and run the mocked client.
    bucketId := "BucketId"
    storageDomainId := "StorageDomainId"
    objectNames := []string {"obj1", "obj2", "obj3"}

    qp := &url.Values{
        "operation": []string{"eject"},
        "blobs": []string{""},
        "bucket_id": []string{bucketId},
        "storage_domain_id": []string{storageDomainId},
    }

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_PUT, "/_rest_/tape", qp, &http.Header{}, &expectedRequest).
        Returning(204, "", nil).
        EjectStorageDomainBlobsSpectraS3(models.NewEjectStorageDomainBlobsSpectraS3Request(bucketId, objectNames, storageDomainId))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
}

func TestGetBlobsOnAzureTargetSpectraS3(t *testing.T) {
    target := "AzureTarget"
    runGetBlobsTest(
        t,
        "/_rest_/azure_target/" + target,
        func(client *Client) (models.BulkObjectList, error) {
            request, err := client.GetBlobsOnAzureTargetSpectraS3(models.NewGetBlobsOnAzureTargetSpectraS3Request(target))
            return request.BulkObjectList, err
        },
    )
}

func TestGetBlobsOnDs3TargetSpectraS3(t *testing.T) {
    target := "Ds3Target"
    runGetBlobsTest(
        t,
        "/_rest_/ds3_target/" + target,
        func(client *Client) (models.BulkObjectList, error) {
            request, err := client.GetBlobsOnDs3TargetSpectraS3(models.NewGetBlobsOnDs3TargetSpectraS3Request(target))
            return request.BulkObjectList, err
        },
    )
}

func TestGetBlobsOnPoolSpectraS3(t *testing.T) {
    target := "Pool"
    runGetBlobsTest(
        t,
        "/_rest_/pool/" + target,
        func(client *Client) (models.BulkObjectList, error) {
            request, err := client.GetBlobsOnPoolSpectraS3(models.NewGetBlobsOnPoolSpectraS3Request(target))
            return request.BulkObjectList, err
        },
    )
}

func TestGetBlobsOnS3TargetSpectraS3(t *testing.T) {
    target := "S3Target"
    runGetBlobsTest(
        t,
        "/_rest_/s3_target/" + target,
        func(client *Client) (models.BulkObjectList, error) {
            request, err := client.GetBlobsOnS3TargetSpectraS3(models.NewGetBlobsOnS3TargetSpectraS3Request(target))
            return request.BulkObjectList, err
        },
    )
}

func TestGetBlobsOnTapeSpectraS3(t *testing.T) {
    target := "Tape"
    runGetBlobsTest(
        t,
        "/_rest_/tape/" + target,
        func(client *Client) (models.BulkObjectList, error) {
            request, err := client.GetBlobsOnTapeSpectraS3(models.NewGetBlobsOnTapeSpectraS3Request(target))
            return request.BulkObjectList, err
        },
    )
}

type getBlobsTest func(*Client) (models.BulkObjectList, error)

func runGetBlobsTest(t *testing.T, path string, callToTest getBlobsTest) {
    expectedResponse := "<Data><Object Bucket=\"default_bucket_name\" Id=\"1bd77dbf-500a-45a1-86ac-d065f026882c\" Latest=\"true\" Length=\"10\" Name=\"obj1\" Offset=\"0\" Version=\"1\"/><Object Bucket=\"default_bucket_name\" Id=\"9afa66e7-3f5a-4913-bae2-ba7c86e4c4f7\" Latest=\"true\" Length=\"10\" Name=\"obj2\" Offset=\"0\" Version=\"1\"/></Data>"

    // Create and run the mocked client.
    qp := &url.Values{ "operation": []string{"get_physical_placement"} }

    client := mockedClient(t).
        Expecting(HTTP_VERB_GET, path, qp, &http.Header{}, nil).
        Returning(200, expectedResponse, nil)

    bulkObject, err := callToTest(client)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    obj1 := bulkObject.Objects[0]
    ds3Testing.AssertNonNilStringPtr(t, "Bucket", "default_bucket_name", obj1.Bucket)
    ds3Testing.AssertNonNilStringPtr(t, "Id", "1bd77dbf-500a-45a1-86ac-d065f026882c", obj1.Id)
    ds3Testing.AssertBoolPtrIsNil(t, "InCache", obj1.InCache)
    ds3Testing.AssertBool(t, "Latest", true, obj1.Latest)
    ds3Testing.AssertInt64(t, "Length", 10, obj1.Length)
    ds3Testing.AssertNonNilStringPtr(t, "Name", "obj1", obj1.Name)
    ds3Testing.AssertInt64(t, "Offset", 0, obj1.Offset)
    ds3Testing.AssertInt64(t, "Version", 1, obj1.Version)
    if obj1.PhysicalPlacement != nil {
        t.Fatalf("Expected nil Physical Placement, but was '%v'.", obj1.PhysicalPlacement)
    }

    obj2 := bulkObject.Objects[1]
    ds3Testing.AssertNonNilStringPtr(t, "Bucket", "default_bucket_name", obj2.Bucket)
    ds3Testing.AssertNonNilStringPtr(t, "Id", "9afa66e7-3f5a-4913-bae2-ba7c86e4c4f7", obj2.Id)
    ds3Testing.AssertBoolPtrIsNil(t, "InCache", obj2.InCache)
    ds3Testing.AssertBool(t, "Latest", true, obj2.Latest)
    ds3Testing.AssertInt64(t, "Length", 10, obj2.Length)
    ds3Testing.AssertNonNilStringPtr(t, "Name", "obj2", obj2.Name)
    ds3Testing.AssertInt64(t, "Offset", 0, obj2.Offset)
    ds3Testing.AssertInt64(t, "Version", 1, obj2.Version)
    if obj2.PhysicalPlacement != nil {
        t.Fatalf("Expected nil Physical Placement, but was '%v'.", obj2.PhysicalPlacement)
    }
}

func TestGetPhysicalPlacementForObjectsSpectraS3(t *testing.T) {
    expectedRequest := "<Objects><Object Name=\"obj1\"></Object><Object Name=\"obj2\"></Object><Object Name=\"obj3\"></Object></Objects>"
    responsePayload := "<Data><AzureTargets/><Ds3Targets/><Pools/><S3Targets/><Tapes/></Data>"

    // Create and run the mocked client.
    bucketName := "BucketName"
    objectNames := []string {"obj1", "obj2", "obj3"}

    qp := &url.Values{ "operation": []string{"get_physical_placement"} }

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_PUT, "/_rest_/bucket/" + bucketName, qp, &http.Header{}, &expectedRequest).
        Returning(200, responsePayload, nil).
        GetPhysicalPlacementForObjectsSpectraS3(models.NewGetPhysicalPlacementForObjectsSpectraS3Request(bucketName, objectNames))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
}

func TestGetPhysicalPlacementForObjectsWithFullDetailsSpectraS3(t *testing.T) {
    expectedRequest := "<Objects><Object Name=\"obj1\"></Object><Object Name=\"obj2\"></Object><Object Name=\"obj3\"></Object></Objects>"
    responsePayload := "<Data><Object Bucket=\"b1\" Id=\"a2897bbd-3e0b-4c0f-83d7-29e1e7669bdd\" InCache=\"false\" Latest=\"true\" Length=\"10\" Name=\"o4\" Offset=\"0\" Version=\"1\"><PhysicalPlacement><AzureTargets/><Ds3Targets/><Pools/><S3Targets/><Tapes/></PhysicalPlacement></Object></Data>"

    // Create and run the mocked client.
    bucketName := "BucketName"
    objectNames := []string {"obj1", "obj2", "obj3"}

    qp := &url.Values{
        "operation": []string{"get_physical_placement"},
        "full_details": []string{""},
    }

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_PUT, "/_rest_/bucket/" + bucketName, qp, &http.Header{}, &expectedRequest).
        Returning(200, responsePayload, nil).
        GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3(models.NewGetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request(bucketName, objectNames))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
}

func TestMarkSuspectBlobAzureTargetsAsDegradedSpectraS3(t *testing.T) {
    runMarkSuspectBlobTest(
        t,
        "/_rest_/suspect_blob_azure_target",
        func(client *Client, ids []string) error {
            _, err := client.MarkSuspectBlobAzureTargetsAsDegradedSpectraS3(models.NewMarkSuspectBlobAzureTargetsAsDegradedSpectraS3Request(ids))
            return err
        },
    )
}

func TestMarkSuspectBlobDs3TargetsAsDegradedSpectraS3(t *testing.T) {
    runMarkSuspectBlobTest(
        t,
        "/_rest_/suspect_blob_ds3_target",
        func(client *Client, ids []string) error {
            _, err := client.MarkSuspectBlobDs3TargetsAsDegradedSpectraS3(models.NewMarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request(ids))
            return err
        },
    )
}

func TestMarkSuspectBlobPoolsAsDegradedSpectraS3(t *testing.T) {
    runMarkSuspectBlobTest(
        t,
        "/_rest_/suspect_blob_pool",
        func(client *Client, ids []string) error {
            _, err := client.MarkSuspectBlobPoolsAsDegradedSpectraS3(models.NewMarkSuspectBlobPoolsAsDegradedSpectraS3Request(ids))
            return err
        },
    )
}

func TestMarkSuspectBlobS3TargetsAsDegradedSpectraS3(t *testing.T) {
    runMarkSuspectBlobTest(
        t,
        "/_rest_/suspect_blob_s3_target",
        func(client *Client, ids []string) error {
            _, err := client.MarkSuspectBlobS3TargetsAsDegradedSpectraS3(models.NewMarkSuspectBlobS3TargetsAsDegradedSpectraS3Request(ids))
            return err
        },
    )
}

func TestMarkSuspectBlobTapesAsDegradedSpectraS3(t *testing.T) {
    runMarkSuspectBlobTest(
        t,
        "/_rest_/suspect_blob_tape",
        func(client *Client, ids []string) error {
            _, err := client.MarkSuspectBlobTapesAsDegradedSpectraS3(models.NewMarkSuspectBlobTapesAsDegradedSpectraS3Request(ids))
            return err
        },
    )
}

type markSuspectBlobTest func(*Client, []string) error

func runMarkSuspectBlobTest(t *testing.T, path string, callToTest markSuspectBlobTest) {
    expectedRequest := "<Ids><Id>id1</Id><Id>id2</Id><Id>id3</Id></Ids>"

    // Create and run the mocked client.
    ids := []string {"id1", "id2", "id3"}

    client := mockedClient(t).
        Expecting(HTTP_VERB_PUT, path, &url.Values{}, &http.Header{}, &expectedRequest).
        Returning(204, "", nil)

    err := callToTest(client, ids)

    // Check the error result.
    ds3Testing.AssertNilError(t, err)
}

func TestVerifyPhysicalPlacementForObjectsSpectraS3(t *testing.T) {
    expectedRequest := "<Objects><Object Name=\"o1\"></Object></Objects>"
    responsePayload := "<Data><AzureTargets/><Ds3Targets/><Pools/><S3Targets/><Tapes><Tape><AssignedToStorageDomain>false</AssignedToStorageDomain><AvailableRawCapacity>10000</AvailableRawCapacity><BarCode>t1</BarCode><BucketId/><DescriptionForIdentification/><EjectDate/><EjectLabel/><EjectLocation/><EjectPending/><FullOfData>false</FullOfData><Id>48d30ecb-84f1-4721-9832-7aa165a1dd77</Id><LastAccessed/><LastCheckpoint/><LastModified/><LastVerified/><PartiallyVerifiedEndOfTape/><PartitionId>76343269-c32a-4cb0-aec4-57a9dccce6ea</PartitionId><PreviousState/><SerialNumber/><State>PENDING_INSPECTION</State><StorageDomainId/><TakeOwnershipPending>false</TakeOwnershipPending><TotalRawCapacity>20000</TotalRawCapacity><Type>LTO5</Type><VerifyPending/><WriteProtected>false</WriteProtected></Tape></Tapes></Data>"

    // Create and run the mocked client.
    bucketName := "b1"
    objectNames := []string {"o1"}

    qp := &url.Values{ "operation": []string{"verify_physical_placement"} }

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/_rest_/bucket/" + bucketName, qp, &http.Header{}, &expectedRequest).
        Returning(200, responsePayload, nil).
        VerifyPhysicalPlacementForObjectsSpectraS3(models.NewVerifyPhysicalPlacementForObjectsSpectraS3Request(bucketName, objectNames))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
    ds3Testing.AssertInt(t, "Number of Tapes", 1, len(response.PhysicalPlacement.Tapes))
}

func TestVerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3(t *testing.T) {
    expectedRequest := "<Objects><Object Name=\"o1\"></Object></Objects>"
    responsePayload := "<Data><Object Bucket=\"b1\" Id=\"ad5bfa96-8356-42e5-97c7-091780f9d2a7\" InCache=\"false\" Latest=\"true\" Length=\"10\" Name=\"o1\" Offset=\"0\" Version=\"1\"><PhysicalPlacement><AzureTargets/><Ds3Targets/><Pools/><S3Targets/><Tapes><Tape><AssignedToStorageDomain>false</AssignedToStorageDomain><AvailableRawCapacity>10000</AvailableRawCapacity><BarCode>t1</BarCode><BucketId/><DescriptionForIdentification/><EjectDate/><EjectLabel/><EjectLocation/><EjectPending/><FullOfData>false</FullOfData><Id>3514700d-4d4f-4e64-8ccd-20750b5514fd</Id><LastAccessed/><LastCheckpoint/><LastModified/><LastVerified/><PartiallyVerifiedEndOfTape/><PartitionId>dc681797-927a-4eb0-9652-d19d06534e50</PartitionId><PreviousState/><SerialNumber/><State>PENDING_INSPECTION</State><StorageDomainId/><TakeOwnershipPending>false</TakeOwnershipPending><TotalRawCapacity>20000</TotalRawCapacity><Type>LTO5</Type><VerifyPending/><WriteProtected>false</WriteProtected></Tape></Tapes></PhysicalPlacement></Object></Data>"

    // Create and run the mocked client.
    bucketName := "b1"
    objectNames := []string {"o1"}

    qp := &url.Values{
        "operation": []string{"verify_physical_placement"},
        "full_details": []string{""},
    }

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_GET, "/_rest_/bucket/" + bucketName, qp, &http.Header{}, &expectedRequest).
        Returning(200, responsePayload, nil).
        VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3(models.NewVerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request(bucketName, objectNames))

    // Check the error result.
    ds3Testing.AssertNilError(t, err)

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
    ds3Testing.AssertInt(t, "Number of BulkObjects", 1, len(response.BulkObjectList.Objects))

    object := response.BulkObjectList.Objects[0]
    ds3Testing.AssertNonNilStringPtr(t, "Bucket", "b1", object.Bucket)
    ds3Testing.AssertNonNilStringPtr(t, "Id", "ad5bfa96-8356-42e5-97c7-091780f9d2a7", object.Id)
    ds3Testing.AssertNonNilBoolPtr(t, "InCache", false, object.InCache)
    ds3Testing.AssertBool(t, "Latest", true, object.Latest)
    ds3Testing.AssertInt64(t, "Length", 10, object.Length)
    ds3Testing.AssertNonNilStringPtr(t, "Name", "o1", object.Name)
    ds3Testing.AssertInt64(t, "Offset", 0, object.Offset)
    ds3Testing.AssertInt64(t, "Version", 1, object.Version)
    if object.PhysicalPlacement == nil {
        t.Fatal("Expected PhysicalPlacement to not be nil")
    }

    pp := object.PhysicalPlacement
    ds3Testing.AssertInt(t, "Number of Azure Targets", 0, len(pp.AzureTargets))
    ds3Testing.AssertInt(t, "Number of Ds3 Targets", 0, len(pp.Ds3Targets))
    ds3Testing.AssertInt(t, "Number of Pools", 0, len(pp.Pools))
    ds3Testing.AssertInt(t, "Number of S3 Targets", 0, len(pp.S3Targets))
    ds3Testing.AssertInt(t, "Number of Tapes", 1, len(pp.Tapes))

    tape := pp.Tapes[0]
    ds3Testing.AssertBool(t, "AssignedToStorageDomain", false, tape.AssignedToStorageDomain)
    ds3Testing.AssertNonNilInt64Ptr(t, "AvailableRawCapacity", 10000, tape.AvailableRawCapacity)
    ds3Testing.AssertNonNilStringPtr(t, "BarCode", "t1", tape.BarCode)
    ds3Testing.AssertStringPtrIsNil(t, "BucketId", tape.BucketId)
    ds3Testing.AssertStringPtrIsNil(t, "DescriptionForIdentification", tape.DescriptionForIdentification)
    ds3Testing.AssertStringPtrIsNil(t, "EjectDate", tape.EjectDate)
    ds3Testing.AssertStringPtrIsNil(t, "EjectLabel", tape.EjectLabel)
    ds3Testing.AssertStringPtrIsNil(t, "EjectLocation", tape.EjectLocation)
    ds3Testing.AssertStringPtrIsNil(t, "EjectPending", tape.EjectPending)
    ds3Testing.AssertBool(t, "FullOfData", false, tape.FullOfData)
    ds3Testing.AssertString(t, "Id", "3514700d-4d4f-4e64-8ccd-20750b5514fd", tape.Id)
    ds3Testing.AssertStringPtrIsNil(t, "LastAccessed", tape.LastAccessed)
    ds3Testing.AssertStringPtrIsNil(t, "LastCheckpoint", tape.LastCheckpoint)
    ds3Testing.AssertStringPtrIsNil(t, "LastModified", tape.LastModified)
    ds3Testing.AssertStringPtrIsNil(t, "LastVerified", tape.LastVerified)
    ds3Testing.AssertStringPtrIsNil(t, "PartiallyVerifiedEndOfTape", tape.PartiallyVerifiedEndOfTape)
    ds3Testing.AssertNonNilStringPtr(t, "PartitionId", "dc681797-927a-4eb0-9652-d19d06534e50", tape.PartitionId)
    if tape.PreviousState != nil {
        t.Fatalf("Expeted previous state to be 'nil' but was '%s'.", tape.PreviousState.String())
    }
    ds3Testing.AssertStringPtrIsNil(t, "SerialNumber", tape.SerialNumber)
    if tape.State != models.TAPE_STATE_PENDING_INSPECTION {
        t.Fatalf("Expected tape state 'TAPE_STATE_PENDING_INSPECTION' but got '%s'.", tape.State.String())
    }
    ds3Testing.AssertStringPtrIsNil(t, "StorageDomainId", tape.StorageDomainId)
    ds3Testing.AssertBool(t, "TakeOwnershipPending", false, tape.TakeOwnershipPending)
    ds3Testing.AssertNonNilInt64Ptr(t, "TotalRawCapacity", 20000, tape.TotalRawCapacity)
    ds3Testing.AssertString(t, "TapeType", "LTO5", tape.Type)
    if tape.VerifyPending != nil {
        t.Fatalf("Expected Verify Pending to be 'nil' but was '%s'.", tape.VerifyPending.String())
    }
    ds3Testing.AssertBool(t, "WriteProtected", false, tape.WriteProtected)
}

func TestHeadObject(t *testing.T) {
    bucketName := "bucket1"
    objectName := "obj1"

    responseHeaders := &http.Header{}
    responseHeaders.Add("x-amz-meta-key", "value")
    responseHeaders.Add("ds3-blob-checksum-type", "MD5")
    responseHeaders.Add("ds3-blob-checksum-offset-0", "4nQGNX4nyz0pi8Hvap79PQ==")
    responseHeaders.Add("ds3-blob-checksum-offset-10485760", "965Aa0/n8DlO1IwXYFh4bg==")
    responseHeaders.Add("ds3-blob-checksum-offset-20971520", "iV2OqJaXJ/jmqgRSb1HmFA==")

    response, err := mockedClient(t).
        Expecting(HTTP_VERB_HEAD, "/" + bucketName + "/" + objectName, &url.Values{}, &http.Header{}, nil).
        Returning(200, "", responseHeaders).
        HeadObject(models.NewHeadObjectRequest(bucketName, objectName))

    ds3Testing.AssertNilError(t, err)

    if response.BlobChecksumType != models.CHECKSUM_TYPE_MD5 {
        t.Fatalf("Expected checksum type to be 'MD5' but was '%s'.", response.BlobChecksumType.String())
    }

    ds3Testing.AssertInt(t, "# of blob checksums", 3, len(response.BlobChecksums))
    ds3Testing.AssertString(t, "checksum at offset '0'", "4nQGNX4nyz0pi8Hvap79PQ==", response.BlobChecksums[0])
    ds3Testing.AssertString(t, "checksum at offset '10485760'", "965Aa0/n8DlO1IwXYFh4bg==", response.BlobChecksums[10485760])
    ds3Testing.AssertString(t, "checksum at offset '20971520'", "iV2OqJaXJ/jmqgRSb1HmFA==", response.BlobChecksums[20971520])
}
