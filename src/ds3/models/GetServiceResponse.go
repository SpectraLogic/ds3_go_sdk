package models

import (
    "net/http"
    "ds3/net"
)

type GetServiceResponse struct {
    Owner Owner
    Buckets []Bucket `xml:"Buckets>Bucket"`
}

type Bucket struct {
    Name string
    CreationDate string
}

func NewGetServiceResponse(response net.Response) (*GetServiceResponse, error) {
    var body GetServiceResponse
    if err := readResponseBody(response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

