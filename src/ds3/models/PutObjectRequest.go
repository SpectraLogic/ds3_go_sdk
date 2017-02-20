package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type PutObjectRequest struct {
    bucketName, objectName string
    content networking.SizedReadCloser
}

func NewPutObjectRequest(bucketName, objectName string, content networking.SizedReadCloser) *PutObjectRequest {
    return &PutObjectRequest{bucketName, objectName, content}
}

func (PutObjectRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (self PutObjectRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (PutObjectRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (PutObjectRequest) Header() *http.Header {
    return &http.Header{}
}

func (self PutObjectRequest) GetContentStream() networking.SizedReadCloser {
    return self.content
}

