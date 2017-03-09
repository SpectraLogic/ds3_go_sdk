package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type ListPartsRequest struct {
    bucketName string
    objectName string
    uploadId string
}

func NewListPartsRequest(bucketName string, objectName string, uploadId string) *ListPartsRequest {
    return &ListPartsRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
    }
}

func (ListPartsRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (listPartsRequest *ListPartsRequest) Path() string {
    return "/" + listPartsRequest.bucketName + "/" + listPartsRequest.objectName
}

func (listPartsRequest *ListPartsRequest) QueryParams() *url.Values {
    return &url.Values{"uploadId": []string{listPartsRequest.uploadId}}
}

func (ListPartsRequest) Header() *http.Header {
    return &http.Header{}
}

func (ListPartsRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (ListPartsRequest) GetChecksum() string {
    return ""
}

func (ListPartsRequest) GetChecksumType() networking.ChecksumType {
    return networking.NONE
}
