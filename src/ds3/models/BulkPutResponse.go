package models

import "ds3/network"

type BulkPutResponse struct {
    Objects [][]Object
}

func NewBulkPutResponse(response net.Response) (*BulkPutResponse, error) {
    objects, err := getObjectsFromBulkResponse(response)
    return &BulkPutResponse{objects}, err
}

