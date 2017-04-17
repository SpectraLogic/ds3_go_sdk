package models

import (
    "ds3/networking"
)

type BulkPutResponse struct {
    Objects [][]Object
}

func NewBulkPutResponse(webResponse networking.WebResponse) (*BulkPutResponse, error) {
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        objects, err := getObjectsFromBulkResponse(webResponse)
        return &BulkPutResponse{objects}, err
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }

}

