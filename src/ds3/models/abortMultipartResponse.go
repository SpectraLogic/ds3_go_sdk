package models

import (
    "net/http"
    "ds3/networking"
)

type AbortMultipartResponse struct {}

func NewAbortMultipartResponse(ds3Response networking.Ds3Response) (*AbortMultipartResponse, error) {
    if err := checkStatusCode(ds3Response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &AbortMultipartResponse{}, nil
    }
}

