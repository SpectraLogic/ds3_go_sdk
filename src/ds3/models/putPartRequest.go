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
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

func NewPutPartRequest(
    bucketName string,
    objectName string,
    partNumber int,
    uploadId string,
    content networking.ReaderWithSizeDecorator,
) *PutPartRequest {
    queryParams := &url.Values{}
    queryParams.Set("partNumber", strconv.Itoa(partNumber))
    queryParams.Set("uploadId", uploadId)

    return &PutPartRequest{
        bucketName: bucketName,
        objectName: objectName,
        partNumber: partNumber,
        uploadId: uploadId,
        content: content,
        queryParams: queryParams,
    }
}

func (PutPartRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (putPartRequest *PutPartRequest) Path() string {
    return "/" + putPartRequest.bucketName + "/" + putPartRequest.objectName
}

func (putPartRequest *PutPartRequest) QueryParams() *url.Values {
    return putPartRequest.queryParams
}

func (PutPartRequest) Header() *http.Header {
    return &http.Header{}
}

func (putPartRequest *PutPartRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return putPartRequest.content
}

func (PutPartRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
