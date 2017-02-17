package models

import (
    "net/url"
    "net/http"
    "strconv"
    "ds3/network"
)

type PutPartRequest struct {
    bucketName, objectName string
    partNumber int
    uploadId string
    content net.SizedReadCloser
}

func NewPutPartRequest(
    bucketName, objectName string,
    partNumber int,
    uploadId string,
    content net.SizedReadCloser,
) *PutPartRequest {
    return &PutPartRequest{bucketName, objectName, partNumber, uploadId, content}
}

func (PutPartRequest) Verb() net.HttpVerb {
    return net.PUT
}

func (self PutPartRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (self PutPartRequest) QueryParams() *url.Values {
    return &url.Values{
        "partNumber": []string{strconv.Itoa(self.partNumber)},
        "uploadId": []string{self.uploadId},
    }
}

func (PutPartRequest) Header() *http.Header {
    return &http.Header{}
}

func (self PutPartRequest) GetContentStream() net.SizedReadCloser {
    return self.content
}

