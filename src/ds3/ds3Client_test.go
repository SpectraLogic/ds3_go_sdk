package ds3

import (
    "testing"
    "strconv"
    "net/url"
    "net/http"
    "io/ioutil"
    "ds3/models"
    "ds3/networking"
    "reflect"
)

func assertNonNilBoolPtr(t *testing.T, xmlTag string, expected bool, actual *bool) {
    if actual == nil {
        t.Fatalf("Expected %s to be '%t' but was 'nil'.", xmlTag, expected)
    }
    assertBool(t, xmlTag, expected, *actual)
}

func assertBoolPtrIsNil(t *testing.T, xmlTag bool, actual *bool) {
    if actual != nil {
        t.Fatalf("Expected %s to be 'nil' but was '%t'.", xmlTag, *actual)
    }
}

func assertBool(t *testing.T, xmlTag string, expected bool, actual bool) {
    if expected != actual {
        t.Fatalf("Expected %s to be '%t' but was '%t'.", xmlTag, expected, actual)
    }
}

func assertNonNilIntPtr(t *testing.T, xmlTag string, expected int64, actual *int64) {
    if actual == nil {
        t.Fatalf("Expected %s to be '%d' but was 'nil'.", xmlTag, expected)
    }
    assertInt(t, xmlTag, expected, *actual)
}

func assertIntPtrIsNil(t *testing.T, xmlTag int64, actual *int64) {
    if actual != nil && *actual != 0 { //TODO remove default value check once nil ptr parsing is fixed
        t.Fatalf("Expected %s to be 'nil' but was '%d'.", xmlTag, *actual)
    }
}

func assertInt(t *testing.T, xmlTag string, expected int64, actual int64) {
    if expected != actual {
        t.Fatalf("Expected %s to be '%d' but was '%d'.", xmlTag, expected, actual)
    }
}

func assertNonNilStringPtr(t *testing.T, xmlTag string, expected string, actual *string) {
    if actual == nil {
        t.Fatalf("Expected %s to be '%s' but was 'nil'.", xmlTag, expected)
    }
    assertString(t, xmlTag, expected, *actual)
}

func assertStringPtrIsNil(t *testing.T, xmlTag string, actual *string) {
    if actual != nil && *actual != "" { //TODO remove empty string comparison once nil ptr parsing is fixed
        t.Fatalf("Expected %s to be 'nil' but was '%s'.", xmlTag, *actual)
    }
}

func assertString(t *testing.T, xmlTag string, expected string, actual string) {
    if expected != actual {
        t.Fatalf("Expected %s to be '%s' but was '%s'.", xmlTag, expected, actual)
    }
}

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
        Expecting(networking.GET, "/", &url.Values{}, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        GetService(models.NewGetServiceRequest())

    // Validate the error contents.
    if err != nil {
        t.Fatalf("Received an unexpected error: %s", err.Error())
    }

    // Check the response.
    if response == nil {
        t.Fatalf("Received an unexpected nil response.")
    }
    if response.ListAllMyBucketsResult.Owner.Id != "ryan" {
        t.Fatalf("Expected owner id of 'ryan' but got '%s'.", response.ListAllMyBucketsResult.Owner.Id)
    }
    if *response.ListAllMyBucketsResult.Owner.DisplayName != "ryan" {
        t.Fatalf("Expected owner display name of 'ryan' but got '%s'.", response.ListAllMyBucketsResult.Owner.DisplayName)
    }
    if len(response.ListAllMyBucketsResult.Buckets) != len(expectedBucketNames) {
        t.Fatalf("Parsed an unexpected number (%d) of buckets.", len(response.ListAllMyBucketsResult.Buckets))
    }
    for i, bucket := range response.ListAllMyBucketsResult.Buckets {
        if *bucket.Name != expectedBucketNames[i] {
            t.Errorf("%dth bucket name is incorrect (%s).", i + 1, bucket.Name)
        }
        if *bucket.CreationDate != expectedCreationDates[i] {
            t.Errorf("%dth bucket creation date is incorrect (%s).", i + 1, bucket.CreationDate)
        }
    }
}

func TestGetBadService(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.GET, "/", &url.Values{}, &http.Header{}, nil).
        Returning(400, "", nil).
        GetService(models.NewGetServiceRequest())

    // Check the response.
    if response != nil {
        t.Fatal("Expected a nil response but received a value.")
    }

    // Validate the error contents.
    if err == nil {
        t.Fatal("Expected an error but didn't get one.")
    } else {
        expectedError := "Received a status code of 400 when [200] was expected. Could not parse the response for additional information."
        actualError := err.Error()
        if actualError != expectedError {
            t.Fatalf("Expected a different error message than received: '%s'", actualError)
        }
    }
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
        Expecting(networking.GET, "/remoteTest16", &url.Values{}, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        GetBucket(models.NewGetBucketRequest("remoteTest16"))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    } else {
        if len(response.ListBucketResult.Objects) != len(keys) {
            t.Fatalf("Expected %d objects but got %d.", len(keys), len(response.ListBucketResult.Objects))
        } else {
            if *response.ListBucketResult.Name != "remoteTest16" {
                t.Fatalf("Expected bucket name 'remoteTest16' but got '%s'.", response.ListBucketResult.Name)
            }
            if *response.ListBucketResult.Prefix != "" {
                t.Fatalf("Expected empty prefix but got '%s'.", response.ListBucketResult.Prefix)
            }
            if *response.ListBucketResult.Marker != "" {
                t.Fatalf("Expected empty marker but got '%s'.", response.ListBucketResult.Marker)
            }
            if response.ListBucketResult.MaxKeys != 1000 {
                t.Fatalf("Expected max keys of 1000 but got %d.", response.ListBucketResult.MaxKeys)
            }
            if response.ListBucketResult.Truncated != false {
                t.Fatalf("Expected that the result would not be truncated, but it was.")
            }
            for i, object := range response.ListBucketResult.Objects {
                if *object.Key != keys[i] {
                    t.Fatalf("Expected key '%s' but got '%s'.", keys[i], object.Key)
                }
                if *object.LastModified != lastModifieds[i] {
                    t.Fatalf("Expected last modified '%s' but got '%s'.", lastModifieds[i], object.LastModified)
                }
                if *object.ETag != etags[i] {
                    t.Fatalf("Expected ETag '%s' but got '%s'.", etags[i], object.ETag)
                }
                if object.Size != sizes[i] {
                    t.Fatalf("Expected size %d but got %d.", sizes[i], object.Size)
                }
                if *object.StorageClass != storageClasses[i] {
                    t.Fatalf("Expected storage class '%s' but got '%s'.", storageClasses[i], object.StorageClass)
                }
                if object.Owner.Id != ids[i] {
                    t.Fatalf("Expected owner id '%s' but got '%s'.", ids[i], object.Owner.Id)
                }
                if *object.Owner.DisplayName != displayNames[i] {
                    t.Fatalf("Expected owner display name '%s' but got '%s'.", displayNames[i], object.Owner.DisplayName)
                }
            }
        }
    }
}

