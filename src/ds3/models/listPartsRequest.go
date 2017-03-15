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
    queryParams *url.Values
}

func NewListPartsRequest(bucketName string, objectName string, uploadId string) *ListPartsRequest {
    queryParams := &url.Values{}
    queryParams.Set("uploadId", uploadId)

    return &ListPartsRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
        queryParams: queryParams,
    }
}

func (ListPartsRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (listPartsRequest *ListPartsRequest) Path() string {
    return "/" + listPartsRequest.bucketName + "/" + listPartsRequest.objectName
}

func (listPartsRequest *ListPartsRequest) QueryParams() *url.Values {
    return listPartsRequest.queryParams
}

func (ListPartsRequest) Header() *http.Header {
    return &http.Header{}
}

func (ListPartsRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}

func (ListPartsRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
