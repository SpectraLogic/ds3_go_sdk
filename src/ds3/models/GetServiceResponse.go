package models

import (
    "io"
    "fmt"
    "errors"
    "net/http"
)

type GetServiceResponse struct {
    Owner Owner
    Buckets []Bucket `xml:"Buckets>Bucket"`
}

type Owner struct {
    Id string `xml:"ID"`
    DisplayName string
}

type Bucket struct {
    Name string
    CreationDate string
}

func NewGetServiceResponse(response *http.Response) (*GetServiceResponse, error) {
    expectedStatusCode := http.StatusOK
    if response.StatusCode == expectedStatusCode {
        return readGetServiceResponseBody(response.Body)
    } else {
        return nil, errors.New(fmt.Sprintf(
            "Expected HTTP status code %d but got %d.",
            expectedStatusCode,
            response.StatusCode,
        ))
    }
}

func readGetServiceResponseBody(stream io.ReadCloser) (*GetServiceResponse, error) {
    var body GetServiceResponse
    err := readResponseBody(stream, &body)
    if err != nil {
        return nil, err
    }
    return &body, nil
}

