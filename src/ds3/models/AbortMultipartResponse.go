package models

import (
    "net/http"
    "ds3/networking"
)

type AbortMultipartResponse struct {}

func NewAbortMultipartResponse(response networking.Response) (*AbortMultipartResponse, error) {
    if err := checkStatusCode(response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &AbortMultipartResponse{}, nil
    }
}

