package models

import (
    "ds3/networking"
)

type PutObjectResponse struct {}

func NewPutObjectResponse(webResponse networking.WebResponse) (*PutObjectResponse, error) {
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &PutObjectResponse{}, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

