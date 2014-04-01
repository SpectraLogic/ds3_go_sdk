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

type StatusCode int

const(
    OK StatusCode = 200
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

type Ds3Request interface {
    Verb() HttpVerb
    Path() string
    QueryParams() *url.Values
    GetContentStream() io.ReadCloser
}

