package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type BulkGetRequest struct {
    bucketName string
    content networking.SizedReadCloser
}

func NewBulkGetRequest(bucketName string, objects []Object) *BulkGetRequest {
    return &BulkGetRequest{bucketName, buildObjectListStream(objects)}
}

func (BulkGetRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (self BulkGetRequest) Path() string {
    return "/_rest_/buckets/" + self.bucketName
}

func (BulkGetRequest) QueryParams() *url.Values {
    return &url.Values{"operation": []string{"start_bulk_get"}}
}

func (BulkGetRequest) Header() *http.Header {
    return &http.Header{}
}

func (self BulkGetRequest) GetContentStream() networking.SizedReadCloser {
    return self.content
}

