package net

import (
    "net/http"
    "bytes"
    "io"
    "net/url"
    "fmt"
)

type Request interface {
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
    io.Seeker
    Size() int64
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

func (verb HttpVerb) String() string {
    switch verb {
        case GET: return "GET"
        case PUT: return "PUT"
        case POST: return "POST"
        case DELETE: return "DELETE"
        case HEAD: return "HEAD"
        case PATCH: return "PATCH"
        default: panic(fmt.Sprintf("Invalid HttpVerb represented by: %d", verb))
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

func (self bytesSizedReadCloser) Read(b []byte) (int, error) {
    return self.reader.Read(b)
}

func (bytesSizedReadCloser) Close() error {
    return nil
}

func (self bytesSizedReadCloser) Seek(offset int64, whence int) (int64, error) {
    return self.reader.Seek(offset, whence)
}

func (self bytesSizedReadCloser) Size() int64 {
    return self.size
}

