package net

import (
    "io"
    "net/http"
)

type Response interface {
    StatusCode() int
    Body() io.ReadCloser
}

type wrappedHttpResponse struct {
    rawResponse *http.Response
}

func (self wrappedHttpResponse) StatusCode() int {
    return self.rawResponse.StatusCode
}

func (self wrappedHttpResponse) Body() io.ReadCloser {
    return self.rawResponse.Body
}

