package models

import (
    "ds3/networking"
)

type GetJobToReplicateSpectraS3Response struct {
    Content string
}

func NewGetJobToReplicateSpectraS3Response(webResponse networking.WebResponse) (*GetJobToReplicateSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        content, err := getResponseBodyAsString(webResponse)
        if err != nil {
            return nil, err
        }
        return &GetJobToReplicateSpectraS3Response{Content: content}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}