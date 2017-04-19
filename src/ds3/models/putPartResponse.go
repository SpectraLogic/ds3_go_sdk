package models

import (
    "strings"
    "ds3/networking"
)

type PutPartResponse struct {
    ETag string
}

func NewPutPartResponse(webResponse networking.WebResponse) (*PutPartResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        etags := (*webResponse.Header())["etag"]
        var etag string
        if len(etags) > 0 {
            etag = strings.Trim(etags[0], "\"")
        }
        return &PutPartResponse{etag}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
