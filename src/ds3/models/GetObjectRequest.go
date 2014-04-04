package models

import (
    "io"
    "net/url"
)

type GetObjectRequest struct {
    bucketName, objectName string
}

func NewGetObjectRequest(bucketName, objectName string) *GetObjectRequest {
    return &GetObjectRequest{bucketName, objectName}
}

func (GetObjectRequest) Verb() HttpVerb {
    return GET
}

func (self GetObjectRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (GetObjectRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (GetObjectRequest) GetContentStream() io.ReadCloser {
    return nil
}

