package models

import "net/url"

type DeleteBucketRequest struct {
    bucketName string
}

func NewDeleteBucketRequest(bucketName string) *DeleteBucketRequest {
    return &DeleteBucketRequest{
        bucketName: bucketName,
    }
}

func (DeleteBucketRequest) Verb() HttpVerb {
    return DELETE
}

func (self DeleteBucketRequest) Path() string {
    return "/" + self.bucketName
}

func (self DeleteBucketRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (DeleteBucketRequest) GetContentStream() SizedReadCloser {
    return nil
}

