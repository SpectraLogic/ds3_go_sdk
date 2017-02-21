package commands

import (
    "errors"
    "ds3"
    "ds3/models"
)

func putBucket(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket name when doing put_bucket.")
    }

    // Run request.
    _, err := client.PutBucket(models.NewPutBucketRequest(args.Bucket))
    return err
}

