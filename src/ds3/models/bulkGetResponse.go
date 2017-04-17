package models

import (
    "ds3/networking"
)

type BulkGetResponse struct {
    Objects [][]Object
}

func NewBulkGetResponse(webResponse networking.WebResponse) (*BulkGetResponse, error) {
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        objects, err := getObjectsFromBulkResponse(webResponse)
        return &BulkGetResponse{objects}, err
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }

}

