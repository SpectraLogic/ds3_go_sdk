package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type BulkPutRequest struct {
    bucketName string
    content networking.SizedReadCloser
}

func NewBulkPutRequest(bucketName string, objects []Object) *BulkPutRequest {
    return &BulkPutRequest{bucketName, buildObjectListStream(objects)}
}

func (BulkPutRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (self BulkPutRequest) Path() string {
    return "/_rest_/buckets/" + self.bucketName
}

func (BulkPutRequest) QueryParams() *url.Values {
    return &url.Values{"operation": []string{"start_bulk_put"}}
}

func (BulkPutRequest) Header() *http.Header {
    return &http.Header{}
}

func (self BulkPutRequest) GetContentStream() networking.SizedReadCloser {
    return self.content
}


