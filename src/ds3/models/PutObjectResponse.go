package models

import (
    "net/http"
    "ds3/networking"
)

type PutObjectResponse struct {}

func NewPutObjectResponse(response networking.Response) (*PutObjectResponse, error) {
    if err := checkStatusCode(response, http.StatusOK); err != nil {
        return nil, err
    } else {
        return &PutObjectResponse{}, nil
    }
}

