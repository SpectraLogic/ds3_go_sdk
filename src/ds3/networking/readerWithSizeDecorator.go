package networking

import (
    "io"
    "ds3/models"
)

// Defines a limited reader that also supports closing
type LimitReadCloser struct {
    reader io.LimitedReader
    closer io.Closer
}

// Creates a ReadCloser that wraps the reader in a LimitedReader, and preserves
// the closability of the original reader
func NewLimitReadCloser(readCloser models.ReadCloserWithSizeDecorator) io.ReadCloser {
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
