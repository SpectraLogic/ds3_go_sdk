package models

import (
    "io"
    "net/http"
    "ds3/networking"
)

type GetObjectResponse struct {
    Content io.ReadCloser
}

func NewGetObjectResponse(ds3Response networking.Ds3Response) (*GetObjectResponse, error) {
    if err := checkStatusCode(ds3Response, http.StatusOK); err != nil {
        return nil, err
    }
    return &GetObjectResponse{ds3Response.Body()}, nil
}

