package models

import (
    "net/url"
    "ds3/net"
)

type GetObjectRequest struct {
    bucketName, objectName string
}

func NewGetObjectRequest(bucketName, objectName string) *GetObjectRequest {
    return &GetObjectRequest{bucketName, objectName}
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

func (GetObjectRequest) GetContentStream() net.SizedReadCloser {
    return nil
}

