package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type ListPartsRequest struct {
    bucketName, objectName string
    uploadId string
}

func NewListPartsRequest(bucketName, objectName, uploadId string) *ListPartsRequest {
    return &ListPartsRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
    }
}

func (ListPartsRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (self ListPartsRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (self ListPartsRequest) QueryParams() *url.Values {
    return &url.Values{"uploadId": []string{self.uploadId}}
}

func (ListPartsRequest) Header() *http.Header {
    return &http.Header{}
}

func (ListPartsRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

