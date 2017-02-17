package models

import (
    "strings"
    "net/http"
    "ds3/network"
)

type CompleteMultipartResponse struct {
    Location string
    Bucket, Key string
    ETag string
}

func NewCompleteMultipartResponse(response net.Response) (*CompleteMultipartResponse, error) {
    var body CompleteMultipartResponse
    if err := readResponseBody(response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    body.ETag = strings.Trim(body.ETag, "\"")
    return &body, nil
}


