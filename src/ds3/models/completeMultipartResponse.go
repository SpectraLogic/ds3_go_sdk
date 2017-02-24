package models

import (
    "strings"
    "net/http"
    "ds3/networking"
)

type CompleteMultipartResponse struct {
    Location string
    Bucket string
    Key string
    ETag string
}

func NewCompleteMultipartResponse(ds3Response networking.Ds3Response) (*CompleteMultipartResponse, error) {
    var body CompleteMultipartResponse
    if err := readResponseBody(ds3Response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    body.ETag = strings.Trim(body.ETag, "\"")
    return &body, nil
}


