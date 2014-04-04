package models

import "net/url"

type PutBucketRequest struct {
    bucketName string
}

func NewPutBucketRequest(bucketName string) *PutBucketRequest {
    return &PutBucketRequest{
        bucketName: bucketName,
    }
}

func (PutBucketRequest) Verb() HttpVerb {
    return PUT
}

func (self PutBucketRequest) Path() string {
    return "/" + self.bucketName
}

func (self PutBucketRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (PutBucketRequest) GetContentStream() SizedReadCloser {
    return nil
}


