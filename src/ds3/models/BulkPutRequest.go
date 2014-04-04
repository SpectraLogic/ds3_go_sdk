package models

import "net/url"

type BulkPutRequest struct {
    bucketName string
    content SizedReadCloser
}

func NewBulkPutRequest(bucketName string, objects []Object) *BulkPutRequest {
    return &BulkPutRequest{bucketName, buildMolReader(objects)}
}

func (BulkPutRequest) Verb() HttpVerb {
    return PUT
}

func (self BulkPutRequest) Path() string {
    return "/_rest_/buckets/" + self.bucketName
}

func (BulkPutRequest) QueryParams() *url.Values {
    return &url.Values{"operation": []string{"start_bulk_put"}}
}

func (self BulkPutRequest) GetContentStream() SizedReadCloser {
    return self.content
}


