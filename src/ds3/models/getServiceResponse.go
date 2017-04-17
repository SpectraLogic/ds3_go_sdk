package models

import (
    "ds3/networking"
)

type GetServiceResponse struct {
    Owner Owner
    Buckets []Bucket `xml:"Buckets>Bucket"`
}

type Bucket struct {
    Name string
    CreationDate string
}

func NewGetServiceResponse(webResponse networking.WebResponse) (*GetServiceResponse, error) {
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetServiceResponse
        if err := readResponseBody(webResponse, &body); err != nil {
            return nil, err
        }
        return &body, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
