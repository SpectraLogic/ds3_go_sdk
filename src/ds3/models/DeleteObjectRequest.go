package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type DeleteObjectRequest struct {
    bucketName, objectName string
}

func NewDeleteObjectRequest(bucketName, objectName string) *DeleteObjectRequest {
    return &DeleteObjectRequest{bucketName, objectName}
}

func (DeleteObjectRequest) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (self DeleteObjectRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (self DeleteObjectRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (DeleteObjectRequest) Header() *http.Header {
    return &http.Header{}
}

func (DeleteObjectRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

