package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type DeleteObjectRequest struct {
    bucketName string
    objectName string
}

func NewDeleteObjectRequest(bucketName string, objectName string) *DeleteObjectRequest {
    return &DeleteObjectRequest{bucketName, objectName}
}

func (DeleteObjectRequest) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteObjectRequest *DeleteObjectRequest) Path() string {
    return "/" + deleteObjectRequest.bucketName + "/" + deleteObjectRequest.objectName
}

func (deleteObjectRequest *DeleteObjectRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (DeleteObjectRequest) Header() *http.Header {
    return &http.Header{}
}

func (DeleteObjectRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (DeleteObjectRequest) GetChecksum() string {
    return ""
}

func (DeleteObjectRequest) GetChecksumType() networking.ChecksumType {
    return networking.NONE
}
