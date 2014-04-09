package models

import (
    "net/url"
    "ds3/net"
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

func (ListPartsRequest) Verb() net.HttpVerb {
    return net.GET
}

func (self ListPartsRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (self ListPartsRequest) QueryParams() *url.Values {
    return &url.Values{"uploadId": []string{self.uploadId}}
}

func (ListPartsRequest) GetContentStream() net.SizedReadCloser {
    return nil
}


