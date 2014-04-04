package models

import "net/url"

type PutObjectRequest struct {
    bucketName, objectName string
    content SizedReadCloser
}

func NewPutObjectRequest(bucketName, objectName string, content SizedReadCloser) *PutObjectRequest {
    return &PutObjectRequest{bucketName, objectName, content}
}

func (PutObjectRequest) Verb() HttpVerb {
    return PUT
}

func (self PutObjectRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (PutObjectRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (self PutObjectRequest) GetContentStream() SizedReadCloser {
    return self.content
}

