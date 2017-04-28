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
        Expecting(networking.GET, "/", &url.Values{}, nil).
        Returning(200, stringResponse, nil).
        GetService(models.NewGetServiceRequest())

    // Check the response.
    if response == nil {
        t.Error("Received an unexpected nil response.")
    }
    if response.Owner.Id != "ryan" {
        t.Errorf("Expected owner id of 'ryan' but got '%s'.", response.Owner.Id)
    }
    if response.Owner.DisplayName != "ryan" {
        t.Errorf("Expected owner display name of 'ryan' but got '%s'.", response.Owner.DisplayName)
    }
    if len(response.Buckets) != len(expectedBucketNames) {
        t.Errorf("Parsed an unexpected number (%d) of buckets.", len(response.Buckets))
    }
    for i, bucket := range response.Buckets {
        if bucket.Name != expectedBucketNames[i] {
            t.Errorf("%dth bucket name is incorrect (%s).", i + 1, bucket.Name)
        }
        if bucket.CreationDate != expectedCreationDates[i] {
            t.Errorf("%dth bucket creation date is incorrect (%s).", i + 1, bucket.CreationDate)
        }
    }

    // Validate the error contents.
    if err != nil {
        t.Errorf("Received an unexpected error: %s", err.Error())
    }
}

func TestGetBadService(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.GET, "/", &url.Values{}, nil).
        Returning(400, "", nil).
        GetService(models.NewGetServiceRequest())

    // Check the response.
    if response != nil {
        t.Error("Expected a nil response but received a value.")
    }

    // Validate the error contents.
    if err == nil {
        t.Error("Expected an error but didn't get one.")
    } else {
        expectedError := "Received a status code of 400 when [200] was expected. Could not parse the response for additional information."
        actualError := err.Error()
        if actualError != expectedError {
            t.Errorf("Expected a different error message than received: '%s'", actualError)
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
        Expecting(networking.GET, "/remoteTest16", &url.Values{}, nil).
        Returning(200, stringResponse, nil).
        GetBucket(models.NewGetBucketRequest("remoteTest16"))

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    } else {
        if len(response.Contents) != len(keys) {
            t.Errorf("Expected %d objects but got %d.", len(keys), len(response.Contents))
        } else {
            if response.Name != "remoteTest16" {
                t.Errorf("Expected bucket name 'remoteTest16' but got '%s'.", response.Name)
            }
            if response.Prefix != "" {
                t.Errorf("Expected empty prefix but got '%s'.", response.Prefix)
            }
            if response.Marker != "" {
                t.Errorf("Expected empty marker but got '%s'.", response.Marker)
            }
            if response.MaxKeys != 1000 {
                t.Errorf("Expected max keys of 1000 but got %d.", response.MaxKeys)
            }
            if response.IsTruncated != false {
                t.Error("Expected that the result would not be truncated, but it was.")
            }
            for i, object := range response.Contents {
                if object.Key != keys[i] {
                    t.Errorf("Expected key '%s' but got '%s'.", keys[i], object.Key)
                }
                if object.LastModified != lastModifieds[i] {
                    t.Errorf("Expected last modified '%s' but got '%s'.", lastModifieds[i], object.LastModified)
                }
                if object.ETag != etags[i] {
                    t.Errorf("Expected ETag '%s' but got '%s'.", etags[i], object.ETag)
                }
                if object.Size != sizes[i] {
                    t.Errorf("Expected size %d but got %d.", sizes[i], object.Size)
                }
                if object.StorageClass != storageClasses[i] {
                    t.Errorf("Expected storage class '%s' but got '%s'.", storageClasses[i], object.StorageClass)
                }
                if object.Owner.Id != ids[i] {
                    t.Errorf("Expected owner id '%s' but got '%s'.", ids[i], object.Owner.Id)
                }
                if object.Owner.DisplayName != displayNames[i] {
                    t.Errorf("Expected owner display name '%s' but got '%s'.", displayNames[i], object.Owner.DisplayName)
                }
            }
        }
    }
}

