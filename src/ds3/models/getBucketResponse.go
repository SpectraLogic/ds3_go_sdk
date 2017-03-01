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

func NewGetBucketResponse(webResponse networking.WebResponse) (*GetBucketResponse, error) {
    var body GetBucketResponse
    if err := readResponseBody(webResponse, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

