package commands

import (
    "errors"
    "spectra/ds3_go_sdk/ds3"
    "spectra/ds3_go_sdk/ds3/models"
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
    sizedContent, fileErr := readFile(args.Key)
    if fileErr != nil {
        return fileErr
    }

    // Run request.
    _, err := client.PutObject(models.NewPutObjectRequest(args.Bucket, args.Key, sizedContent))
    return err
}

