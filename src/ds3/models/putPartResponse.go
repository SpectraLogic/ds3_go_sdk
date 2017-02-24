package models

import (
    "strings"
    "net/http"
    "ds3/networking"
)

type PutPartResponse struct {
    ETag string
}

func NewPutPartResponse(ds3Response networking.Ds3Response) (*PutPartResponse, error) {
    if err := checkStatusCode(ds3Response, http.StatusOK); err != nil {
        return nil, err
    } else {
        etags := (*ds3Response.Header())["etag"]
        var etag string
        if len(etags) > 0 {
            etag = strings.Trim(etags[0], "\"")
        }
        return &PutPartResponse{etag}, nil
    }
}

