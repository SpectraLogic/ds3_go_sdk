package models

import (
    "net/url"
    "net/http"
    "strconv"
    "ds3/networking"
)

type GetBucketRequest struct {
    bucketName string
    marker string
    prefix string
    maxKeys int
}

func NewGetBucketRequest(bucketName string) *GetBucketRequest {
    return &GetBucketRequest{
        bucketName: bucketName,
    }
}

func (getBucketRequest *GetBucketRequest) WithMarker(marker string) *GetBucketRequest {
    getBucketRequest.marker = marker
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithPrefix(prefix string) *GetBucketRequest {
    getBucketRequest.prefix = prefix
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithMaxKeys(maxKeys int) *GetBucketRequest {
    getBucketRequest.maxKeys = maxKeys
    return getBucketRequest
}

func (GetBucketRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (getBucketRequest *GetBucketRequest) Path() string {
    return "/" + getBucketRequest.bucketName
}

func (getBucketRequest *GetBucketRequest) QueryParams() *url.Values {
    values := make(url.Values)
    if getBucketRequest.marker != "" {
        values.Add("marker", getBucketRequest.marker)
    }
    if getBucketRequest.prefix != "" {
        values.Add("prefix", getBucketRequest.prefix)
    }
    if getBucketRequest.maxKeys > 0 {
        values.Add("max-keys", strconv.Itoa(getBucketRequest.maxKeys))
    }
    return &values
}

func (GetBucketRequest) Header() *http.Header {
    return &http.Header{}
}

func (GetBucketRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

