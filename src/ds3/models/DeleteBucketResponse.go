package models

import (
    "net/http"
    "ds3/networking"
)

type DeleteBucketResponse struct {}

func NewDeleteBucketResponse(response networking.Response) (*DeleteBucketResponse, error) {
    if err := checkStatusCode(response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &DeleteBucketResponse{}, nil
    }
}

