package net

import (
    "io"
    "strings"
    "net/http"
)

type Response interface {
    StatusCode() int
    Body() io.ReadCloser
    Header() map[string][]string
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

func (self wrappedHttpResponse) Header() map[string][]string {
    // The HTTP spec says headers keys are case insensitive, so we'll just
    // tolower them before processing the response so we can always get the
    // right thing.
    result := make(map[string][]string)
    header := self.rawResponse.Header
    for k, v := range header {
        result[strings.ToLower(k)] = v
    }
    return result
}

