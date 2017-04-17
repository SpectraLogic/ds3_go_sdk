package models

import (
    "ds3/networking"
)

type DeleteBucketResponse struct {}

func NewDeleteBucketResponse(webResponse networking.WebResponse) (*DeleteBucketResponse, error) {
    expectedStatusCodes := []int { 204 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteBucketResponse{}, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
