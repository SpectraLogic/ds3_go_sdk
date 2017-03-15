package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type BulkPutRequest struct {
    bucketName string
    content networking.SizedReadCloser
    queryParams *url.Values
}

func NewBulkPutRequest(bucketName string, objects []Object) *BulkPutRequest {
    queryParams := &url.Values{}
    queryParams.Set("operation", "start_bulk_put")

    return &BulkPutRequest{
        bucketName: bucketName,
        content: buildObjectListStream(objects),
        queryParams: queryParams,
    }
}

func (BulkPutRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (bulkPutRequest *BulkPutRequest) Path() string {
    return "/_rest_/buckets/" + bulkPutRequest.bucketName
}

func (bulkPutRequest *BulkPutRequest) QueryParams() *url.Values {
    return bulkPutRequest.queryParams
}

func (BulkPutRequest) Header() *http.Header {
    return &http.Header{}
}

func (bulkPutRequest *BulkPutRequest) GetContentStream() networking.SizedReadCloser {
    return bulkPutRequest.content
}

func (BulkPutRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
