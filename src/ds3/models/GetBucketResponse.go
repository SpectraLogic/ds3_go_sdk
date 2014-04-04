package models

import (
    "net/http"
)

type GetBucketResponse struct {
    Contents []Object
    CreationDate string
    Delimiter string
    IsTruncated bool
    Marker string
    MaxKeys int
    NextMarker string
    Prefix string
}

func NewGetBucketResponse(response *http.Response) (*GetBucketResponse, error) {
    var body GetBucketResponse
    if err := readResponseBody(response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

