package commands

import "os"

// Implements the SizedReadCloser interface for a file.
// Decorates an os.File and adds a Size function which determines the length of the file
type fileWithSizeDecorator struct {
    file *os.File
}

func readFile(path string) (*fileWithSizeDecorator, error) {
    content, fileErr := os.Open(path)
    if fileErr != nil {
        return nil, fileErr
    }
    return &fileWithSizeDecorator{content}, nil
}

func (fileWithSizeDecorator *fileWithSizeDecorator) Size() (int64, error) {
    fi, err := fileWithSizeDecorator.file.Stat()
    if err != nil {
        return 0, err
    }
    return fi.Size(), nil
}

func (fileWithSizeDecorator *fileWithSizeDecorator) Read(p []byte) (int, error) {
    return fileWithSizeDecorator.file.Read(p)
}

func (fileWithSizeDecorator *fileWithSizeDecorator) Seek(offset int64, whence int) (int64, error) {
    return fileWithSizeDecorator.file.Seek(offset, whence)
}

func (fileWithSizeDecorator *fileWithSizeDecorator) Close() error {
    return fileWithSizeDecorator.file.Close()
}