func TestPutBucket(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.PUT, "/bucketName", &url.Values{}, &http.Header{}, nil).
        Returning(200, "", nil).
        PutBucket(models.NewPutBucketRequest("bucketName"))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestDeleteBucket(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.DELETE, "/bucketName", &url.Values{}, &http.Header{}, nil).
        Returning(204, "", nil).
        DeleteBucket(models.NewDeleteBucketRequest("bucketName"))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

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
            Expecting(networking.DELETE, "/_rest_/folder/FolderName", queryParams, &http.Header{}, nil).
            Returning(204, "", nil).
            DeleteFolderRecursivelySpectraS3(models.NewDeleteFolderRecursivelySpectraS3Request(bucketId, folderName))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestDeleteObject(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.DELETE, "/bucketName/my/file.txt", &url.Values{}, &http.Header{}, nil).
        Returning(204, "", nil).
        DeleteObject(models.NewDeleteObjectRequest("bucketName", "my/file.txt"))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestGetBadBucket(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.GET, "/remoteTest16", &url.Values{}, &http.Header{}, nil).
        Returning(400, "", nil).
        GetBucket(models.NewGetBucketRequest("remoteTest16"))

    // Check the error result.
    if err == nil {
        t.Fatal("Expected an error but got nil.")
    } else {
        expectedError := "Received a status code of 400 when [200] was expected. Could not parse the response for additional information."
        actualError := err.Error()
        if actualError != expectedError {
            t.Fatalf("Expected a different error message than received: '%s'", actualError)
        }
    }

    // Check the response value.
    if response != nil {
        t.Fatal("Response was supposed to be nil.")
    }
}

func TestGetObject(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.GET, "/bucketName/object", &url.Values{}, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        GetObject(models.NewGetObjectRequest("bucketName", "object"))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    } else if response.Content == nil {
        t.Fatal("Response content was unexpectedly nil.")
    } else {
        defer response.Content.Close()
        bs, readErr := ioutil.ReadAll(response.Content)
        if readErr != nil {
            t.Fatalf("Unexpected error '%s'.", readErr.Error())
        } else if string(bs) != stringResponse {
            t.Fatalf("Expected '%s' but got '%s'.", stringResponse, string(bs))
        }
    }
}

func TestGetObjectRange(t *testing.T) {
    stringResponse := "object contents"
    requestHeaders := &http.Header{"Range": []string{"bytes=20-179"}}

    // Create and run the mocked client.
    request := models.NewGetObjectRequest("bucketName", "object").
            WithRange(20, 179)

    response, err := mockedClient(t).
        Expecting(networking.GET, "/bucketName/object", &url.Values{}, requestHeaders, nil).
        Returning(200, stringResponse, &http.Header{"Range": []string{"bytes=20-179"}}).
        GetObject(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    } else if response.Content == nil {
        t.Fatal("Response content was unexpectedly nil.")
    } else {
        defer response.Content.Close()
        bs, readErr := ioutil.ReadAll(response.Content)
        if readErr != nil {
            t.Fatalf("Unexpected error '%s'.", readErr.Error())
        } else if string(bs) != stringResponse {
            t.Fatalf("Expected '%s' but got '%s'.", stringResponse, string(bs))
        }
    }
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
            Expecting(networking.GET, "/_rest_/object", queryParams, &http.Header{}, nil).
            Returning(200, stringResponse, responseHeaders).
            GetObjectsDetailsSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

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
        if object.BucketId != bucketId {
            t.Fatalf("Expected bucket id to be '%s' but was '%s'.", bucketId, object.BucketId)
        }
        if object.CreationDate == nil {
            t.Fatalf("Expected creation date of '%s' but was nil.", expectedCreationDates[i])
        }
        if *object.CreationDate != expectedCreationDates[i] {
            t.Fatalf("Expected creation date of '%s' but was '%s'.", expectedCreationDates[i], *object.CreationDate)
        }
        if object.Id != expectedIds[i] {
            t.Fatalf("Expected id to be '%s' but was '%s'.", expectedIds[i], object.Id)
        }
        if object.Name == nil {
            t.Fatalf("Expected name of '%s' but was nil.", expectedNames[i])
        }
        if *object.Name != expectedNames[i] {
            t.Fatalf("Expected name of '%s' but was '%s'.", expectedNames[i], *object.Name)
        }
        if object.Type != models.S3_OBJECT_TYPE_DATA {
            t.Fatalf("Expected type of '%d' but was '%d'.", models.S3_OBJECT_TYPE_DATA, object.Type)
        }
        if object.Version != 1 {
            t.Fatalf("Expected version of 1, but was '%d'.", object.Version)
        }
    }
}

func TestGetJobToReplicateSpectraS3(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    request := models.NewGetJobToReplicateSpectraS3Request("23a876ec-2fac-4dc8-b8e6-98d6026e7f4a")

    expectedParams := &url.Values{"replicate": []string{""}}
    response, err := mockedClient(t).
        Expecting(networking.GET, "/_rest_/job/23a876ec-2fac-4dc8-b8e6-98d6026e7f4a", expectedParams, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        GetJobToReplicateSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    } else {
        if response.Content != stringResponse {
            t.Fatalf("Unexpected response: expected '%s' but was '%s'", stringResponse, response.Content)
        }
    }
}

