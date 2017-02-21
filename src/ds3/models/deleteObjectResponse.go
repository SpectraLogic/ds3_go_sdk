package models

import (
    "net/http"
    "ds3/networking"
)

type DeleteObjectResponse struct {}

func NewDeleteObjectResponse(ds3Response networking.Ds3Response) (*DeleteObjectResponse, error) {
    if err := checkStatusCode(ds3Response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &DeleteObjectResponse{}, nil
    }
}


