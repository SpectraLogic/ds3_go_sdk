package models

import (
    "net/http"
    "ds3/networking"
)

type DeleteObjectResponse struct {}

func NewDeleteObjectResponse(response networking.Response) (*DeleteObjectResponse, error) {
    if err := checkStatusCode(response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &DeleteObjectResponse{}, nil
    }
}