func TestPutObject(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.PUT, "/bucketName/object", &url.Values{}, &http.Header{}, nil).
        Returning(200, "", nil).
        PutObject(models.NewPutObjectRequest(
            "bucketName",
            "object",
        networking.BuildByteReaderWithSizeDecorator([]byte(stringResponse)),
        ))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestPutObjectWithMetaData(t *testing.T) {
    stringResponse := "object contents"

    expectedMetaData := &http.Header{}
    expectedMetaData.Add("x-amz-meta-test1", "test1value")
    expectedMetaData.Add("x-amz-meta-test2", "test2value")
    expectedMetaData.Add("x-amz-meta-test2", "test2value2")

    request := models.NewPutObjectRequest(
        "bucketName",
        "object",
        networking.BuildByteReaderWithSizeDecorator([]byte(stringResponse))).
            WithMetaData("x-amz-meta-test1", "test1value").
            WithMetaData("test2", "test2value").
            WithMetaData("TEST2", "test2value2")

    if request.Header() == nil || len(*request.Header()) < 1 {
        t.Fatal("Expected there to be headers")
    }

    // Create and run the mocked client.
    response, err := mockedClient(t).
            Expecting(networking.PUT, "/bucketName/object", &url.Values{}, expectedMetaData, nil).
            Returning(200, "", nil).
            PutObject(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestBulkPut(t *testing.T) {
    runBulkTest(
        t,
        "start_bulk_put",
        func(client *Client, objects []models.Ds3Object) ([]models.Objects, error) {
            request, err := client.PutBulkJobSpectraS3(models.NewPutBulkJobSpectraS3Request("bucketName", objects))
            return request.MasterObjectList.Objects, err
        },
    )
}

func TestBulkGet(t *testing.T) {
    runBulkTest(
        t,
        "start_bulk_get",
        func(client *Client, objects []models.Ds3Object) ([]models.Objects, error) {
            request, err := client.GetBulkJobSpectraS3(models.NewGetBulkJobSpectraS3Request("bucketName", objects))
            return request.MasterObjectList.Objects, err
        },
    )
}

type bulkTest func(*Client, []models.Ds3Object) ([]models.Objects, error)

func runBulkTest(t *testing.T, operation string, callToTest bulkTest) {
    keys := []string { "file2", "file1", "file3" }
    sizes := []int64 { 1202, 256, 2523 }

    stringRequest := "<objects><object name=\"file1\" size=\"256\"></object><object name=\"file2\" size=\"1202\"></object><object name=\"file3\" size=\"2523\"></object></objects>"
    stringResponse := "<MasterObjectList><Objects><Object Name='file2' Length='1202'/><Object Name='file1' Length='256'/><Object Name='file3' Length='2523'/></Objects></MasterObjectList>"

    inputObjects := []models.Ds3Object {
        {Name: "file1", Size: 256 },
        {Name: "file2", Size: 1202 },
        {Name: "file3", Size: 2523 },
    }

    // Create and run the mocked client.
    client := mockedClient(t).
        Expecting(
        networking.PUT,
            "/_rest_/bucket/bucketName",
            &url.Values{"operation": []string{operation}},
            &http.Header{},
            &stringRequest,
        ).
        Returning(200, stringResponse, nil)
    response, err := callToTest(client, inputObjects)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

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
        if *obj.Name != keys[i] {
            t.Errorf("Expected key %s but got %s.", keys[i], obj.Name)
        }
        if obj.Length != sizes[i] {
            t.Errorf("Expected size %d but got %d.", sizes[i], obj.Length)
        }
    }
}

func TestInitiateMultipart(t *testing.T) {
    stringResponse := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><InitiateMultipartUploadResult xmlns=\"http://s3.amazonaws.com/doc/2006-03-01/\"><Bucket>example-bucket</Bucket><Key>example-object</Key><UploadId>VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA</UploadId></InitiateMultipartUploadResult>"

    // Create and run the mocked client.
    qs := &url.Values{"uploads": []string{""}}
    response, err := mockedClient(t).
        Expecting(networking.POST, "/bucketName/object", qs, &http.Header{}, nil).
        Returning(200, stringResponse, nil).
        InitiateMultiPartUpload(models.NewInitiateMultiPartUploadRequest(
            "bucketName",
            "object",
        ))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
    if response.InitiateMultipartUploadResult.Bucket == nil {
        t.Fatal("Expected bucket 'example-bucket' but got 'nil'.")
    }
    if *response.InitiateMultipartUploadResult.Bucket != "example-bucket" {
        t.Fatalf("Expected bucket 'example-bucket' but got '%s'.", response.InitiateMultipartUploadResult.Bucket)
    }
    if *response.InitiateMultipartUploadResult.Key != "example-object" {
        t.Fatalf("Expected key 'example-object' but got '%s'.", response.InitiateMultipartUploadResult.Key)
    }
    if *response.InitiateMultipartUploadResult.UploadId != "VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA" {
        t.Fatalf("Expected upload id 'VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA' but got '%s'.", response.InitiateMultipartUploadResult.UploadId)
    }
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
        Expecting(networking.PUT, "/bucketName/object", qs, &http.Header{}, nil).
        Returning(200, "", responseHeaders).
        PutMultiPartUploadPart(models.NewPutMultiPartUploadPartRequest(
            "bucketName",
            "object",
            networking.BuildByteReaderWithSizeDecorator([]byte(content)),
            partNumber,
            uploadId,
        ))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }

    assertString(t, "etag header", eTag, response.Headers.Get("etag"))
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
        Expecting(networking.POST, "/bucketName/object", qs, &http.Header{}, &expectedRequest).
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
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatalf("Response was unexpectedly nil.")
    }
    if response.CompleteMultipartUploadResult.Location == nil {
        t.Fatalf("Expected location '%s' but got 'nil'.", location)
    }
    if *response.CompleteMultipartUploadResult.Location != location {
        t.Fatalf("Expected location '%s' but got '%s'.", location, response.CompleteMultipartUploadResult.Location)
    }
    if *response.CompleteMultipartUploadResult.Bucket != bucket {
        t.Fatalf("Expected bucket '%s' but got '%s'.", bucket, response.CompleteMultipartUploadResult.Bucket)
    }
    if *response.CompleteMultipartUploadResult.Key != key {
        t.Fatalf("Expected key '%s' but got '%s'.", key, response.CompleteMultipartUploadResult.Key)
    }
    if response.CompleteMultipartUploadResult.ETag == nil {
        t.Fatalf("Expected etag '%s' but got 'nil'.", etag)
    }
    if *response.CompleteMultipartUploadResult.ETag != etag {
        t.Fatalf("Expected etag '%s' but got '%s'.", etag, response.CompleteMultipartUploadResult.ETag)
    }
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
        Expecting(networking.POST, "/bucketName", qs, &http.Header{}, &expectedRequest).
        Returning(200, expectedResponse, &http.Header{}).
        DeleteObjects(models.NewDeleteObjectsRequest(bucket, objectNames))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

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
            Expecting(networking.PUT, "/_rest_/job_chunk/203f6886-b058-4f7c-a012-8779176453b1", qp, &http.Header{}, nil).
            Returning(200, responseString, nil).
            AllocateJobChunkSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    if response.Objects.ChunkId != chunkId {
        t.Fatalf("Expected chunk id '%s' but got '%s'.", chunkId, response.Objects.ChunkId)
    }
    if response.Objects.ChunkNumber != 3 {
        t.Fatalf("Expected chunk number 3 but got '%s'.", response.Objects.ChunkNumber)
    }
    if response.Objects.NodeId == nil {
        t.Fatalf("Expected node id '%s' but got nil.", nodeId)
    }
    if *response.Objects.NodeId != nodeId {
        t.Fatalf("Expected node id '%s' but got '%s'.", nodeId, *response.Objects.NodeId)
    }

    expectedNames := []string{"client00obj000004-8000000", "client00obj000004-8000000", "client00obj000003-8000000", "client00obj000003-8000000"}
    expectedOffsets := []int64{0, 5368709120, 5368709120, 0}
    expectedLengths := []int64{5368709120, 2823290880, 2823290880, 5368709120}

    if len(response.Objects.Objects) != len(expectedNames) {
        t.Fatalf("Expected number of objects '%d' but got '%d'.", len(response.Objects.Objects), len(expectedNames))
    }
    for i, obj := range response.Objects.Objects {
        if obj.Name == nil {
            t.Errorf("Expected name '%s' but got nil.", obj.Name)
        }
        if *obj.Name != expectedNames[i] {
            t.Errorf("Expected name '%s' but got '%s'.", obj.Name, expectedNames[i])
        }
        if obj.InCache == nil {
            t.Error("Expected in-cache to be true, but was nil.")
        }
        if !*obj.InCache {
            t.Errorf("Expected in-cache to be true, but was '%t'.", *obj.InCache)
        }
        if obj.Offset != expectedOffsets[i] {
            t.Errorf("Expected offset '%d' but got '%d'.", obj.Offset, expectedOffsets[i])
        }
        if obj.Length != expectedLengths[i] {
            t.Errorf("Expected length '%d' but got '%d'.", obj.Length, expectedLengths[i])
        }
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
    if masterObjectList.BucketName == nil {
        t.Fatalf("Expected bucket name '%s' but got nil.", bucketName)
    }
    if *masterObjectList.BucketName != bucketName {
        t.Fatalf("Expected bucket name '%s' but got '%s'.", bucketName, masterObjectList.BucketName)
    }
    if masterObjectList.JobId != test_master_object_list_id {
        t.Fatalf("Expected job id '%s' but got '%s'.", test_master_object_list_id, masterObjectList.JobId)
    }
    if masterObjectList.Priority != models.PRIORITY_NORMAL {
        t.Fatalf("Expected priority '%d' but got '%d'.", models.PRIORITY_NORMAL, masterObjectList.Priority)
    }
    if masterObjectList.RequestType != models.JOB_REQUEST_TYPE_GET {
        t.Fatalf("Expected request type '%d' but got '%d'.", models.JOB_REQUEST_TYPE_GET, masterObjectList.RequestType)
    }
    if masterObjectList.StartDate != startDate {
        t.Fatalf("Expected start date '%s' but got '%s'.", startDate, masterObjectList.StartDate)
    }

    // verify nodes
    if len(masterObjectList.Nodes) != 2 {
        t.Fatalf("Expected there to be 2 nodes but got '%d'.", len(masterObjectList.Nodes))
    }

    node1 := masterObjectList.Nodes[0]
    if node1.EndPoint == nil || *node1.EndPoint != "10.1.18.12" {
        t.Fatalf("Expected end point 10.1.18.12, but got '%s'.", node1.EndPoint)
    }
    if node1.HttpPort == nil || *node1.HttpPort != 80 {
        t.Fatalf("Expected http port 80 but got '%d'.", node1.HttpPort)
    }
    if node1.HttpsPort == nil || *node1.HttpsPort != 443 {
        t.Fatalf("Expected https port 443 but got '%d'.", node1.HttpsPort)
    }
    if node1.Id != "a02053b9-0147-11e4-8d6a-002590c1177c" {
        t.Fatalf("Expected id a02053b9-0147-11e4-8d6a-002590c1177c but got '%s'.", node1.Id)
    }
    node2 := masterObjectList.Nodes[1]
    if node2.EndPoint == nil || *node2.EndPoint != "10.1.18.13" {
        t.Fatalf("Expected end point 10.1.18.13, but got '%s'.", node2.EndPoint)
    }
    if node2.HttpPort != nil {
        t.Fatalf("Expected http port to be nil but got '%d'.", node2.HttpPort)
    }
    if node2.HttpsPort == nil || *node2.HttpsPort != 443 {
        t.Fatalf("Expected https port 443 but got '%d'.", node2.HttpsPort)
    }
    if node2.Id != "95e97010-8e70-4733-926c-aeeb21796848" {
        t.Fatalf("Expected id 95e97010-8e70-4733-926c-aeeb21796848 but got '%s'.", node2.Id)
    }

    // verify objects
    if len(masterObjectList.Objects) != 2 {
        t.Fatalf("Expected 2 object but got '%d'.", len(masterObjectList.Objects))
    }

    obj1 := masterObjectList.Objects[0]
    if obj1.ChunkId != "f58370c2-2538-4e78-a9f8-e4d2676bdf44" {
        t.Fatalf("Expected chunk id f58370c2-2538-4e78-a9f8-e4d2676bdf44 but got '%s'.", obj1.ChunkId)
    }
    if obj1.ChunkNumber != 0 {
        t.Fatalf("Expected chunk number 0 but got '%d'.", obj1.ChunkNumber)
    }
    if obj1.NodeId == nil || *obj1.NodeId != "a02053b9-0147-11e4-8d6a-002590c1177c" {
        t.Fatalf("Expected node id a02053b9-0147-11e4-8d6a-002590c1177c but got '%s'.", *obj1.NodeId)
    }

    if len(obj1.Objects) != 2 {
        t.Fatalf("Expected 2 bulk objects but got '%d'.", len(obj1.Objects))
    }

    obj1BulkObj1 := obj1.Objects[0]
    if obj1BulkObj1.Name == nil || *obj1BulkObj1.Name != "client00obj000004-8000000" {
        t.Fatalf("Expected name 'client00obj000004-8000000' but got '%s'.", obj1BulkObj1.Name)
    }
    if obj1BulkObj1.InCache == nil || !*obj1BulkObj1.InCache {
        t.Fatalf("Expected in cache 'true' but got '%t'", *obj1BulkObj1.InCache)
    }
    if obj1BulkObj1.Length != 5368709120 {
        t.Fatalf("Expected length '5368709120' but got '%d'.", obj1BulkObj1.Length)
    }
    if obj1BulkObj1.Offset != 0 {
        t.Fatalf("Expected offset '0' but got '%d'.", obj1BulkObj1.Offset)
    }

    obj1BulkObj2 := obj1.Objects[1]
    if obj1BulkObj2.Name == nil || *obj1BulkObj2.Name != "client00obj000004-8000000" {
        t.Fatalf("Expected name 'client00obj000004-8000000' but got '%s'.", obj1BulkObj2.Name)
    }
    if obj1BulkObj2.InCache == nil || !*obj1BulkObj2.InCache {
        t.Fatalf("Expected in cache 'true' but got '%t'", *obj1BulkObj2.InCache)
    }
    if obj1BulkObj2.Length != 2823290880 {
        t.Fatalf("Expected length '2823290880' but got '%d'.", obj1BulkObj2.Length)
    }
    if obj1BulkObj2.Offset != 5368709120 {
        t.Fatalf("Expected offset '5368709120' but got '%d'.", obj1BulkObj2.Offset)
    }

    obj2 := masterObjectList.Objects[1]
    if obj2.ChunkId != "4137d768-25bb-4942-9d36-b92dfbe75e01" {
        t.Fatalf("Expected chunk id 4137d768-25bb-4942-9d36-b92dfbe75e01 but got '%s'.", obj2.ChunkId)
    }
    if obj2.ChunkNumber != 1 {
        t.Fatalf("Expected chunk number 1 but got '%d'.", obj2.ChunkNumber)
    }
    if obj2.NodeId == nil || *obj2.NodeId != "95e97010-8e70-4733-926c-aeeb21796848" {
        t.Fatalf("Expected node id 95e97010-8e70-4733-926c-aeeb21796848 but got '%s'.", *obj2.NodeId)
    }
    if len(obj2.Objects) != 2 {
        t.Fatalf("Expected 2 bulk objects but got '%d'.", len(obj2.Objects))
    }

    obj2BulkObj1 := obj2.Objects[0]
    if obj2BulkObj1.Name == nil || *obj2BulkObj1.Name != "client00obj000008-8000000" {
        t.Fatalf("Expected name 'client00obj000008-8000000' but got '%s'.", obj2BulkObj1.Name)
    }
    if obj2BulkObj1.InCache == nil || !*obj2BulkObj1.InCache {
        t.Fatalf("Expected in cache 'true' but got '%t'", *obj2BulkObj1.InCache)
    }
    if obj2BulkObj1.Length != 2823290880 {
        t.Fatalf("Expected length '2823290880' but got '%d'.", obj2BulkObj1.Length)
    }
    if obj2BulkObj1.Offset != 5368709120 {
        t.Fatalf("Expected offset '5368709120' but got '%d'.", obj2BulkObj1.Offset)
    }

    obj2BulkObj2 := obj2.Objects[1]
    if obj2BulkObj2.Name == nil || *obj2BulkObj2.Name != "client00obj000008-8000000" {
        t.Fatalf("Expected name 'client00obj000008-8000000' but got '%s'.", obj2BulkObj2.Name)
    }
    if obj2BulkObj2.InCache == nil || !*obj2BulkObj2.InCache {
        t.Fatalf("Expected in cache 'true' but got '%t'", *obj2BulkObj2.InCache)
    }
    if obj2BulkObj2.Length != 5368709120 {
        t.Fatalf("Expected length '5368709120' but got '%d'.", obj2BulkObj2.Length)
    }
    if obj2BulkObj2.Offset != 0 {
        t.Fatalf("Expected offset '0' but got '%d'.", obj2BulkObj2.Offset)
    }
}

