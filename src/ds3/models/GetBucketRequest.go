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

func (self *GetBucketRequest) WithMarker(marker string) *GetBucketRequest {
    self.marker = marker
    return self
}

func (self *GetBucketRequest) WithPrefix(prefix string) *GetBucketRequest {
    self.prefix = prefix
    return self
}

func (self *GetBucketRequest) WithMaxKeys(maxKeys int) *GetBucketRequest {
    self.maxKeys = maxKeys
    return self
}

func (GetBucketRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (self GetBucketRequest) Path() string {
    return "/" + self.bucketName
}

func (self GetBucketRequest) QueryParams() *url.Values {
    values := make(url.Values)
    if self.marker != "" {
        values.Add("marker", self.marker)
    }
    if self.prefix != "" {
        values.Add("prefix", self.prefix)
    }
    if self.maxKeys > 0 {
        values.Add("max-keys", strconv.Itoa(self.maxKeys))
    }
    return &values
}

func (GetBucketRequest) Header() *http.Header {
    return &http.Header{}
}

func (GetBucketRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}
