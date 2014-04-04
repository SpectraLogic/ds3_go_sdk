package commands

import (
    "os"
    "errors"
    "ds3"
    "ds3/models"
)

func putObject(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket name when doing put_object.")
    }
    if args.Key == "" {
        return errors.New("Must specify an object key when doing put_object.")
    }

    // Open the file.
    content, fileErr := os.Open(args.Key)
    if fileErr != nil {
        return fileErr
    }
    sizedContent := &fileSizeReadCloser{content}

    // Run request.
    _, err := client.PutObject(models.NewPutObjectRequest(args.Bucket, args.Key, sizedContent))
    return err
}

type fileSizeReadCloser struct {
    file *os.File
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

