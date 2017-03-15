package models

import (
    "fmt"
    "net/url"
    "net/http"
    "ds3/networking"
)

type GetObjectRequest struct {
    bucketName string
    objectName string
    rangeHeader *rangeHeader
    checksum networking.Checksum
    queryParams *url.Values
}

type rangeHeader struct {
    start, end int
}

func NewGetObjectRequest(bucketName string, objectName string) *GetObjectRequest {
    return &GetObjectRequest{
        bucketName: bucketName,
        objectName: objectName,
        checksum:   networking.NewNoneChecksum(), //Default checksum type of None
        queryParams: &url.Values{},
    }
}

func (getObjectRequest *GetObjectRequest) WithRange(start, end int) *GetObjectRequest {
    getObjectRequest.rangeHeader = &rangeHeader{start, end}
    return getObjectRequest
}

func (getObjectRequest *GetObjectRequest) WithChecksum(contentHash string, checksumType networking.ChecksumType) *GetObjectRequest {
    getObjectRequest.checksum.ContentHash = contentHash
    getObjectRequest.checksum.Type = checksumType
    return getObjectRequest
}

func (GetObjectRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (getObjectRequest *GetObjectRequest) Path() string {
    return "/" + getObjectRequest.bucketName + "/" + getObjectRequest.objectName
}

func (getObjectRequest *GetObjectRequest) QueryParams() *url.Values {
    return getObjectRequest.queryParams
}

func (getObjectRequest *GetObjectRequest) Header() *http.Header {
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

func (getObjectRequest *GetObjectRequest) GetChecksum() networking.Checksum {
    return getObjectRequest.checksum
}