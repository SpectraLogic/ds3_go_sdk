package commands

import "os"

// Implements the SizedReadCloser interface for a file.
type fileSizeReadCloser struct {
    file *os.File
}

func readFile(path string) (*fileSizeReadCloser, error) {
    content, fileErr := os.Open(path)
    if fileErr != nil {
        return nil, fileErr
    }
    return &fileSizeReadCloser{content}, nil
}

func (fileSizeReadCloser *fileSizeReadCloser) Size() int64 {
    fi, err := fileSizeReadCloser.file.Stat()
    if err != nil {
        panic(err)
    }
    return fi.Size()
}

func (fileSizeReadCloser *fileSizeReadCloser) Read(p []byte) (int, error) {
    return fileSizeReadCloser.file.Read(p)
}

func (fileSizeReadCloser *fileSizeReadCloser) Seek(offset int64, whence int) (int64, error) {
    return fileSizeReadCloser.file.Seek(offset, whence)
}

func (fileSizeReadCloser *fileSizeReadCloser) Close() error {
    return fileSizeReadCloser.file.Close()
}

