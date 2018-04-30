package commands

import (
    "errors"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
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


