package models

import (
    "ds3/networking"
)

type BulkGetResponse struct {
    Objects [][]Object
}

func NewBulkGetResponse(webResponse networking.WebResponse) (*BulkGetResponse, error) {
    objects, err := getObjectsFromBulkResponse(webResponse)
    return &BulkGetResponse{objects}, err
}

