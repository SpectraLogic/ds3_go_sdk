package models

import (
    "net/http"
    "ds3/networking"
)

type ListPartsResponse struct {
    Bucket string
    Key string
    UploadId string
    Initiator Owner
    Owner Owner
    StorageClass string
    PartNumberMarker int
    NextPartNumberMarker int
    MaxParts int
    IsTruncated bool
    Parts []UploadedPart `xml:"Part"`
}

type UploadedPart struct {
    PartNumber int
    LastModified string
    ETag string
    Size int64
}

func NewListPartsResponse(response networking.Response) (*ListPartsResponse, error) {
    var body ListPartsResponse
    if err := readResponseBody(response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

