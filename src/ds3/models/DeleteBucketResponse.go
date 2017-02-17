package models

import (
    "net/http"
    "ds3/network"
)

type DeleteBucketResponse struct {}

func NewDeleteBucketResponse(response net.Response) (*DeleteBucketResponse, error) {
    if err := checkStatusCode(response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &DeleteBucketResponse{}, nil
    }
}

