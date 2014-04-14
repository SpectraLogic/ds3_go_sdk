package models

import (
    "net/url"
    "ds3/net"
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

func (ListMultipartRequest) GetContentStream() net.SizedReadCloser {
    return nil
}




