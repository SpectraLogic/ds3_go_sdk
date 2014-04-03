package commands

import (
    "fmt"
    "errors"
    "math"
    "ds3"
    "ds3/models"
)

func getBucket(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket name when doing get_bucket.")
    }

    remainingKeys := math.MaxInt
    if args.MaxKeys > 0 {
        remainingKeys = args.MaxKeys
    }

    marker := ""
    for {
        // Build the request.
        request := models.NewGetBucketRequest(args.Bucket)
        if args.KeyPrefix != "" {
            request.WithPrefix(args.KeyPrefix)
        }
        if marker != "" {
            request.WithMarker(marker)
        }

        // Send the request.
        response, err := client.GetBucket(request)
        if err != nil {
            return err
        }

        // Output the results.
        for _, obj := range(response.Contents) {
            printObject(obj)
        }

        // Subtract the number of keys that we got from the number of keys that
        // we need to get.
        remainingKeys -= len(response.Contents)

        // Take note of the next marker to get.
        marker = response.NextMarker

        // Take care of the do...while.
        if !response.IsTruncated || remainingKeys <= 0 {
            break
        }
    }

    return err
}

//TODO: better result printing
func printObject(obj models.Object) {
    fmt.Println(object.Key)
}

