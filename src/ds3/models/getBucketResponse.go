package models

import (
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
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBucketResponse
        if err := readResponseBody(webResponse, &body); err != nil {
            return nil, err
        }
        return &body, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
