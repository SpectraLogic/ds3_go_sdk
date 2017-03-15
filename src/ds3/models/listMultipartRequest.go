package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type ListMultipartRequest struct {
    queryParams *url.Values
}

func NewListMultipartRequest() *ListMultipartRequest {
    queryParams := &url.Values{}
    queryParams.Set("uploads", "")

    return &ListMultipartRequest{
        queryParams: queryParams,
    }
}

func (ListMultipartRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (ListMultipartRequest) Path() string {
    return "/"
}

func (listMultipartRequest *ListMultipartRequest) QueryParams() *url.Values {
    return listMultipartRequest.queryParams
}

func (ListMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (ListMultipartRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (ListMultipartRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}