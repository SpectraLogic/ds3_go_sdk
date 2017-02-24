package networking

import (
    "io"
    "strings"
    "net/http"
)

type Ds3Response interface {
    StatusCode() int
    Body() io.ReadCloser
    Header() *http.Header
}

type wrappedHttpResponse struct {
    rawResponse *http.Response
}

func (wrappedHttpResponse *wrappedHttpResponse) StatusCode() int {
    return wrappedHttpResponse.rawResponse.StatusCode
}

func (wrappedHttpResponse *wrappedHttpResponse) Body() io.ReadCloser {
    return wrappedHttpResponse.rawResponse.Body
}

func (wrappedHttpResponse *wrappedHttpResponse) Header() *http.Header {
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

