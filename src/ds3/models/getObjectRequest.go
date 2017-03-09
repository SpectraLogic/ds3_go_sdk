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
    checksum string
    checksumType networking.ChecksumType
}

type rangeHeader struct {
    start, end int
}

func NewGetObjectRequest(bucketName string, objectName string) *GetObjectRequest {
    return &GetObjectRequest{
        bucketName:   bucketName,
        objectName:   objectName,
        checksumType: networking.NONE } //Default checksum type of None
}

func (getObjectRequest *GetObjectRequest) WithRange(start, end int) *GetObjectRequest {
    getObjectRequest.rangeHeader = &rangeHeader{start, end}
    return getObjectRequest
}

func (getObjectRequest *GetObjectRequest) WithChecksum(checksum string, checksumType networking.ChecksumType) *GetObjectRequest {
    getObjectRequest.checksum = checksum
    getObjectRequest.checksumType = checksumType
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

func (getObjectRequest *GetObjectRequest) GetChecksum() string {
    return getObjectRequest.checksum
}

func (getObjectRequest *GetObjectRequest) GetChecksumType() networking.ChecksumType {
    return getObjectRequest.checksumType
}