func TestGetJobChunksReadyForClientProcessingSpectraS3(t *testing.T) {
    qp := &url.Values{"job": []string{test_master_object_list_id}}

    request := models.NewGetJobChunksReadyForClientProcessingSpectraS3Request(test_master_object_list_id)
    response, err := mockedClient(t).
            Expecting(networking.GET, "/_rest_/job_chunk", qp, &http.Header{}, nil).
            Returning(200, test_master_object_list_xml, nil).
            GetJobChunksReadyForClientProcessingSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    verifyMasterObjectList(t, response.MasterObjectList)
}

func TestGetJobSpectraS3(t *testing.T) {
    request := models.NewGetJobSpectraS3Request(test_master_object_list_id)
    response, err := mockedClient(t).
            Expecting(networking.GET, "/_rest_/job/1a85e743-ec8f-4789-afec-97e587a26936", &url.Values{}, &http.Header{}, nil).
            Returning(200, test_master_object_list_xml, nil).
            GetJobSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    verifyMasterObjectList(t, response.MasterObjectList)
}

func TestModifyJobSpectraS3(t *testing.T) {
    request := models.NewModifyJobSpectraS3Request(test_master_object_list_id)
    response, err := mockedClient(t).
            Expecting(networking.PUT, "/_rest_/job/1a85e743-ec8f-4789-afec-97e587a26936", &url.Values{}, &http.Header{}, nil).
            Returning(200, test_master_object_list_xml, nil).
            ModifyJobSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

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
            Expecting(networking.GET, "/_rest_/job", &url.Values{}, &http.Header{}, nil).
            Returning(200, responseString, nil).
            GetJobsSpectraS3(models.NewGetJobsSpectraS3Request())

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    if len(response.JobList.Jobs) != 2 {
        t.Fatalf("Expected 2 jobs but got '%d'.", len(response.JobList.Jobs))
    }

    job1 := response.JobList.Jobs[0]
    if job1.BucketName == nil || *job1.BucketName != "bucket_1" {
        t.Fatalf("Expected name 'bucket_1' but got '%s'.", *job1.BucketName)
    }
    if job1.CachedSizeInBytes != 69880 {
        t.Fatalf("Expected cached size in bytes '69880' but got '%d'.", job1.CachedSizeInBytes)
    }
    if job1.ChunkClientProcessingOrderGuarantee != models.JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER {
        t.Fatalf("Expected chunk client processing order guarantee '%d' but got '%d'.",
            models.JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER,
            job1.ChunkClientProcessingOrderGuarantee,
        )
    }
    if job1.CompletedSizeInBytes != 0 {
        t.Fatalf("Expected completed size in bytes '0' but got '%d'.", job1.CompletedSizeInBytes)
    }
    if job1.JobId != "0807ff11-a9f6-4d55-bb92-b452c1bb00c7" {
        t.Fatalf("Expected job id '0807ff11-a9f6-4d55-bb92-b452c1bb00c7' but got '%s'.", job1.JobId)
    }
    if job1.OriginalSizeInBytes != 69880 {
        t.Fatalf("Expected original size in bytes '69880' but got '%d'.", job1.OriginalSizeInBytes)
    }
    if job1.Priority != models.PRIORITY_NORMAL {
        t.Fatalf("Expected priority '%d' but got '%d'.", models.PRIORITY_NORMAL, job1.Priority)
    }
    if job1.RequestType != models.JOB_REQUEST_TYPE_PUT {
        t.Fatalf("Expected request type '%d' but got '%d'.", models.JOB_REQUEST_TYPE_PUT, job1.RequestType)
    }
    if job1.StartDate != "2014-09-04T17:23:45.000Z" {
        t.Fatalf("Expected start date '2014-09-04T17:23:45.000Z' but got '%s'.", job1.StartDate)
    }
    if job1.UserId != "a7d3eff9-e6d2-4e37-8a0b-84e76211a18a" {
        t.Fatalf("Expected user id 'a7d3eff9-e6d2-4e37-8a0b-84e76211a18a' but got '%s'.", job1.UserId)
    }
    if job1.UserName == nil || *job1.UserName != "spectra" {
        t.Fatalf("Expected user name 'spectra' but got '%s'.", *job1.UserName)
    }
    if len(job1.Nodes) != 1 {
        t.Fatalf("Expected '1' nodes but got '%d'.", len(job1.Nodes))
    }

    job1node := job1.Nodes[0]
    if job1node.EndPoint == nil || *job1node.EndPoint != "10.10.10.10" {
        t.Fatalf("Expected end point '10.10.10.10' but got '%s'.", *job1node.EndPoint)
    }
    if job1node.HttpPort == nil || *job1node.HttpPort != 80{
        t.Fatalf("Expected http port '80' but got '%d'.", *job1node.HttpPort)
    }
    if job1node.HttpsPort == nil || *job1node.HttpsPort != 443 {
        t.Fatalf("Expected https port '443' but got '%d'.", *job1node.HttpsPort)
    }
    if job1node.Id != "edb8cc38-32f2-11e4-bce1-080027ecf0d4" {
        t.Fatalf("Expected id 'edb8cc38-32f2-11e4-bce1-080027ecf0d4' but got '%s'.", job1node.Id)
    }

    job2 := response.JobList.Jobs[1]
    if job2.BucketName == nil || *job2.BucketName != "bucket_2" {
        t.Fatalf("Expected name 'bucket_2' but got '%s'.", *job2.BucketName)
    }
    if job2.CachedSizeInBytes != 0 {
        t.Fatalf("Expected cached size in bytes '0' but got '%d'.", job2.CachedSizeInBytes)
    }
    if job2.ChunkClientProcessingOrderGuarantee != models.JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER {
        t.Fatalf("Expected chunk client processing order guarantee '%d' but got '%d'.",
            models.JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER,
            job2.ChunkClientProcessingOrderGuarantee,
        )
    }
    if job2.CompletedSizeInBytes != 0 {
        t.Fatalf("Expected completed size in bytes '0' but got '%d'.", job2.CompletedSizeInBytes)
    }
    if job2.JobId != "c18554ba-e3a8-4905-91fd-3e6eec71bf45" {
        t.Fatalf("Expected job id 'c18554ba-e3a8-4905-91fd-3e6eec71bf45' but got '%s'.", job2.JobId)
    }
    if job2.OriginalSizeInBytes != 69880 {
        t.Fatalf("Expected original size in bytes '69880' but got '%d'.", job2.OriginalSizeInBytes)
    }
    if job2.Priority != models.PRIORITY_HIGH {
        t.Fatalf("Expected priority '%d' but got '%d'.", models.PRIORITY_HIGH, job2.Priority)
    }
    if job2.RequestType != models.JOB_REQUEST_TYPE_GET {
        t.Fatalf("Expected request type '%d' but got '%d'.", models.JOB_REQUEST_TYPE_GET, job2.RequestType)
    }
    if job2.StartDate != "2014-09-04T17:24:04.000Z" {
        t.Fatalf("Expected start date '2014-09-04T17:24:04.000Z' but got '%s'.", job2.StartDate)
    }
    if job2.UserId != "a7d3eff9-e6d2-4e37-8a0b-84e76211a18a" {
        t.Fatalf("Expected user id 'a7d3eff9-e6d2-4e37-8a0b-84e76211a18a' but got '%s'.", job2.UserId)
    }
    if job2.UserName == nil || *job1.UserName != "spectra" {
        t.Fatalf("Expected user name 'spectra' but got '%s'.", *job2.UserName)
    }
    if len(job2.Nodes) != 1 {
        t.Fatalf("Expected '1' nodes but got '%d'.", len(job2.Nodes))
    }

    job2node := job2.Nodes[0]
    if job2node.EndPoint == nil || *job2node.EndPoint != "10.10.10.10" {
        t.Fatalf("Expected end point '10.10.10.10' but got '%s'.", *job2node.EndPoint)
    }
    if job2node.HttpPort == nil || *job2node.HttpPort != 80 {
        t.Fatalf("Expected http port '80' but got '%d'.", *job2node.HttpPort)
    }
    if job2node.HttpsPort == nil || *job2node.HttpsPort != 443 {
        t.Fatalf("Expected https port '443' but got '%d'.", *job2node.HttpsPort)
    }
    if job2node.Id != "edb8cc38-32f2-11e4-bce1-080027ecf0d4" {
        t.Fatalf("Expected id 'edb8cc38-32f2-11e4-bce1-080027ecf0d4' but got '%s'.", job2node.Id)
    }
}

