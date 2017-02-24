package models

import (
    "net/http"
    "ds3/networking"
)

type PutObjectResponse struct {}

func NewPutObjectResponse(ds3Response networking.Ds3Response) (*PutObjectResponse, error) {
    if err := checkStatusCode(ds3Response, http.StatusOK); err != nil {
        return nil, err
    } else {
        return &PutObjectResponse{}, nil
    }
}

