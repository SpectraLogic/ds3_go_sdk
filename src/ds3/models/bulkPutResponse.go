package models

import (
    "ds3/networking"
)

type BulkPutResponse struct {
    Objects [][]Object
}

func NewBulkPutResponse(webResponse networking.WebResponse) (*BulkPutResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        objects, err := getObjectsFromBulkResponse(webResponse)
        return &BulkPutResponse{objects}, err
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }

}

