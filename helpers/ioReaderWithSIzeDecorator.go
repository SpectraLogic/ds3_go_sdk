package helpers

import "io"

// Defines a simple wrapper for an io.Reader and io.Closer which specifies size
// Implements ReaderWithSizeDecorator.
type IoReaderWithSizeDecorator struct {
    reader io.Reader
    size int64
}

func NewIoReaderWithSizeDecorator(reader io.Reader, size int64) *IoReaderWithSizeDecorator {
    return &IoReaderWithSizeDecorator{reader:reader, size:size}
}

func (ior *IoReaderWithSizeDecorator) Read(b []byte) (int, error) {
    return ior.reader.Read(b)
}

func (ior *IoReaderWithSizeDecorator) Size() (int64, error) {
    return ior.size, nil
}