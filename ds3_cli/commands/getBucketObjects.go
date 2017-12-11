package commands

import (
    "errors"
    "spectra/ds3_go_sdk/ds3"
    "spectra/ds3_go_sdk/ds3/models"
)

// Gets all objects in a bucket, performing multiple requests if the results
// are paged. Also supports limiting to an arbitrary number of keys independent
// of page size.
func getBucketObjects(client *ds3.Client, args *Arguments) ([]models.Ds3PutObject, error) {
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
    var results []models.Ds3PutObject
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
        for _, obj := range response.ListBucketResult.Objects {
            ds3object := models.Ds3PutObject{Name:*obj.Key, Size:obj.Size}
            results = append(results, ds3object)
        }

        // Subtract the number of keys that we got from the number of keys that
        // we need to get.
        remainingKeys -= len(response.ListBucketResult.Objects)

        // Take note of the next marker to get.
        marker = *response.ListBucketResult.NextMarker

        // Take care of the do...while.
        if response.ListBucketResult.Truncated == false || remainingKeys <= 0 {
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

