package commands

import (
    "errors"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "strconv"
)

func headObject(client *ds3.Client, args *Arguments) error {
    // Validate the arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket.")
    }
    bucket := args.Bucket
    key := args.Key
    if key == "" {
        // do whole bucket
        request:= models.NewGetBucketRequestWithPrefix(bucket, args.KeyPrefix)
        response, err := client.GetBucket(request)
        if err != nil {
            return fmt.Errorf("Failed to get bucket %s\n%v", bucket, err)
        }

        for _, object := range response.ListBucketResult.Objects {
            err = processHeadObject(client, bucket, *object.Key)
            if err != nil {
                return fmt.Errorf("Failed head_object for Bucket %s, key %s\n%v",
                                    bucket, *object.Key, err)
            }
        }
        return nil
    } else {
        // single object
        return processHeadObject(client, bucket, key)
    }
}

func processHeadObject(client *ds3.Client, bucket string, key string) error {
    // Build the request.
    request := models.NewHeadObjectRequest(bucket, key)

    // Perform the request.
    response, requestErr := client.HeadObject(request)
    if requestErr != nil {
        return requestErr
    }

    checksums := response.BlobChecksums
    checksumType := response.BlobChecksumType

    fmt.Printf("Bucket: %s\n", bucket)
    fmt.Printf("Key: %s\n", key)
    fmt.Printf("Type: %s\n", checksumType)
    for i, c := range checksums {
        fmt.Printf("%s: %s\n",strconv.Itoa(int(i)), c)
    }
    fmt.Printf("-----------------------------\n")

    return nil
}

