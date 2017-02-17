package models

import (
    "net/url"
    "net/http"
    "ds3/network"
)

type ListMultipartRequest struct { }

func NewListMultipartRequest() *ListMultipartRequest {
    return &ListMultipartRequest{}
}

func (ListMultipartRequest) Verb() net.HttpVerb {
    return net.GET
}

func (ListMultipartRequest) Path() string {
    return "/"
}

func (ListMultipartRequest) QueryParams() *url.Values {
    return &url.Values{"uploads": []string{""}}
}

func (ListMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (ListMultipartRequest) GetContentStream() net.SizedReadCloser {
    return nil
}

