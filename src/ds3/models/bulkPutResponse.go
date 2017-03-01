package models

import (
    "ds3/networking"
)

type BulkPutResponse struct {
    Objects [][]Object
}

func NewBulkPutResponse(webResponse networking.WebResponse) (*BulkPutResponse, error) {
    objects, err := getObjectsFromBulkResponse(webResponse)
    return &BulkPutResponse{objects}, err
}

