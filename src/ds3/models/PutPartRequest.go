package models

import (
    "net/url"
    "net/http"
    "strconv"
    "ds3/networking"
)

type PutPartRequest struct {
    bucketName, objectName string
    partNumber int
    uploadId string
    content networking.SizedReadCloser
}

func NewPutPartRequest(
    bucketName, objectName string,
    partNumber int,
    uploadId string,
    content networking.SizedReadCloser,
) *PutPartRequest {
    return &PutPartRequest{bucketName, objectName, partNumber, uploadId, content}
}

func (PutPartRequest) Verb() networking.HttpVerb {
    return networking.PUT
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

func (self PutPartRequest) GetContentStream() networking.SizedReadCloser {
    return self.content
}

