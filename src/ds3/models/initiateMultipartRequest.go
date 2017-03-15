package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type InitiateMultipartRequest struct {
    bucketName string
    objectName string
    queryParams *url.Values
}

func NewInitiateMultipartRequest(bucketName string, objectName string) *InitiateMultipartRequest {
    queryParams := &url.Values{}
    queryParams.Set("uploads", "")

    return &InitiateMultipartRequest{
        bucketName: bucketName,
        objectName: objectName,
        queryParams: queryParams,
    }
}

func (InitiateMultipartRequest) Verb() networking.HttpVerb {
    return networking.POST
}

func (initiateMultipartRequest *InitiateMultipartRequest) Path() string {
    return "/" + initiateMultipartRequest.bucketName + "/" + initiateMultipartRequest.objectName
}

func (initiateMultipartRequest *InitiateMultipartRequest) QueryParams() *url.Values {
    return initiateMultipartRequest.queryParams
}

func (InitiateMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (InitiateMultipartRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (InitiateMultipartRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
