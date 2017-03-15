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
    queryParams *url.Values
}

func NewGetBucketRequest(bucketName string) *GetBucketRequest {
    return &GetBucketRequest{
        bucketName: bucketName,
        queryParams: &url.Values{},
    }
}

func (getBucketRequest *GetBucketRequest) WithMarker(marker string) *GetBucketRequest {
    getBucketRequest.marker = marker
    getBucketRequest.queryParams.Set("marker", marker)
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithPrefix(prefix string) *GetBucketRequest {
    getBucketRequest.prefix = prefix
    getBucketRequest.queryParams.Set("prefix", prefix)
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithMaxKeys(maxKeys int) *GetBucketRequest {
    getBucketRequest.maxKeys = maxKeys
    getBucketRequest.queryParams.Set("max-keys", strconv.Itoa(getBucketRequest.maxKeys))
    return getBucketRequest
}

func (GetBucketRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (getBucketRequest *GetBucketRequest) Path() string {
    return "/" + getBucketRequest.bucketName
}

func (getBucketRequest *GetBucketRequest) QueryParams() *url.Values {
    return getBucketRequest.queryParams
}

func (GetBucketRequest) Header() *http.Header {
    return &http.Header{}
}

func (GetBucketRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (GetBucketRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
