package models

import (
    "net/url"
    "ds3/net"
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

func (GetServiceRequest) GetContentStream() net.SizedReadCloser {
    return nil
}

