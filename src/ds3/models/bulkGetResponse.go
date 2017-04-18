package models

import (
    "ds3/networking"
)

type BulkGetResponse struct {
    Objects [][]Object
}

func NewBulkGetResponse(webResponse networking.WebResponse) (*BulkGetResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        objects, err := getObjectsFromBulkResponse(webResponse)
        return &BulkGetResponse{objects}, err
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }

}

