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

func NewCompleteMultipartResponse(webResponse networking.WebResponse) (*CompleteMultipartResponse, error) {
    var body CompleteMultipartResponse
    if err := readResponseBody(webResponse, http.StatusOK, &body); err != nil {
        return nil, err
    }
    body.ETag = strings.Trim(body.ETag, "\"")
    return &body, nil
}


