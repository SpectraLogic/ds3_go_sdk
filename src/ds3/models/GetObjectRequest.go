package models

import (
    "fmt"
    "net/url"
    "net/http"
    "ds3/net"
)

type GetObjectRequest struct {
    bucketName, objectName string
    rangeHeader *rangeHeader
}

type rangeHeader struct {
    start, end int
}

func NewGetObjectRequest(bucketName, objectName string) *GetObjectRequest {
    return &GetObjectRequest{bucketName, objectName, nil}
}

func (self *GetObjectRequest) WithRange(start, end int) *GetObjectRequest {
    self.rangeHeader = &rangeHeader{start, end}
    return self
}

func (GetObjectRequest) Verb() net.HttpVerb {
    return net.GET
}

func (self GetObjectRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (GetObjectRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (self GetObjectRequest) Header() *http.Header {
    if self.rangeHeader == nil {
        return &http.Header{}
    } else {
        rng := fmt.Sprintf("bytes=%d-%d", self.rangeHeader.start, self.rangeHeader.end)
        return &http.Header{ "Range": []string{ rng } }
    }
}

func (GetObjectRequest) GetContentStream() net.SizedReadCloser {
    return nil
}

