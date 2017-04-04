package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type CompleteMultipartRequest struct {
    bucketName string
    objectName string
    uploadId string
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

func NewCompleteMultipartRequest(bucketName string, objectName string, uploadId string, parts []Part) *CompleteMultipartRequest {
    queryParams := &url.Values{}
    queryParams.Set("uploadId", uploadId)

    return &CompleteMultipartRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
        content: buildPartsListStream(parts),
        queryParams: queryParams,
    }
}

func (CompleteMultipartRequest) Verb() networking.HttpVerb {
    return networking.POST
}

func (completeMultipartRequest *CompleteMultipartRequest) Path() string {
    return "/" + completeMultipartRequest.bucketName + "/" + completeMultipartRequest.objectName
}

func (completeMultipartRequest *CompleteMultipartRequest) QueryParams() *url.Values {
    return completeMultipartRequest.queryParams
}

func (CompleteMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (completeMultipartRequest *CompleteMultipartRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return completeMultipartRequest.content
}

func (CompleteMultipartRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}