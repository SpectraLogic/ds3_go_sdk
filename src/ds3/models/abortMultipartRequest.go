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
    queryParams *url.Values
}

func NewAbortMultipartRequest(bucketName string, objectName string, uploadId string) *AbortMultipartRequest {
    queryParams := &url.Values{}
    queryParams.Set("", uploadId)

    return &AbortMultipartRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
        queryParams: queryParams,
    }
}

func (AbortMultipartRequest) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (abortMultipartRequest *AbortMultipartRequest) Path() string {
    return "/" + abortMultipartRequest.bucketName + "/" + abortMultipartRequest.objectName
}

func (abortMultipartRequest *AbortMultipartRequest) QueryParams() *url.Values {
    return abortMultipartRequest.queryParams
}

func (AbortMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (AbortMultipartRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (AbortMultipartRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}