package models

import (
"net/url"
"net/http"
"ds3/networking"
)

type DeleteObjectsRequest struct {
    bucketName string
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

func NewDeleteObjectsRequest(bucketName string, objectNames []string) *DeleteObjectsRequest {
    queryParams := &url.Values{}
    queryParams.Set("delete", "")

    return &DeleteObjectsRequest{
        bucketName:  bucketName,
        content:     buildDeleteObjectsPayload(objectNames),
        queryParams: queryParams,
    }
}

func (deleteObjectsRequest *DeleteObjectsRequest) WithRollBack() *DeleteObjectsRequest {
    deleteObjectsRequest.queryParams.Set("roll_back", "")
    return deleteObjectsRequest
}


func (DeleteObjectsRequest) Verb() networking.HttpVerb {
    return networking.POST
}

func (deleteObjectsRequest *DeleteObjectsRequest) Path() string {
    return "/" + deleteObjectsRequest.bucketName
}

func (deleteObjectsRequest *DeleteObjectsRequest) QueryParams() *url.Values {
    return deleteObjectsRequest.queryParams
}

func (DeleteObjectsRequest) Header() *http.Header {
    return &http.Header{}
}

func (DeleteObjectsRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}

func (deleteObjectsRequest *DeleteObjectsRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return deleteObjectsRequest.content
}