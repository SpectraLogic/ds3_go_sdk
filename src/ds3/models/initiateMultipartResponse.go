package models

import (
    "ds3/networking"
)

type InitiateMultipartResponse struct {
    Bucket string
    Key string
    UploadId string
}

func NewInitiateMultipartResponse(webResponse networking.WebResponse) (*InitiateMultipartResponse, error) {
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body InitiateMultipartResponse
        if err := readResponseBody(webResponse, &body); err != nil {
            return nil, err
        }
        return &body, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

