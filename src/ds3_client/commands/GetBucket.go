package commands

import (
    "fmt"
    "errors"
    "ds3"
    "ds3/models"
)

const maxInt = int(^uint(0) >> 1)
const defaultMaxKeys = 1000

func getMinInt(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func getBucket(client *ds3.Client, args *Arguments) error {
    objects, err := getBucketObjects(client, args)
    if err != nil {
        return err
    }
    for _, obj := range objects {
        //TODO: better result printing
        fmt.Println(obj.Key)
    }
    return nil
}

func getBucketObjects(client *ds3.Client, args *Arguments) ([]models.Object, error) {
    // Validate arguments.
    if args.Bucket == "" {
        return nil, errors.New("Must specify a bucket name when doing get_bucket.")
    }

    // Figure out how many keys to restrict to.
    remainingKeys := maxInt
    if args.MaxKeys > 0 {
        remainingKeys = args.MaxKeys
    }

    // Do a do...while pattern to do as many requests as needed.
    var results []models.Object
    marker := ""
    for {
        // Build the request.
        request := buildRequest(args, remainingKeys, marker)

        // Send the request.
        response, err := client.GetBucket(request)
        if err != nil {
            return nil, err
        }

        // Output the results.
        for _, obj := range(response.Contents) {
            results = append(results, obj)
        }

        // Subtract the number of keys that we got from the number of keys that
        // we need to get.
        remainingKeys -= len(response.Contents)

        // Take note of the next marker to get.
        marker = response.NextMarker

        // Take care of the do...while.
        if response.IsTruncated == false || remainingKeys <= 0 {
            break
        }
    }

    return results, nil
}

func buildRequest(args *Arguments, remainingKeys int, marker string) *models.GetBucketRequest {
    request := models.NewGetBucketRequest(args.Bucket)
    request.WithMaxKeys(getMinInt(remainingKeys, defaultMaxKeys))
    if args.KeyPrefix != "" {
        request.WithPrefix(args.KeyPrefix)
    }
    if marker != "" {
        request.WithMarker(marker)
    }
    return request
}

