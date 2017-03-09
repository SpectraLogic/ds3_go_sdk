package models

import (
    "net/url"
    "net/http"
    "strconv"
    "ds3/networking"
)

type PutPartRequest struct {
    bucketName string
    objectName string
    partNumber int
    uploadId string
    content networking.SizedReadCloser
}

func NewPutPartRequest(
    bucketName string,
    objectName string,
    partNumber int,
    uploadId string,
    content networking.SizedReadCloser,
) *PutPartRequest {
    return &PutPartRequest{bucketName, objectName, partNumber, uploadId, content}
}

func (PutPartRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (putPartRequest *PutPartRequest) Path() string {
    return "/" + putPartRequest.bucketName + "/" + putPartRequest.objectName
}

func (putPartRequest *PutPartRequest) QueryParams() *url.Values {
    return &url.Values{
        "partNumber": []string{strconv.Itoa(putPartRequest.partNumber)},
        "uploadId": []string{putPartRequest.uploadId},
    }
}

func (PutPartRequest) Header() *http.Header {
    return &http.Header{}
}

func (putPartRequest *PutPartRequest) GetContentStream() networking.SizedReadCloser {
    return putPartRequest.content
}

func (PutPartRequest) GetChecksum() string {
    return ""
}

func (PutPartRequest) GetChecksumType() networking.ChecksumType {
    return networking.NONE
}
