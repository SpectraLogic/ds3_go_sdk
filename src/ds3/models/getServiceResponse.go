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

func NewGetServiceResponse(ds3Response networking.Ds3Response) (*GetServiceResponse, error) {
    var body GetServiceResponse
    if err := readResponseBody(ds3Response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

