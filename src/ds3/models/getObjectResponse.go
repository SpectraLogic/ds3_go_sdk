package models

import (
    "io"
    "net/http"
    "ds3/networking"
)

type GetObjectResponse struct {
    Content io.ReadCloser
}

func NewGetObjectResponse(webResponse networking.WebResponse) (*GetObjectResponse, error) {
    if err := checkStatusCode(webResponse, http.StatusOK); err != nil {
        return nil, err
    }
    return &GetObjectResponse{webResponse.Body()}, nil
}

