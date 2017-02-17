package models

import (
    "net/http"
    "ds3/networking"
)

type InitiateMultipartResponse struct {
    Bucket, Key, UploadId string
}

func NewInitiateMultipartResponse(response networking.Response) (*InitiateMultipartResponse, error) {
    var body InitiateMultipartResponse
    if err := readResponseBody(response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

