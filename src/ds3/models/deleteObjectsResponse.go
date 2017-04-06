package models

import (
    "ds3/networking"
    "net/http"
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
    var body DeleteObjectsResponse //DeleteResult
    if err := readResponseBody(webResponse, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return &body, nil
}