func TestPutBucket(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.PUT, "/bucketName", &url.Values{}, nil).
        Returning(200, "", nil).
        PutBucket(models.NewPutBucketRequest("bucketName"))

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    }
}

func TestDeleteBucket(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.DELETE, "/bucketName", &url.Values{}, nil).
        Returning(204, "", nil).
        DeleteBucket(models.NewDeleteBucketRequest("bucketName"))

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    }
}

func TestDeleteObject(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.DELETE, "/bucketName/my/file.txt", &url.Values{}, nil).
        Returning(204, "", nil).
        DeleteObject(models.NewDeleteObjectRequest("bucketName", "my/file.txt"))

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    }
}

func TestGetBadBucket(t *testing.T) {
    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.GET, "/remoteTest16", &url.Values{}, nil).
        Returning(400, "", nil).
        GetBucket(models.NewGetBucketRequest("remoteTest16"))

    // Check the error result.
    if err == nil {
        t.Error("Expected an error but got nil.")
    } else {
        expectedError := "Received a status code of 400 when [200] was expected. Could not parse the response for additional information."
        actualError := err.Error()
        if actualError != expectedError {
            t.Errorf("Expected a different error message than received: '%s'", actualError)
        }
    }

    // Check the response value.
    if response != nil {
        t.Error("Response was supposed to be nil.")
    }
}

func TestGetObject(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.GET, "/bucketName/object", &url.Values{}, nil).
        Returning(200, stringResponse, nil).
        GetObject(models.NewGetObjectRequest("bucketName", "object"))

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    } else if response.Content == nil {
        t.Error("Response content was unexpectedly nil.")
    } else {
        defer response.Content.Close()
        bs, readErr := ioutil.ReadAll(response.Content)
        if readErr != nil {
            t.Errorf("Unexpected error '%s'.", readErr.Error())
        } else if string(bs) != stringResponse {
            t.Errorf("Expected '%s' but got '%s'.", stringResponse, string(bs))
        }
    }
}

func TestGetObjectRange(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    request := models.NewGetObjectRequest("bucketName", "object")
    request.WithRange(20, 179)
    response, err := mockedClient(t).
        Expecting(networking.GET, "/bucketName/object", &url.Values{}, nil).
        Returning(200, stringResponse, &http.Header{"Range": []string{"bytes=20-179"}}).
        GetObject(request)

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    } else if response.Content == nil {
        t.Error("Response content was unexpectedly nil.")
    } else {
        defer response.Content.Close()
        bs, readErr := ioutil.ReadAll(response.Content)
        if readErr != nil {
            t.Errorf("Unexpected error '%s'.", readErr.Error())
        } else if string(bs) != stringResponse {
            t.Errorf("Expected '%s' but got '%s'.", stringResponse, string(bs))
        }
    }
}

func TestGetJobToReplicateSpectraS3(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    request := models.NewGetJobToReplicateSpectraS3Request("23a876ec-2fac-4dc8-b8e6-98d6026e7f4a")

    expectedParams := &url.Values{"replicate": []string{""}}
    response, err := mockedClient(t).
        Expecting(networking.GET, "/_rest_/job/23a876ec-2fac-4dc8-b8e6-98d6026e7f4a", expectedParams, nil).
        Returning(200, stringResponse, nil).
        GetJobToReplicateSpectraS3(request)

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    } else {
        if response.Content != stringResponse {
            t.Errorf("Unexpected response: expected '%s' but was '%s'", stringResponse, response.Content)
        }
    }
}

func TestPutObject(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(networking.PUT, "/bucketName/object", &url.Values{}, nil).
        Returning(200, "", nil).
        PutObject(models.NewPutObjectRequest(
            "bucketName",
            "object",
        networking.BuildByteReaderWithSizeDecorator([]byte(stringResponse)),
        ))

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    }
}

func TestBulkPut(t *testing.T) {
    runBulkTest(
        t,
        "start_bulk_put",
        func(client *Client, objects []models.Ds3Object) ([][]models.Object, error) {
            request, err := client.BulkPut(models.NewBulkPutRequest("bucketName", objects))
            return request.Objects, err
        },
    )
}

