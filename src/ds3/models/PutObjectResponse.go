package models

import (
    "net/http"
)

type PutObjectResponse struct {}

func NewPutObjectResponse(response *http.Response) (*PutObjectResponse, error) {
    if err := checkStatusCode(response, http.StatusOK); err != nil {
        return nil, err
    } else {
        return &PutObjectResponse{}, nil
    }
}



