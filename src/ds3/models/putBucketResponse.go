package models

import (
    "ds3/networking"
)

type PutBucketResponse struct {}

func NewPutBucketResponse(webResponse networking.WebResponse) (*PutBucketResponse, error) {
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &PutBucketResponse{}, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

