package models

import (
    "net/http"
    "ds3/networking"
)

type DeleteBucketResponse struct {}

func NewDeleteBucketResponse(ds3Response networking.Ds3Response) (*DeleteBucketResponse, error) {
    if err := checkStatusCode(ds3Response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &DeleteBucketResponse{}, nil
    }
}

