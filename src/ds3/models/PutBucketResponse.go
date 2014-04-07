package models

import (
    "net/http"
    "ds3/net"
)

type PutBucketResponse struct {}

func NewPutBucketResponse(response net.Response) (*PutBucketResponse, error) {
    if err := checkStatusCode(response, http.StatusOK); err != nil {
        return nil, err
    } else {
        return &PutBucketResponse{}, nil
    }
}

