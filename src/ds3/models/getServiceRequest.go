package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type GetServiceRequest struct {}

func NewGetServiceRequest() *GetServiceRequest {
    return &GetServiceRequest{}
}

func (GetServiceRequest) Verb() networking.HttpVerb {
    return networking.GET
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

func (GetServiceRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (GetServiceRequest) GetChecksum() networking.Checksum {
    return networking.Checksum{
        Type: networking.NONE,
        ContentHash: "" }
}
