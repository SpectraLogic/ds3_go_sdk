package models

import (
    "fmt"
    "net/url"
    "net/http"
    "ds3/networking"
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

func (getObjectRequest *GetObjectRequest) WithRange(start, end int) *GetObjectRequest {
    getObjectRequest.rangeHeader = &rangeHeader{start, end}
    return getObjectRequest
}

func (GetObjectRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (getObjectRequest *GetObjectRequest) Path() string {
    return "/" + getObjectRequest.bucketName + "/" + getObjectRequest.objectName
}

func (GetObjectRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (getObjectRequest GetObjectRequest) Header() *http.Header {
    if getObjectRequest.rangeHeader == nil {
        return &http.Header{}
    } else {
        rng := fmt.Sprintf("bytes=%d-%d", getObjectRequest.rangeHeader.start, getObjectRequest.rangeHeader.end)
        return &http.Header{ "Range": []string{ rng } }
    }
}

func (GetObjectRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

