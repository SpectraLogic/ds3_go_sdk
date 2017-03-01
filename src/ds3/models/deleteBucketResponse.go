package models

import (
    "net/http"
    "ds3/networking"
)

type DeleteBucketResponse struct {}

func NewDeleteBucketResponse(webResponse networking.WebResponse) (*DeleteBucketResponse, error) {
    if err := checkStatusCode(webResponse, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &DeleteBucketResponse{}, nil
    }
}

