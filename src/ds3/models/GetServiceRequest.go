package models

import (
    "io"
    "net/url"
)

type GetServiceRequest struct {}

func NewGetServiceRequest() *GetServiceRequest {
    return &GetServiceRequest{}
}

func (GetServiceRequest) Verb() HttpVerb {
    return GET
}

func (GetServiceRequest) Path() string {
    return "/"
}

func (GetServiceRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (GetServiceRequest) GetContentStream() io.ReadCloser {
    return nil
}

