package models

import "ds3/network"

type BulkGetResponse struct {
    Objects [][]Object
}

func NewBulkGetResponse(response net.Response) (*BulkGetResponse, error) {
    objects, err := getObjectsFromBulkResponse(response)
    return &BulkGetResponse{objects}, err
}

