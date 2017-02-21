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

func NewInitiateMultipartResponse(ds3Response networking.Ds3Response) (*InitiateMultipartResponse, error) {
    var body InitiateMultipartResponse
    if err := readResponseBody(ds3Response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}

