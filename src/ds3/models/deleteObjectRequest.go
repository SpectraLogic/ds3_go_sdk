package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type DeleteObjectRequest struct {
    bucketName string
    objectName string
    queryParams *url.Values
}

func NewDeleteObjectRequest(bucketName string, objectName string) *DeleteObjectRequest {
    return &DeleteObjectRequest{
        bucketName: bucketName,
        objectName: objectName,
        queryParams: &url.Values{},}
}

func (DeleteObjectRequest) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteObjectRequest *DeleteObjectRequest) Path() string {
    return "/" + deleteObjectRequest.bucketName + "/" + deleteObjectRequest.objectName
}

func (deleteObjectRequest *DeleteObjectRequest) QueryParams() *url.Values {
    return deleteObjectRequest.queryParams
}

func (DeleteObjectRequest) Header() *http.Header {
    return &http.Header{}
}

func (DeleteObjectRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}

func (DeleteObjectRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
