package models

import (
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

func NewListPartsResponse(webResponse networking.WebResponse) (*ListPartsResponse, error) {
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ListPartsResponse
        if err := readResponseBody(webResponse, &body); err != nil {
            return nil, err
        }
        return &body, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

