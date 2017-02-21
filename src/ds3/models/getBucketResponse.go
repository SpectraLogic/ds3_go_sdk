package models

import (
    "net/http"
    "ds3/networking"
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

func NewGetBucketResponse(ds3Response networking.Ds3Response) (*GetBucketResponse, error) {
    var body GetBucketResponse
    if err := readResponseBody(ds3Response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

