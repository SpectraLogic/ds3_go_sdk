package models

import (
    "net/http"
    "ds3/networking"
)

type AbortMultipartResponse struct {}

func NewAbortMultipartResponse(webResponse networking.WebResponse) (*AbortMultipartResponse, error) {
    if err := checkStatusCode(webResponse, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &AbortMultipartResponse{}, nil
    }
}

