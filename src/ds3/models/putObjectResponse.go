package models

import (
    "net/http"
    "ds3/networking"
)

type PutObjectResponse struct {}

func NewPutObjectResponse(webResponse networking.WebResponse) (*PutObjectResponse, error) {
    if err := checkStatusCode(webResponse, http.StatusOK); err != nil {
        return nil, err
    } else {
        return &PutObjectResponse{}, nil
    }
}

