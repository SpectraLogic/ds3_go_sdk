package ds3

import (
    "testing"
    "net/url"
    "io/ioutil"
    "ds3/net"
    "ds3/models"
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
        Expecting(net.GET, "/", &url.Values{}, nil).
        Returning(200, stringResponse).
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
        Expecting(net.GET, "/", &url.Values{}, nil).
        Returning(400, "").
        GetService(models.NewGetServiceRequest())

    // Check the response.
    if response != nil {
        t.Error("Expected a nil response but received a value.")
    }

    // Validate the error contents.
    if err == nil {
        t.Error("Expected an error but didn't get one.")
    } else {
        expectedError := "Received a status code of 400 when 200 was expected. Could not parse the response for additional information."
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
        Expecting(net.GET, "/remoteTest16", &url.Values{}, nil).
        Returning(200, stringResponse).
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
        Expecting(net.PUT, "/bucketName", &url.Values{}, nil).
        Returning(200, "").
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
        Expecting(net.DELETE, "/bucketName", &url.Values{}, nil).
        Returning(204, "").
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
        Expecting(net.DELETE, "/bucketName/my/file.txt", &url.Values{}, nil).
        Returning(204, "").
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
        Expecting(net.GET, "/remoteTest16", &url.Values{}, nil).
        Returning(400, "").
        GetBucket(models.NewGetBucketRequest("remoteTest16"))

    // Check the error result.
    if err == nil {
        t.Error("Expected an error but got nil.")
    } else {
        expectedError := "Received a status code of 400 when 200 was expected. Could not parse the response for additional information."
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
        Expecting(net.GET, "/bucketName/object", &url.Values{}, nil).
        Returning(200, stringResponse).
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

func TestPutObject(t *testing.T) {
    stringResponse := "object contents"

    // Create and run the mocked client.
    response, err := mockedClient(t).
        Expecting(net.PUT, "/bucketName/object", &url.Values{}, nil).
        Returning(200, "").
        PutObject(models.NewPutObjectRequest(
            "bucketName",
            "object",
            net.BuildSizedReadCloser([]byte(stringResponse)),
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

