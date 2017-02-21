package commands

import (
    "errors"
    "ds3"
    "ds3/models"
)

func deleteBucket(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket name when doing delete_bucket.")
    }

    // Run request.
    _, err := client.DeleteBucket(models.NewDeleteBucketRequest(args.Bucket))
    return err
}


