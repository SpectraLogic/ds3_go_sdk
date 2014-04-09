package models

import (
    "net/http"
    "ds3/net"
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
    return &body, nil
}


