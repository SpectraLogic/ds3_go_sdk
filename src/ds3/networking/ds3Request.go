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
    GetContentStream() ReaderWithSizeDecorator
    GetChecksum() Checksum
}

type Checksum struct {
    ContentHash string
    Type        ChecksumType
}

func NewNoneChecksum() Checksum {
    return Checksum{
        Type: NONE,
        ContentHash: "" }
}

// We need a Size method so we can pass the appropriate Content-Length header.
// Size isn't readily available in Go standard interfaces, so we created a new
// interface for it.
type ReaderWithSizeDecorator interface {
    io.Reader
    io.Closer
    Size() (int64, error)
}

type ChecksumType int

const(
    CRC_32 ChecksumType = iota
    CRC_32C
    MD5
    SHA_256
    SHA_512
    NONE
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

// Defines a ReaderWithSizeDecorator based on an array of bytes.
type byteReaderWithSizeDecorator struct {
    reader *bytes.Reader
    size int64
}

func BuildByteReaderWithSizeDecorator(contentBytes []byte) ReaderWithSizeDecorator {
    return &byteReaderWithSizeDecorator{
        bytes.NewReader(contentBytes),
        int64(len(contentBytes)),
    }
}

func (byteReaderWithSizeDecorator *byteReaderWithSizeDecorator) Read(b []byte) (int, error) {
    return byteReaderWithSizeDecorator.reader.Read(b)
}

func (byteReaderWithSizeDecorator) Close() error {
    return nil
}

func (byteReaderWithSizeDecorator *byteReaderWithSizeDecorator) Seek(offset int64, whence int) (int64, error) {
    return byteReaderWithSizeDecorator.reader.Seek(offset, whence)
}

func (byteReaderWithSizeDecorator *byteReaderWithSizeDecorator) Size() (int64, error) {
    return byteReaderWithSizeDecorator.size, nil
}

