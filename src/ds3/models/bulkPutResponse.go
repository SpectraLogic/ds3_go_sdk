package models

import (
    "ds3/networking"
)

type BulkPutResponse struct {
    Objects [][]Object
}

func NewBulkPutResponse(ds3Response networking.Ds3Response) (*BulkPutResponse, error) {
    objects, err := getObjectsFromBulkResponse(ds3Response)
    return &BulkPutResponse{objects}, err
}

