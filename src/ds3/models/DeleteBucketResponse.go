package models

import "net/http"

type DeleteBucketResponse struct {}

func NewDeleteBucketResponse(response *http.Response) (*DeleteBucketResponse, error) {
    if err := checkStatusCode(response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &DeleteBucketResponse{}, nil
    }
}