func TestCancelJobSpectraS3(t *testing.T) {
    jobId := "1a85e743-ec8f-4789-afec-97e587a26936"
    qp := &url.Values{"force": []string{""}}
    request := models.NewCancelJobSpectraS3Request(jobId)
    response, err := mockedClient(t).
            Expecting(networking.DELETE, "/_rest_/job/1a85e743-ec8f-4789-afec-97e587a26936", qp, &http.Header{}, nil).
            Returning(204, "", nil).
            CancelJobSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestDeleteTapeDriveSpectraS3(t *testing.T) {
    jobId := "1a85e743-ec8f-4789-afec-97e587a26936"

    request := models.NewDeleteTapeDriveSpectraS3Request(jobId)
    response, err := mockedClient(t).
            Expecting(networking.DELETE, "/_rest_/tape_drive/1a85e743-ec8f-4789-afec-97e587a26936", &url.Values{}, &http.Header{}, nil).
            Returning(204, "", nil).
            DeleteTapeDriveSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestDeleteTapePartitionSpectraS3(t *testing.T) {
    id := "1a85e743-ec8f-4789-afec-97e587a26936"

    request := models.NewDeleteTapePartitionSpectraS3Request(id)
    response, err := mockedClient(t).
            Expecting(networking.DELETE, "/_rest_/tape_partition/1a85e743-ec8f-4789-afec-97e587a26936", &url.Values{}, &http.Header{}, nil).
            Returning(204, "", nil).
            DeleteTapePartitionSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
}

func TestVerifySystemHealthSpectraS3(t *testing.T) {
    responsePayload := "<Data><MsRequiredToVerifyDataPlannerHealth>10</MsRequiredToVerifyDataPlannerHealth></Data>"
    response, err := mockedClient(t).
            Expecting(networking.GET, "/_rest_/system_health", &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, nil).
            VerifySystemHealthSpectraS3(models.NewVerifySystemHealthSpectraS3Request())

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    if response.HealthVerificationResult.MsRequiredToVerifyDataPlannerHealth != 10 {
        t.Fatalf("Expected ms requred to verify data planner health '10' but was '%d'.", response.HealthVerificationResult.MsRequiredToVerifyDataPlannerHealth)
    }
}

func TestGetSystemInformationSpectraS3(t *testing.T) {
    responsePayload := "<Data>" +
            "<ApiVersion>518B3F2A95B71AC7325EFB12B2937376.15F3CC0489CBCD4648ECFF0FBF371B8A</ApiVersion>" +
            "<BuildInformation><Branch/><Revision/><Version/></BuildInformation>" +
            "<SerialNumber>UNKNOWN</SerialNumber>" +
            "</Data>"
    response, err := mockedClient(t).
            Expecting(networking.GET, "/_rest_/system_information", &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, nil).
            GetSystemInformationSpectraS3(models.NewGetSystemInformationSpectraS3Request())

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }
    if response.SystemInformation.ApiVersion == nil || *response.SystemInformation.ApiVersion != "518B3F2A95B71AC7325EFB12B2937376.15F3CC0489CBCD4648ECFF0FBF371B8A" {
        t.Fatalf("Expected api version '518B3F2A95B71AC7325EFB12B2937376.15F3CC0489CBCD4648ECFF0FBF371B8A' but was '%s'.", response.SystemInformation.ApiVersion)
    }
    if response.SystemInformation.SerialNumber == nil || *response.SystemInformation.SerialNumber != "UNKNOWN" {
        t.Fatalf("Expected serial number 'UNKNOWN' but was '%s'.", *response.SystemInformation.SerialNumber)
    }
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
            Expecting(networking.GET, "/_rest_/tape_library", &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, responseHeaders).
            GetTapeLibrariesSpectraS3(models.NewGetTapeLibrariesSpectraS3Request())

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    if len(response.TapeLibraryList.TapeLibraries) != 2 {
        t.Fatalf("Expected '2' tape libraries but got '%d'.", len(response.TapeLibraryList.TapeLibraries))
    }

    lib1 := response.TapeLibraryList.TapeLibraries[0]
    if lib1.Id != "f4dae25d-e52a-4430-82bd-525e4f15493c" {
        t.Fatalf("Expected id 'f4dae25d-e52a-4430-82bd-525e4f15493c' but was '%s'.", lib1.Id)
    }
    if lib1.ManagementUrl == nil || *lib1.ManagementUrl != "a" {
        t.Fatalf("Expected management url 'a' but was '%s'.", *lib1.ManagementUrl)
    }
    if lib1.Name == nil || *lib1.Name != "test library" {
        t.Fatalf("Expected name 'test library' but was '%s'.", lib1.Name)
    }
    if lib1.SerialNumber == nil || *lib1.SerialNumber != "test library" {
        t.Fatalf("Expected serial number 'test library' but was '%s'.", *lib1.SerialNumber)
    }

    lib2 := response.TapeLibraryList.TapeLibraries[1]
    if lib2.Id != "82bdab72-d79a-4b43-95d7-f2c16cd9aa45" {
        t.Fatalf("Expected id '82bdab72-d79a-4b43-95d7-f2c16cd9aa45' but was '%s'.", lib2.Id)
    }
    if lib2.ManagementUrl == nil || *lib2.ManagementUrl != "a" {
        t.Fatalf("Expected management url 'a' but was '%s'.", *lib2.ManagementUrl)
    }
    if lib2.Name == nil || *lib2.Name != "test library 2" {
        t.Fatalf("Expected name 'test library 2' but was '%s'.", lib2.Name)
    }
    if lib2.SerialNumber == nil || *lib2.SerialNumber != "test library 2" {
        t.Fatalf("Expected serial number 'test library 2' but was '%s'.", *lib2.SerialNumber)
    }

    assertString(t, "Page-Truncated header", "2", response.Headers.Get("Page-Truncated"))
    assertString(t, "Total-Result-Count header", "3", response.Headers.Get("Total-Result-Count"))
}

func TestGetTapeLibrarySpectraS3(t *testing.T) {
    responsePayload := "<Data><Id>e23030e5-9b8d-4594-bdd1-15d3c45abb9f</Id><ManagementUrl>a</ManagementUrl><Name>125ca16e-60e3-43b2-a26f-0bc81843745f</Name><SerialNumber>test library</SerialNumber></Data>"

    id := "e23030e5-9b8d-4594-bdd1-15d3c45abb9f"

    response, err := mockedClient(t).
            Expecting(networking.GET, "/_rest_/tape_library/" + id, &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, nil).
            GetTapeLibrarySpectraS3(models.NewGetTapeLibrarySpectraS3Request(id))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    library := response.TapeLibrary
    if library.Id != id {
        t.Fatalf("Expected id '%s' but was '%s'.", id, library.Id)
    }
    if library.ManagementUrl == nil || *library.ManagementUrl != "a" {
        t.Fatalf("Expected management url 'a' but was '%s'.", *library.ManagementUrl)
    }
    if library.Name == nil || *library.Name != "125ca16e-60e3-43b2-a26f-0bc81843745f" {
        t.Fatalf("Expected name '125ca16e-60e3-43b2-a26f-0bc81843745f' but was '%s'.", library.Name)
    }
    if library.SerialNumber == nil || *library.SerialNumber != "test library" {
        t.Fatalf("Expected serial number 'test library' but was '%s'.", *library.SerialNumber)
    }
}

func TestGetTapeDriveSpectraS3(t *testing.T) {
    responsePayload := "<Data><ErrorMessage/><ForceTapeRemoval>false</ForceTapeRemoval><Id>ff5df6c8-7e24-4e4f-815d-a8a1a4cddc98</Id><PartitionId>ca69b187-47cf-425e-b92f-c09bacc7d3b3</PartitionId><SerialNumber>test tape drive</SerialNumber><State>NORMAL</State><TapeId>0ea07c32-8ff6-443f-b7c8-420667b0df84</TapeId><Type>UNKNOWN</Type></Data>"

    id := "ff5df6c8-7e24-4e4f-815d-a8a1a4cddc98"

    response, err := mockedClient(t).
            Expecting(networking.GET, "/_rest_/tape_drive/" + id, &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, nil).
            GetTapeDriveSpectraS3(models.NewGetTapeDriveSpectraS3Request(id))

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    drive := response.TapeDrive
    if drive.ErrorMessage != nil && *drive.ErrorMessage != "" {
        t.Fatalf("Expected error message to be 'nil' but was '%s'.", *drive.ErrorMessage)
    }
    if drive.ForceTapeRemoval {
        t.Fatalf("Expected force tap removal to be 'false' but was '%t'.", drive.ForceTapeRemoval)
    }
    if drive.Id != id {
        t.Fatalf("Expected id '%s' but was '%s'.", id, drive.Id)
    }
    if drive.PartitionId != "ca69b187-47cf-425e-b92f-c09bacc7d3b3" {
        t.Fatalf("Expected partition id 'ca69b187-47cf-425e-b92f-c09bacc7d3b3' but was '%s'.", drive.PartitionId)
    }
    if drive.SerialNumber == nil || *drive.SerialNumber != "test tape drive" {
        t.Fatalf("Expected serial number 'test tape drive' but was '%s'.", *drive.SerialNumber)
    }
    if drive.State != models.TAPE_DRIVE_STATE_NORMAL {
        t.Fatalf("Expected drive stat '%d' but was '%d'.", models.TAPE_DRIVE_STATE_NORMAL, drive.State)
    }
    if drive.TapeId == nil || *drive.TapeId != "0ea07c32-8ff6-443f-b7c8-420667b0df84" {
        t.Fatalf("Expected tape id '0ea07c32-8ff6-443f-b7c8-420667b0df84' but was '%s'.", *drive.TapeId)
    }
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
            Expecting(networking.GET, "/_rest_/tape", &url.Values{}, &http.Header{}, nil).
            Returning(200, responsePayload, responseHeaders).
            GetTapesSpectraS3(models.NewGetTapesSpectraS3Request())

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    if len(response.TapeList.Tapes) != 1 {
        t.Fatalf("Expected '1' tapes but got '%d'.", len(response.TapeList.Tapes))
    }

    tape := response.TapeList.Tapes[0]
    assertBool(t, "AssignedToStorageDomain", false, tape.AssignedToStorageDomain)
    assertNonNilIntPtr(t, "AvailableRawCapacity", 2408082046976, tape.AvailableRawCapacity)
    assertNonNilStringPtr(t, "BarCode", "101000L6", tape.BarCode)
    assertStringPtrIsNil(t, "BucketId", tape.BucketId)
    assertStringPtrIsNil(t, "DescriptionForIdentification", tape.DescriptionForIdentification)
    assertStringPtrIsNil(t, "EjectDate", tape.EjectDate)
    assertStringPtrIsNil(t, "EjectLabel", tape.EjectLabel)
    assertStringPtrIsNil(t, "EjectLocation", tape.EjectLocation)
    assertStringPtrIsNil(t, "EjectPending", tape.EjectPending)
    assertBool(t, "FullOfData", false, tape.FullOfData)
    assertString(t, "Id", "c7c431df-f95d-4533-b350-ffd7a8a5caac", tape.Id)
    assertNonNilStringPtr(t, "LastAccessed", "2015-09-04T06:53:08.236", tape.LastAccessed)
    assertNonNilStringPtr(t, "LastCheckpoint", "eb77ea67-3c83-47ec-8714-cd46a97dc392:2", tape.LastCheckpoint)
    assertNonNilStringPtr(t, "LastModified", "2015-08-21T16:14:30.714", tape.LastModified)
    assertStringPtrIsNil(t, "LastVerified", tape.LastVerified)
    assertNonNilStringPtr(t, "PartitionId", "4f8a5cbb-9837-41d9-afd1-cebed41f18f7", tape.PartitionId)
    if *tape.PreviousState != models.UNDEFINED {
        t.Fatalf("Expected previous state '%d' but was '%d'.", models.UNDEFINED, *tape.PreviousState)
    }
    assertNonNilStringPtr(t, "SerialNumber", "HP-W130501213", tape.SerialNumber)
    assertString(t, "State", models.TAPE_STATE_NORMAL.String(), tape.State.String())
    assertNonNilIntPtr(t, "TotalRawCapacity", 2408088338432, tape.TotalRawCapacity)
    assertString(t, "Type", models.TAPE_TYPE_LTO6.String(), tape.Type.String())
    assertBool(t, "WriteProtected", false, tape.WriteProtected)

    assertString(t, "Page-Truncated header", "2", response.Headers.Get("Page-Truncated"))
    assertString(t, "Total-Result-Count header", "3", response.Headers.Get("Total-Result-Count"))
}

func TestDeletePermanentlyLostTapeSpectraS3(t *testing.T) {
    id := "1a85e743-ec8f-4789-afec-97e587a26936"

    request := models.NewDeletePermanentlyLostTapeSpectraS3Request(id)
    response, err := mockedClient(t).
        Expecting(networking.DELETE, "/_rest_/tape/" + id, &url.Values{}, &http.Header{}, nil).
        Returning(204, "", nil).
        DeletePermanentlyLostTapeSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

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
        Expecting(networking.GET, "/_rest_/tape/" + id, &url.Values{}, &http.Header{}, nil).
        Returning(200, responsePayload, nil).
        GetTapeSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Fatal("Response was unexpectedly nil.")
    }

    tape := response.Tape
    assertBool(t, "AssignedToStorageDomain", false, tape.AssignedToStorageDomain)
    assertNonNilIntPtr(t, "AvailableRawCapacity", 2408082046976, tape.AvailableRawCapacity)
    assertNonNilStringPtr(t, "BarCode", "101000L6", tape.BarCode)
    assertStringPtrIsNil(t, "BucketId", tape.BucketId)
    assertStringPtrIsNil(t, "DescriptionForIdentification", tape.DescriptionForIdentification)
    assertStringPtrIsNil(t, "EjectDate", tape.EjectDate)
    assertStringPtrIsNil(t, "EjectLabel", tape.EjectLabel)
    assertStringPtrIsNil(t, "EjectLocation", tape.EjectLocation)
    assertStringPtrIsNil(t, "EjectPending", tape.EjectPending)
    assertBool(t, "FullOfData", false, tape.FullOfData)
    assertString(t, "Id", "c7c431df-f95d-4533-b350-ffd7a8a5caac", tape.Id)
    assertNonNilStringPtr(t, "LastAccessed", "2015-09-04T06:53:08.236", tape.LastAccessed)
    assertNonNilStringPtr(t, "LastCheckpoint", "eb77ea67-3c83-47ec-8714-cd46a97dc392:2", tape.LastCheckpoint)
    assertNonNilStringPtr(t, "LastModified", "2015-08-21T16:14:30.714", tape.LastModified)
    assertStringPtrIsNil(t, "LastVerified", tape.LastVerified)
    assertNonNilStringPtr(t, "PartitionId", "4f8a5cbb-9837-41d9-afd1-cebed41f18f7", tape.PartitionId)
    if *tape.PreviousState != models.UNDEFINED {
        t.Fatalf("Expected previous state '%d' but was '%d'.", models.UNDEFINED, *tape.PreviousState)
    }
    assertNonNilStringPtr(t, "SerialNumber", "HP-W130501213", tape.SerialNumber)
    assertString(t, "State", models.TAPE_STATE_NORMAL.String(), tape.State.String())
    assertNonNilIntPtr(t, "TotalRawCapacity", 2408088338432, tape.TotalRawCapacity)
    assertString(t, "Type", models.TAPE_TYPE_LTO6.String(), tape.Type.String())
    assertBool(t, "WriteProtected", false, tape.WriteProtected)
}
