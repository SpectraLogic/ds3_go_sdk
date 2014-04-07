package models

import (
    "io"
    "net/url"
    "fmt"
)

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

// We need a Size method so we can pass the appropriate Content-Length header.
// Size isn't readily available in Go standard interfaces, so we created a new
// interface for it.
type SizedReadCloser interface {
    io.ReadCloser
    Size() int64
}

type Ds3Request interface {
    Verb() HttpVerb
    Path() string
    QueryParams() *url.Values
    GetContentStream() SizedReadCloser
}