func TestBulkGet(t *testing.T) {
    runBulkTest(
        t,
        "start_bulk_get",
        func(client *Client, objects []models.Ds3Object) ([][]models.Object, error) {
            request, err := client.BulkGet(models.NewBulkGetRequest("bucketName", objects))
            return request.Objects, err
        },
    )
}

type bulkTest func(*Client, []models.Ds3Object) ([][]models.Object, error)

func runBulkTest(t *testing.T, operation string, callToTest bulkTest) {
    keys := []string { "file2", "file1", "file3" }
    sizes := []int64 { 1202, 256, 2523 }

    stringRequest := "<objects><object name=\"file1\" size=\"256\"></object><object name=\"file2\" size=\"1202\"></object><object name=\"file3\" size=\"2523\"></object></objects>"
    stringResponse := "<masterobjectlist><objects><object name='file2' size='1202'/><object name='file1' size='256'/><object name='file3' size='2523'/></objects></masterobjectlist>"

    inputObjects := []models.Ds3Object {
        {Name: "file1", Size: 256 },
        {Name: "file2", Size: 1202 },
        {Name: "file3", Size: 2523 },
    }

    // Create and run the mocked client.
    client := mockedClient(t).
        Expecting(
        networking.PUT,
            "/_rest_/buckets/bucketName",
            &url.Values{"operation": []string{operation}},
            &stringRequest,
        ).
        Returning(200, stringResponse, nil)
    response, err := callToTest(client, inputObjects)

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    }
    if len(response) != 1 {
        t.Errorf("Expected 1 object list but got %d.", len(response))
    }
    if len(response[0]) != len(keys) {
        t.Errorf("Expected %d objects but got %d.", len(keys), len(response[0]))
    }
    for i, obj := range response[0] {
        if obj.Key != keys[i] {
            t.Errorf("Expected key %s but got %s.", keys[i], obj.Key)
        }
        if obj.Size != sizes[i] {
            t.Errorf("Expected size %d but got %d.", sizes[i], obj.Size)
        }
    }
}

func TestInitiateMultipart(t *testing.T) {
    stringResponse := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><InitiateMultipartUploadResult xmlns=\"http://s3.amazonaws.com/doc/2006-03-01/\"><Bucket>example-bucket</Bucket><Key>example-object</Key><UploadId>VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA</UploadId></InitiateMultipartUploadResult>"

    // Create and run the mocked client.
    qs := &url.Values{"uploads": []string{""}}
    response, err := mockedClient(t).
        Expecting(networking.POST, "/bucketName/object", qs, nil).
        Returning(200, stringResponse, nil).
        InitiateMultipart(models.NewInitiateMultipartRequest(
            "bucketName",
            "object",
        ))

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    }
    if response.Bucket != "example-bucket" {
        t.Errorf("Expected bucket 'example-bucket' but got '%s'.", response.Bucket)
    }
    if response.Key != "example-object" {
        t.Errorf("Expected key 'example-object' but got '%s'.", response.Key)
    }
    if response.UploadId != "VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA" {
        t.Errorf("Expected upload id 'VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA' but got '%s'.", response.UploadId)
    }
}

func TestPutPart(t *testing.T) {
    content := "this is the part content"
    partNumber := 2
    uploadId := "VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA"
    etag := "b54357faf0632cce46e942fa68356b38"

    // Create and run the mocked client.
    qs := &url.Values{
        "partNumber": []string{strconv.Itoa(partNumber)},
        "uploadId": []string{uploadId},
    }
    response, err := mockedClient(t).
        Expecting(networking.PUT, "/bucketName/object", qs, nil).
        Returning(200, "", &http.Header{"etag": []string{"\"" + etag + "\""}}).
        PutPart(models.NewPutPartRequest(
            "bucketName",
            "object",
            partNumber,
            uploadId,
            networking.BuildByteReaderWithSizeDecorator([]byte(content)),
        ))

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    }
    if response.ETag != etag {
        t.Errorf("Expected etag '%s' but got '%s'.", etag, response.ETag)
    }
}

