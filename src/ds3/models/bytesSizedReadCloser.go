package models

import (
    "bytes"
    "ds3/net"
)

// Defines a SizedReadCloser based on an array of bytes.
type bytesSizedReadCloser struct {
    reader *bytes.Reader
    size int64
}

func buildSizedReadCloser(contentBytes []byte) net.SizedReadCloser{
    return &bytesSizedReadCloser{bytes.NewReader(contentBytes), int64(len(contentBytes))}
}

func (self bytesSizedReadCloser) Read(b []byte) (int, error) {
    return self.reader.Read(b)
}

func (bytesSizedReadCloser) Close() error {
    return nil
}

func (self bytesSizedReadCloser) Size() int64 {
    return self.size
}

