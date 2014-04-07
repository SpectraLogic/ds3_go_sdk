package models

import (
    "net/url"
    "ds3/net"
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

func (DeleteBucketRequest) GetContentStream() net.SizedReadCloser {
    return nil
}

