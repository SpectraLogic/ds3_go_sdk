package models

import (
    "ds3/networking"
)

type DeleteBucketResponse struct {}

func NewDeleteBucketResponse(webResponse networking.WebResponse) (*DeleteBucketResponse, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteBucketResponse{}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
