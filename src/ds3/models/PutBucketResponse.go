package models

import (
    "net/http"
    "ds3/networking"
)

type PutBucketResponse struct {}

func NewPutBucketResponse(response networking.Response) (*PutBucketResponse, error) {
    if err := checkStatusCode(response, http.StatusOK); err != nil {
        return nil, err
    } else {
        return &PutBucketResponse{}, nil
    }
}

