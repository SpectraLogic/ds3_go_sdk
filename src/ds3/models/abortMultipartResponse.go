package models

import (
    "ds3/networking"
)

type AbortMultipartResponse struct {}

func NewAbortMultipartResponse(webResponse networking.WebResponse) (*AbortMultipartResponse, error) {
    expectedStatusCodes := []int { 204 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &AbortMultipartResponse{}, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

