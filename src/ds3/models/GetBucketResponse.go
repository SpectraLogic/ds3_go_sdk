package models

import (
    "net/http"
    "ds3/network"
)

type GetBucketResponse struct {
    Name string
    Contents []Object
    CreationDate string
    Delimiter string
    IsTruncated bool
    Marker string
    MaxKeys int
    NextMarker string
    Prefix string
}

func NewGetBucketResponse(response net.Response) (*GetBucketResponse, error) {
    var body GetBucketResponse
    if err := readResponseBody(response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

