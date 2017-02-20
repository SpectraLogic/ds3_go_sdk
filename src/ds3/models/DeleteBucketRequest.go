package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type DeleteBucketRequest struct {
    bucketName string
}

func NewDeleteBucketRequest(bucketName string) *DeleteBucketRequest {
    return &DeleteBucketRequest{
        bucketName: bucketName,
    }
}

func (DeleteBucketRequest) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (self DeleteBucketRequest) Path() string {
    return "/" + self.bucketName
}

func (self DeleteBucketRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (DeleteBucketRequest) Header() *http.Header {
    return &http.Header{}
}

func (DeleteBucketRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

