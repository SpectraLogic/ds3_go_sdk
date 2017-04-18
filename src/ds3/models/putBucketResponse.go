package models

import (
    "ds3/networking"
)

type PutBucketResponse struct {}

func NewPutBucketResponse(webResponse networking.WebResponse) (*PutBucketResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &PutBucketResponse{}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

