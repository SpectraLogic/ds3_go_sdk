package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type DeleteBucketRequest struct {
    bucketName string
    queryParams *url.Values
}

func NewDeleteBucketRequest(bucketName string) *DeleteBucketRequest {
    return &DeleteBucketRequest{
        bucketName: bucketName,
        queryParams: &url.Values{},
    }
}

func (DeleteBucketRequest) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteBucketRequest *DeleteBucketRequest) Path() string {
    return "/" + deleteBucketRequest.bucketName
}

func (deleteBucketRequest *DeleteBucketRequest) QueryParams() *url.Values {
    return deleteBucketRequest.queryParams
}

func (DeleteBucketRequest) Header() *http.Header {
    return &http.Header{}
}

func (DeleteBucketRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (DeleteBucketRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
