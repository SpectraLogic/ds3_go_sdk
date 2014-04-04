package models

import (
    "net/http"
)

type PutBucketResponse struct {}

func NewPutBucketResponse(response *http.Response) (*PutBucketResponse, error) {
    if err := checkStatusCode(response, http.StatusOK); err != nil {
        return nil, err
    } else {
        return &PutBucketResponse{}, nil
    }
}


