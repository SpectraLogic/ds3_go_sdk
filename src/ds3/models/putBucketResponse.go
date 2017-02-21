package models

import (
    "net/http"
    "ds3/networking"
)

type PutBucketResponse struct {}

func NewPutBucketResponse(ds3Response networking.Ds3Response) (*PutBucketResponse, error) {
    if err := checkStatusCode(ds3Response, http.StatusOK); err != nil {
        return nil, err
    } else {
        return &PutBucketResponse{}, nil
    }
}

