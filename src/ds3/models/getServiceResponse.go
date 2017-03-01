package models

import (
    "net/http"
    "ds3/networking"
)

type GetServiceResponse struct {
    Owner Owner
    Buckets []Bucket `xml:"Buckets>Bucket"`
}

type Bucket struct {
    Name string
    CreationDate string
}

func NewGetServiceResponse(webResponse networking.WebResponse) (*GetServiceResponse, error) {
    var body GetServiceResponse
    if err := readResponseBody(webResponse, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

