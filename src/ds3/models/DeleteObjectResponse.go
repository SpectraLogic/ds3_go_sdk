package models

import "net/http"

type DeleteObjectResponse struct {}

func NewDeleteObjectResponse(response *http.Response) (*DeleteObjectResponse, error) {
    if err := checkStatusCode(response, http.StatusNoContent); err != nil {
        return nil, err
    } else {
        return &DeleteObjectResponse{}, nil
    }
}


