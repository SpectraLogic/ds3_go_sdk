package models

import (
    "ds3/networking"
)

type BulkGetResponse struct {
    Objects [][]Object
}

func NewBulkGetResponse(ds3Response networking.Ds3Response) (*BulkGetResponse, error) {
    objects, err := getObjectsFromBulkResponse(ds3Response)
    return &BulkGetResponse{objects}, err
}

