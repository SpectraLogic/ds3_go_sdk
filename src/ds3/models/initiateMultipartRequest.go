package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type InitiateMultipartRequest struct {
    bucketName string
    objectName string
}

func NewInitiateMultipartRequest(bucketName string, objectName string) *InitiateMultipartRequest {
    return &InitiateMultipartRequest{
        bucketName: bucketName,
        objectName: objectName,
    }
}

func (InitiateMultipartRequest) Verb() networking.HttpVerb {
    return networking.POST
}

func (initiateMultipartRequest *InitiateMultipartRequest) Path() string {
    return "/" + initiateMultipartRequest.bucketName + "/" + initiateMultipartRequest.objectName
}

func (initiateMultipartRequest *InitiateMultipartRequest) QueryParams() *url.Values {
    return &url.Values{"uploads": []string{""}}
}

func (InitiateMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (InitiateMultipartRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (InitiateMultipartRequest) GetChecksum() networking.Checksum {
    return networking.Checksum{
        Type: networking.NONE,
        ContentHash: "" }
}
