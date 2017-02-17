package models

import (
    "net/url"
    "net/http"
    "ds3/network"
)

type DeleteBucketRequest struct {
    bucketName string
}

func NewDeleteBucketRequest(bucketName string) *DeleteBucketRequest {
    return &DeleteBucketRequest{
        bucketName: bucketName,
    }
}

func (DeleteBucketRequest) Verb() net.HttpVerb {
    return net.DELETE
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

func (DeleteBucketRequest) GetContentStream() net.SizedReadCloser {
    return nil
}

