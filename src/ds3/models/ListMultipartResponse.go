package models

import (
    "net/http"
    "ds3/networking"
)

type ListMultipartResponse struct {
    Bucket string
    KeyMarker string
    UploadIdMarker string
    NextKeyMarker string
    NextUploadIdMarker string
    MaxUploads int
    IsTruncated bool
    Uploads []Upload `xml:"Upload"`
}

type Upload struct {
    Key string
    UploadId string
    Initiator Owner
    Owner Owner
    StorageClass string
    Initiated string
}

func NewListMultipartResponse(response networking.Response) (*ListMultipartResponse, error) {
    var body ListMultipartResponse
    if err := readResponseBody(response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}


