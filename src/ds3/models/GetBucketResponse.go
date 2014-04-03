package models

import (
    "net/http"
)

type GetBucketResponse struct {
    Contents []Object
    CreationDate string
    Delimiter string
    IsTruncated string
    Marker string
    MaxKeys int
    NextMarker string
    Prefix string
}

type Object struct {
    ETag string
    Key string
    LastModified string
    Owner Owner
    Size int
    StorageClass string
}

func NewGetBucketResponse(response *http.Response) (*GetBucketResponse, error) {
    var body GetBucketResponse
    if err := readResponseBody(response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

