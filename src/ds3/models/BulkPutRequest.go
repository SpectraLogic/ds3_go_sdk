package models

import (
    "net/url"
    "net/http"
    "ds3/net"
)

type BulkPutRequest struct {
    bucketName string
    content net.SizedReadCloser
}

func NewBulkPutRequest(bucketName string, objects []Object) *BulkPutRequest {
    return &BulkPutRequest{bucketName, buildObjectListStream(objects)}
}

func (BulkPutRequest) Verb() net.HttpVerb {
    return net.PUT
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

func (self BulkPutRequest) GetContentStream() net.SizedReadCloser {
    return self.content
}


