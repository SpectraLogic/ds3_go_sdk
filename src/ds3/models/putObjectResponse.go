package models

import (
    "ds3/networking"
)

type PutObjectResponse struct {}

func NewPutObjectResponse(webResponse networking.WebResponse) (*PutObjectResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &PutObjectResponse{}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

