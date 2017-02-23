package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
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

func (InitiateMultipartRequest) Verb() networking.HttpVerb {
    return networking.POST
}

func (self InitiateMultipartRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (self InitiateMultipartRequest) QueryParams() *url.Values {
    return &url.Values{"uploads": []string{""}}
}

func (InitiateMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (InitiateMultipartRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}
