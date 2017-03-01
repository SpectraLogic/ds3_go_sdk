package models

import (
    "net/http"
    "ds3/networking"
)

type PutBucketResponse struct {}

func NewPutBucketResponse(webResponse networking.WebResponse) (*PutBucketResponse, error) {
    if err := checkStatusCode(webResponse, http.StatusOK); err != nil {
        return nil, err
    } else {
        return &PutBucketResponse{}, nil
    }
}

