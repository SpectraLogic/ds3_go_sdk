package models

import "net/http"

type BulkGetResponse struct {
    Objects [][]Object
}

func NewBulkGetResponse(response *http.Response) (*BulkGetResponse, error) {
    objects, err := newBulkResponse(response)
    return &BulkGetResponse{objects}, err
}


