package models

import (
    "ds3/networking"
)

type DeleteObjectsResponse struct {
    Deleted []DeletedObject `xml:"Deleted"`
    Errors []DeleteError `xml:"Error"`
}

type DeletedObject struct {
    Key string `xml:"Key"`
}

type DeleteError struct {
    Code string `xml:"Code"`
    Key string `xml:"Key"`
    Message string `xml:"Message"`
}

func NewDeleteObjectsResponse(webResponse networking.WebResponse) (*DeleteObjectsResponse, error) {
    expectedStatusCodes := []int { 200 }

    if err := checkStatusCode(webResponse, expectedStatusCodes); err != nil {
        return nil, err
    }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body DeleteObjectsResponse //DeleteResult
        if err := readResponseBody(webResponse, &body); err != nil {
            return nil, err
        }
        return &body, nil
    default:
        //Should never get here
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}