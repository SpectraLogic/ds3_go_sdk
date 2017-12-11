package models

import "io"

// We need a Size method so we can pass the appropriate Content-Length header.
// Size isn't readily available in Go standard interfaces, so we created a new
// interface for it.
// Also hides the Close function from http library which prevents automatic
// closing of stream within http library.
type ReaderWithSizeDecorator interface {
    io.Reader
    Size() (int64, error)
}

// ReadCloser with size knowledge. Networking layer will automatically close
// the reader when finished transmitting request.
type ReadCloserWithSizeDecorator interface {
    io.ReadCloser
    Size() (int64, error)
}