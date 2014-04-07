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

func (self fileSizeReadCloser) Size() int64 {
    fi, err := self.file.Stat()
    if err != nil {
        panic(err)
    }
    return fi.Size()
}

func (self fileSizeReadCloser) Read(p []byte) (int, error) {
    return self.file.Read(p)
}

func (self fileSizeReadCloser) Close() error {
    return self.file.Close()
}

