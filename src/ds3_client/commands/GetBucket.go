package commands

import (
    "fmt"
    "errors"
    "ds3"
    "ds3/models"
)

func getBucket(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket name when doing get_bucket.")
    }

    //TODO: do the loop that takes care of the chaining multiple requests as well.
    // Run the request.
    request := models.NewGetBucketRequest(args.Bucket)
    if args.KeyPrefix != "" {
        request.WithPrefix(args.KeyPrefix)
    }
    getBucketResponse, err := client.GetBucket(request)
    if err == nil {
        //TODO: better result printing
        for _, object := range(getBucketResponse.Contents) {
            fmt.Println(object.Key)
        }
        //fmt.Println(getBucketResponse)
    }
    return err
}

