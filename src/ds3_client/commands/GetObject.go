package commands

import (
    "os"
    "io"
    "errors"
    "ds3"
    "ds3/models"
)

func getObject(client *ds3.Client, args *Arguments) error {
    // Validate the arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket.")
    }
    if args.Key == "" {
        return errors.New("Must specify an object key.")
    }

    // Make sure the file doesn't already exist.
    if _, err := os.Stat(args.Key); err == nil {
        return errors.New("File already exists")
    }

    // Perform the request.
    response, requestErr := client.GetObject(models.NewGetObjectRequest(args.Bucket, args.Key))
    if requestErr != nil {
        return requestErr
    }
    defer response.Content.Close()

    // Open the file to write.
    file, fileErr := os.Create(args.Key)
    if fileErr != nil {
        return fileErr
    }
    defer file.Close()

    // Copy the request stream to the file.
    io.Copy(file, response.Content)

    return nil
}

