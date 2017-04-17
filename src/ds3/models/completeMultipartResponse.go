package models

import (
    "strings"
    "ds3/networking"
)

type CompleteMultipartResponse struct {
    Location string
    Bucket string
    Key string
    ETag string
}

func NewCompleteMultipartResponse(webResponse networking.WebResponse) (*CompleteMultipartResponse, error) {
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CompleteMultipartResponse
        if err := readResponseBody(webResponse, &body); err != nil {
            return nil, err
        }
        body.ETag = strings.Trim(body.ETag, "\"")
        return &body, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}


