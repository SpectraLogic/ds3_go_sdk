package models

import (
    "net/http"
    "ds3/networking"
)

type DeleteObjectResponse struct {}

func NewDeleteObjectResponse(webResponse networking.WebResponse) (*DeleteObjectResponse, error) {
    if err := checkStatusCode(webResponse, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &DeleteObjectResponse{}, nil
    }
}


