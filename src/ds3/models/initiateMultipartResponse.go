package models

import (
    "net/http"
    "ds3/networking"
)

type InitiateMultipartResponse struct {
    Bucket string
    Key string
    UploadId string
}

func NewInitiateMultipartResponse(webResponse networking.WebResponse) (*InitiateMultipartResponse, error) {
    var body InitiateMultipartResponse
    if err := readResponseBody(webResponse, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

