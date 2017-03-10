package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type AbortMultipartRequest struct {
    bucketName string
    objectName string
    uploadId string
}

func NewAbortMultipartRequest(bucketName string, objectName string, uploadId string) *AbortMultipartRequest {
    return &AbortMultipartRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
    }
}

func (AbortMultipartRequest) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (abortMultipartRequest *AbortMultipartRequest) Path() string {
    return "/" + abortMultipartRequest.bucketName + "/" + abortMultipartRequest.objectName
}

func (abortMultipartRequest *AbortMultipartRequest) QueryParams() *url.Values {
    return &url.Values{"uploadId": []string{abortMultipartRequest.uploadId}}
}

func (AbortMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (AbortMultipartRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (AbortMultipartRequest) GetChecksum() networking.Checksum {
    return networking.Checksum{
        Type: networking.NONE,
        ContentHash: "" }
}