package networking

import (
    "io"
    "strings"
    "net/http"
)

type Response interface {
    StatusCode() int
    Body() io.ReadCloser
    Header() *http.Header
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

func (self wrappedHttpResponse) Header() *http.Header {
    // The HTTP spec says headers keys are case insensitive, so we'll just
    // tolower them before processing the response so we can always get the
    // right thing.
    result := make(http.Header)
    header := self.rawResponse.Header
    for k, v := range header {
        result[strings.ToLower(k)] = v
    }
    return &result
}

