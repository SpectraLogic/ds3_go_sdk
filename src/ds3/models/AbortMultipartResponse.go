package models

import (
    "net/http"
    "ds3/network"
)

type AbortMultipartResponse struct {}

func NewAbortMultipartResponse(response net.Response) (*AbortMultipartResponse, error) {
    if err := checkStatusCode(response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &AbortMultipartResponse{}, nil
    }
}

