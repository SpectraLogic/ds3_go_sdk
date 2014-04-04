package models

import "net/url"

type BulkGetRequest struct {
    bucketName string
    content SizedReadCloser
}

func NewBulkGetRequest(bucketName string, objects []Object) *BulkGetRequest {
    return &BulkGetRequest{bucketName, buildMolReader(objects)}
}

func (BulkGetRequest) Verb() HttpVerb {
    return PUT
}

func (self BulkGetRequest) Path() string {
    return "/_rest_/buckets/" + self.bucketName
}

func (BulkGetRequest) QueryParams() *url.Values {
    return &url.Values{"operation": []string{"start_bulk_get"}}
}

func (self BulkGetRequest) GetContentStream() SizedReadCloser {
    return self.content
}

