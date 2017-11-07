package networking

import (
"io"
)

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

// Defines a limited reader that also supports closing
type LimitReadCloser struct {
    reader io.LimitedReader
    closer io.Closer
}

// Creates a ReadCloser that wraps the reader in a LimitedReader, and preserves
// the closability of the original reader
func NewLimitReadCloser(readCloser ReadCloserWithSizeDecorator) io.ReadCloser {
    size, _ := readCloser.Size()
    limitReader := io.LimitedReader{R:readCloser, N:size}
    return &LimitReadCloser{
        reader: limitReader,
        closer: readCloser,
    }
}

// Reads from the limited reader
func (limitedReadCloser *LimitReadCloser) Read(p []byte) (n int, err error) {
    return limitedReadCloser.reader.Read(p)
}

// Closes the original reader
func (limitedReadCloser *LimitReadCloser) Close() error {
    return limitedReadCloser.closer.Close()
}
