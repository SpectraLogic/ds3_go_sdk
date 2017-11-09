package models

import (
    "io"
    "strings"
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
    // The HTTP spec says headers keys are case insensitive, so we'll just
    // to lower them before processing the response so we can always get the
    // right thing.
    result := make(http.Header)
    header := wrappedHttpResponse.rawResponse.Header
    for k, v := range header {
        result[strings.ToLower(k)] = v
    }
    return &result
}

