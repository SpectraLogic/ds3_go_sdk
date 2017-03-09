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
    checksum string
    checksumType networking.ChecksumType
}

func NewPutObjectRequest(bucketName string, objectName string, content networking.SizedReadCloser) *PutObjectRequest {
    return &PutObjectRequest{
        bucketName: bucketName,
        objectName:objectName,
        content:content,
        checksumType: networking.NONE } // Default checksum type of None
}

func (putObjectRequest *PutObjectRequest) WithChecksum(checksum string, checksumType networking.ChecksumType) *PutObjectRequest {
    putObjectRequest.checksum = checksum
    putObjectRequest.checksumType = checksumType
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

func (putObjectRequest *PutObjectRequest) GetChecksum() string {
    return putObjectRequest.checksum
}

func (putObjectRequest *PutObjectRequest) GetChecksumType() networking.ChecksumType {
    return putObjectRequest.checksumType
}