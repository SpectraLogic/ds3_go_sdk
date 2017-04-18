package models

import (
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

func NewListMultipartResponse(webResponse networking.WebResponse) (*ListMultipartResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ListMultipartResponse
        if err := readResponseBody(webResponse, &body); err != nil {
            return nil, err
        }
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
