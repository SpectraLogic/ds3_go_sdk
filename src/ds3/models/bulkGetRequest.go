package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type BulkGetRequest struct {
    bucketName string
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

func NewBulkGetRequest(bucketName string, objects []Ds3Object) *BulkGetRequest {
    queryParams := &url.Values{}
    queryParams.Set("operation", "start_bulk_get")

    return &BulkGetRequest{
        bucketName: bucketName,
        content: buildDs3ObjectListStream(objects),
        queryParams: queryParams,
    }
}

func (BulkGetRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (bulkGetRequest *BulkGetRequest) Path() string {
    return "/_rest_/buckets/" + bulkGetRequest.bucketName
}

func (bulkGetRequest *BulkGetRequest) QueryParams() *url.Values {
    return bulkGetRequest.queryParams
}

func (BulkGetRequest) Header() *http.Header {
    return &http.Header{}
}

func (bulkGetRequest *BulkGetRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return bulkGetRequest.content
}

func (BulkGetRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
