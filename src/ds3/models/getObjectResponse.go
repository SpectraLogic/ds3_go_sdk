package models

import (
    "io"
    "ds3/networking"
)

type GetObjectResponse struct {
    Content io.ReadCloser
}

func NewGetObjectResponse(webResponse networking.WebResponse) (*GetObjectResponse, error) {
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &GetObjectResponse{webResponse.Body()}, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

