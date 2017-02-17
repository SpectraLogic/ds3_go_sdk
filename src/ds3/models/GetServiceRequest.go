package models

import (
    "net/url"
    "net/http"
    "ds3/network"
)

type GetServiceRequest struct {}

func NewGetServiceRequest() *GetServiceRequest {
    return &GetServiceRequest{}
}

func (GetServiceRequest) Verb() net.HttpVerb {
    return net.GET
}

func (GetServiceRequest) Path() string {
    return "/"
}

func (GetServiceRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (GetServiceRequest) Header() *http.Header {
    return &http.Header{}
}

func (GetServiceRequest) GetContentStream() net.SizedReadCloser {
    return nil
}

