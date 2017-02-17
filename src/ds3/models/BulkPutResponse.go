package models

import (
    "ds3/networking"
)

type BulkPutResponse struct {
    Objects [][]Object
}

func NewBulkPutResponse(response networking.Response) (*BulkPutResponse, error) {
    objects, err := getObjectsFromBulkResponse(response)
    return &BulkPutResponse{objects}, err
}

