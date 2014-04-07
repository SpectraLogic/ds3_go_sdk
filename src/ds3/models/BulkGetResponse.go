package models

import "net/http"

type BulkGetResponse struct {
    Objects [][]Object
}

func NewBulkGetResponse(response *http.Response) (*BulkGetResponse, error) {
    objects, err := getObjectsFromBulkResponse(response)
    return &BulkGetResponse{objects}, err
}


