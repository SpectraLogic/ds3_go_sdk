package networking

import (
    "net/http"
    "bytes"
    "io"
    "net/url"
    "fmt"
    "errors"
)

type Ds3Request interface {
    Verb() HttpVerb
    Path() string
    QueryParams() *url.Values
    Header() *http.Header
    GetContentStream() SizedReadCloser
}

// We need a Size method so we can pass the appropriate Content-Length header.
// Size isn't readily available in Go standard interfaces, so we created a new
// interface for it.
type SizedReadCloser interface {
    io.Reader
    io.Closer
    Size() (int64, error)
}

type HttpVerb int

const(
    GET HttpVerb = iota
    PUT
    POST
    DELETE
    HEAD
    PATCH
)

func (verb HttpVerb) String() (string, error) {
    switch verb {
        case GET: return "GET", nil
        case PUT: return "PUT", nil
        case POST: return "POST", nil
        case DELETE: return "DELETE", nil
        case HEAD: return "HEAD", nil
        case PATCH: return "PATCH", nil
        default: return "", errors.New(fmt.Sprintf("Invalid HttpVerb represented by: %d", verb))
    }
}

// Defines a SizedReadCloser based on an array of bytes.
type bytesSizedReadCloser struct {
    reader *bytes.Reader
    size int64
}

func BuildSizedReadCloser(contentBytes []byte) SizedReadCloser{
    return &bytesSizedReadCloser{
        bytes.NewReader(contentBytes),
        int64(len(contentBytes)),
    }
}

func (bytesSizedReadCloser *bytesSizedReadCloser) Read(b []byte) (int, error) {
    return bytesSizedReadCloser.reader.Read(b)
}

func (bytesSizedReadCloser) Close() error {
    return nil
}

func (bytesSizedReadCloser *bytesSizedReadCloser) Seek(offset int64, whence int) (int64, error) {
    return bytesSizedReadCloser.reader.Seek(offset, whence)
}

func (bytesSizedReadCloser *bytesSizedReadCloser) Size() (int64, error) {
    return bytesSizedReadCloser.size, nil
}

