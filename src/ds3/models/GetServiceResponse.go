package models

import (
    "errors"
    "io"
    "io/ioutil"
    "net/http"
    "encoding/xml"
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
    if response.StatusCode == int(OK) {
        return readResponseBody(response.Body)
    } else {
        return nil, errors.New("Response status error") //TODO: fix error handling
    }
}

//TODO: improve error handling
func readResponseBody(stream io.ReadCloser) (*GetServiceResponse, error) {
    defer stream.Close()

    // Get the bytes or forward the error.
    bytes, err := ioutil.ReadAll(stream)
    if err != nil {
        return nil, err
    }

    // Deserialize or forward the error.
    var body GetServiceResponse
    if err = xml.Unmarshal(bytes, &body); err != nil {
        return nil, err
    }

    // Return the result.
    return &body, nil
}

