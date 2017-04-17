package models

import (
    "ds3/networking"
)

type DeleteObjectResponse struct {}

func NewDeleteObjectResponse(webResponse networking.WebResponse) (*DeleteObjectResponse, error) {
    expectedStatusCodes := []int { 204 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteObjectResponse{}, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
