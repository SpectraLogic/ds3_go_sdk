package channels

import "io"

type Ds3ObjectReadCloser struct {
	io.Reader
}

func NewDs3ObjectReadCloserDecorator(reader io.Reader) io.ReadCloser {
	return Ds3ObjectReadCloser { Reader: reader }
}

func (Ds3ObjectReadCloser Ds3ObjectReadCloser) Close() error {
	readCloser, isReadCloser := Ds3ObjectReadCloser.Reader.(io.ReadCloser)
	if isReadCloser && readCloser != nil {
		return readCloser.Close()
	}
	return nil
}




