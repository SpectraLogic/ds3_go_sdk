package models

import (
    "io"
    "net/http"
    "ds3/net"
)

type GetObjectResponse struct {
    Content io.ReadCloser
}

func NewGetObjectResponse(response net.Response) (*GetObjectResponse, error) {
    if err := checkStatusCode(response, http.StatusOK); err != nil {
        return nil, err
    }
    return &GetObjectResponse{response.Body()}, nil
}

