package models

import (
    "errors"
    "io"
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
    if response.StatusCode == http.StatusOK {
        return readGetServiceResponseBody(response.Body)
    } else {
        return nil, errors.New("Response status error") //TODO: fix error handling
    }
}

//TODO: improve error handling
func readGetServiceResponseBody(stream io.ReadCloser) (*GetServiceResponse, error) {
    var body GetServiceResponse
    err := readResponseBody(stream, &body)
    if err != nil {
        return nil, err
    }
    return &body, nil
}

