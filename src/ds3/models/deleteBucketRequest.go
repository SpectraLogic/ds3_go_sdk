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

func (deleteBucketRequest *DeleteBucketRequest) Path() string {
    return "/" + deleteBucketRequest.bucketName
}

func (deleteBucketRequest *DeleteBucketRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (DeleteBucketRequest) Header() *http.Header {
    return &http.Header{}
}

func (DeleteBucketRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (DeleteBucketRequest) GetChecksum() networking.Checksum {
    return networking.Checksum{
        Type: networking.NONE,
        ContentHash: "" }
}
