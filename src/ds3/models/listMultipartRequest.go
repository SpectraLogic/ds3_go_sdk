package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type ListMultipartRequest struct { }

func NewListMultipartRequest() *ListMultipartRequest {
    return &ListMultipartRequest{}
}

func (ListMultipartRequest) Verb() networking.HttpVerb {
    return networking.GET
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

func (ListMultipartRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (ListMultipartRequest) GetChecksum() string {
    return ""
}

func (ListMultipartRequest) GetChecksumType() networking.ChecksumType {
    return networking.NONE
}