func TestCompleteMultipart(t *testing.T) {
    bucket := "bucketName"
    key := "object"
    location := "http://my-server/bucketName/object"
    etag := "b54357faf0632cce46e942fa68356b38"
    uploadId := "VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA"
    expectedRequest := "<CompleteMultipartUpload><Part><PartNumber>1</PartNumber><ETag>7a112844c1a2327e617f530cb06dccf8</ETag></Part><Part><PartNumber>2</PartNumber><ETag>7162e29f4e40da7f521d0794b57770ba</ETag></Part></CompleteMultipartUpload>"
    expectedResponse := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><CompleteMultipartUploadResult xmlns=\"http://s3.amazonaws.com/doc/2006-03-01/\"><Location>http://my-server/bucketName/object</Location><Bucket>bucketName</Bucket><Key>object</Key><ETag>\"b54357faf0632cce46e942fa68356b38\"</ETag></CompleteMultipartUploadResult>"

    // Create and run the mocked client.
    qs := &url.Values{
        "uploadId": []string{uploadId},
    }
    response, err := mockedClient(t).
        Expecting(networking.POST, "/bucketName/object", qs, &expectedRequest).
        Returning(200, expectedResponse, &http.Header{"etag": []string{etag}}).
        CompleteMultipart(models.NewCompleteMultipartRequest(
            bucket,
            key,
            uploadId,
            []models.Part{
                {PartNumber: 1, ETag: "7a112844c1a2327e617f530cb06dccf8"},
                {PartNumber: 2, ETag: "7162e29f4e40da7f521d0794b57770ba"},
            },
        ))

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    }
    if response.Location != location {
        t.Errorf("Expected location '%s' but got '%s'.", location, response.Location)
    }
    if response.Bucket != bucket {
        t.Errorf("Expected bucket '%s' but got '%s'.", bucket, response.Bucket)
    }
    if response.Key != key {
        t.Errorf("Expected key '%s' but got '%s'.", key, response.Key)
    }
    if response.ETag != etag {
        t.Errorf("Expected etag '%s' but got '%s'.", etag, response.ETag)
    }
}

func TestDeleteObjects(t *testing.T) {
    bucket := "bucketName"
    objectNames := []string {"obj1", "obj2", "obj3"}
    expectedRequest := "<Delete><Object><Key>obj1</Key></Object><Object><Key>obj2</Key></Object><Object><Key>obj3</Key></Object></Delete>"
    expectedResponse := "<DeleteResult><Deleted><Key>obj1</Key></Deleted><Deleted><Key>obj2</Key></Deleted><Error><Code>ObjectNotFound</Code><Key>obj3</Key><Message>Object not found</Message></Error></DeleteResult>"

    expectedDeleted := []models.DeletedObject{{"obj1"}, {"obj2"}}

    expectedErrors := []models.DeleteError {
        {
            Code:    "ObjectNotFound",
            Key:     "obj3",
            Message: "Object not found",
        },
    }

    // Create and run the mocked client.
    qs := &url.Values{
        "delete": []string{""},
    }
    response, err := mockedClient(t).
        Expecting(networking.POST, "/bucketName", qs, &expectedRequest).
        Returning(200, expectedResponse, &http.Header{}).
        DeleteObjects(models.NewDeleteObjectsRequest(bucket, objectNames))

    // Check the error result.
    if err != nil {
        t.Errorf("Unexpected error '%s'.", err.Error())
    }

    // Check the response value.
    if response == nil {
        t.Error("Response was unexpectedly nil.")
    }

    if !reflect.DeepEqual(response.Deleted, expectedDeleted) {
        t.Errorf("Expected '%s' but got '%s'", expectedDeleted, response.Deleted)
    }

    if !reflect.DeepEqual(response.Errors, expectedErrors) {
        t.Errorf("Expected '%s' but got '%s'", expectedErrors, response.Errors)
    }
}
