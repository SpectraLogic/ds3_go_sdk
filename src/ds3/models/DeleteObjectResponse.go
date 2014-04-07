package models

import (
    "net/http"
    "ds3/net"
)

type DeleteObjectResponse struct {}

func NewDeleteObjectResponse(response net.Response) (*DeleteObjectResponse, error) {
    if err := checkStatusCode(response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &DeleteObjectResponse{}, nil
    }
}


