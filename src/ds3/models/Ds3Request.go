package models

import (
    "io"
    "net/url"
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
        default: return "GET"//TODO
    }
}

type Ds3Request interface {
    Verb() HttpVerb
    Path() string
    QueryParams() *url.Values
    GetContentStream() io.ReadCloser
}

