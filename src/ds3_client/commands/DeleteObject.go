package commands

import (
    "errors"
    "ds3"
    "ds3/models"
)

func deleteObject(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket name when doing delete_object.")
    }
    if args.Key == "" {
        return errors.New("Must specify an object key when doing delete_object.")
    }

    // Run request.
    _, err := client.DeleteObject(models.NewDeleteObjectRequest(args.Bucket, args.Key))
    return err
}



