package models

import (
    "io"
    "net/http"
)

type GetObjectResponse struct {
    Content io.ReadCloser
}

func NewGetObjectResponse(response *http.Response) (*GetObjectResponse, error) {
    if err := checkStatusCode(response, http.StatusOK); err != nil {
        return nil, err
    }
    return &GetObjectResponse{response.Body}, nil
}

