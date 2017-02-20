package models

import (
    "ds3/networking"
)

type BulkGetResponse struct {
    Objects [][]Object
}

func NewBulkGetResponse(response networking.Response) (*BulkGetResponse, error) {
    objects, err := getObjectsFromBulkResponse(response)
    return &BulkGetResponse{objects}, err
}

