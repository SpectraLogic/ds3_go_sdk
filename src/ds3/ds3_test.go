package ds3

import (
    "testing"
    "net/url"
    "ds3/net"
    "ds3/models"
)

func TestGetService(t *testing.T) {
    stringResponse := "<ListAllMyBucketsResult xmlns=\"http://doc.s3.amazonaws.com/2006-03-01\"><Owner><ID>ryan</ID><DisplayName>ryan</DisplayName></Owner><Buckets><Bucket><Name>testBucket2</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest1</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest2</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest3</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest4</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest5</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>bulkTest6</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>testBucket3</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>testBucket1</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket><Bucket><Name>testbucket</Name><CreationDate>2013-12-11T23:20:09</CreationDate></Bucket></Buckets></ListAllMyBucketsResult>";
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
        Expecting(net.GET, "/", &url.Values{}).
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
        Expecting(net.GET, "/", &url.Values{}).
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

