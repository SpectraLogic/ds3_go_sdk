package models

import (
    "io"
    "io/ioutil"
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
    return ioutil.NopCloser(nopReader{})
}

//TODO: we might need this somewhere else as well
type nopReader struct {}
func (nopReader) Read(p []byte) (int, error) {
    return 0, nil
}

