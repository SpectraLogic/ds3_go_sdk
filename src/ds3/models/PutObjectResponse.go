package models

import (
    "net/http"
    "ds3/network"
)

type PutObjectResponse struct {}

func NewPutObjectResponse(response net.Response) (*PutObjectResponse, error) {
    if err := checkStatusCode(response, http.StatusOK); err != nil {
        return nil, err
    } else {
        return &PutObjectResponse{}, nil
    }
}

