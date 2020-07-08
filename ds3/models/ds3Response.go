package models

import (
    "io"
    "net/http"
)

type WebResponse interface {
    StatusCode() int
    Body() io.ReadCloser
    Header() *http.Header
}

type WrappedHttpResponse struct {
    rawResponse *http.Response
}

func NewWrappedHttpResponse(rawResponse *http.Response) *WrappedHttpResponse {
    return &WrappedHttpResponse{rawResponse: rawResponse}
}

func (wrappedHttpResponse *WrappedHttpResponse) StatusCode() int {
    return wrappedHttpResponse.rawResponse.StatusCode
}

func (wrappedHttpResponse *WrappedHttpResponse) Body() io.ReadCloser {
    return wrappedHttpResponse.rawResponse.Body
}

func (wrappedHttpResponse *WrappedHttpResponse) Header() *http.Header {
    result := make(http.Header)
    header := wrappedHttpResponse.rawResponse.Header
    for k, v := range header {
        result[k] = v
    }
    return &result
}

