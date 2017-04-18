package models

import (
    "ds3/networking"
)

type AbortMultipartResponse struct {}

func NewAbortMultipartResponse(webResponse networking.WebResponse) (*AbortMultipartResponse, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &AbortMultipartResponse{}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

