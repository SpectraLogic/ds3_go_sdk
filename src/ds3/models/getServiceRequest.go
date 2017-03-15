package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type GetServiceRequest struct {
    queryParams *url.Values
}

func NewGetServiceRequest() *GetServiceRequest {
    return &GetServiceRequest{
        queryParams: &url.Values{},
    }
}

func (GetServiceRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (GetServiceRequest) Path() string {
    return "/"
}

func (getServiceRequest *GetServiceRequest) QueryParams() *url.Values {
    return getServiceRequest.queryParams
}

func (GetServiceRequest) Header() *http.Header {
    return &http.Header{}
}

func (GetServiceRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (GetServiceRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
