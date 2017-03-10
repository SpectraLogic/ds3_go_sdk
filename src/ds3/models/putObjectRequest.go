package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type PutObjectRequest struct {
    bucketName string
    objectName string
    content networking.SizedReadCloser
    checksum networking.Checksum
}

func NewPutObjectRequest(bucketName string, objectName string, content networking.SizedReadCloser) *PutObjectRequest {
    return &PutObjectRequest{
        bucketName: bucketName,
        objectName:objectName,
        content:content,
        checksum: networking.Checksum{ //Default checksum type of None
            Type: networking.NONE,
            ContentHash: "",
        }}
}

func (putObjectRequest *PutObjectRequest) WithChecksum(contentHash string, checksumType networking.ChecksumType) *PutObjectRequest {
    putObjectRequest.checksum.ContentHash = contentHash
    putObjectRequest.checksum.Type = checksumType
    return putObjectRequest
}

func (PutObjectRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (putObjectRequest *PutObjectRequest) Path() string {
    return "/" + putObjectRequest.bucketName + "/" + putObjectRequest.objectName
}

func (PutObjectRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (PutObjectRequest) Header() *http.Header {
    return &http.Header{}
}

func (putObjectRequest *PutObjectRequest) GetContentStream() networking.SizedReadCloser {
    return putObjectRequest.content
}

func (putObjectRequest *PutObjectRequest) GetChecksum() networking.Checksum {
    return putObjectRequest.checksum
}