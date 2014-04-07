package models

import "net/http"

type BulkPutResponse struct {
    Objects [][]Object
}

func NewBulkPutResponse(response *http.Response) (*BulkPutResponse, error) {
    objects, err := getObjectsFromBulkResponse(response)
    return &BulkPutResponse{objects}, err
}

