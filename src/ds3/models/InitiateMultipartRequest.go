package models

import (
    "net/url"
    "ds3/net"
)

type InitiateMultipartRequest struct {
    bucketName, objectName string
}

func NewInitiateMultipartRequest(bucketName, objectName string) *InitiateMultipartRequest {
    return &InitiateMultipartRequest{
        bucketName: bucketName,
        objectName: objectName,
    }
}

func (InitiateMultipartRequest) Verb() net.HttpVerb {
    return net.POST
}

func (self InitiateMultipartRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (self InitiateMultipartRequest) QueryParams() *url.Values {
    return &url.Values{"uploads": []string{""}}
}

func (InitiateMultipartRequest) GetContentStream() net.SizedReadCloser {
    return nil
}



