package commands

import (
    "io"
    "os"
    "path"
    "errors"
    "ds3"
    "ds3/models"
)

func bulkGet(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket name when doing a bulk_get.")
    }

    // Determine the objects that we need to queue a bulk get for.
    objects, err := getBucketObjects(client, args)
    if err != nil {
        return err
    }

    // Run request.
    response, err := client.BulkGet(models.NewBulkGetRequest(args.Bucket, objects))
    if err != nil {
        return err
    }

    // Handle the responses in parallel
    return handleBulkResponse(response.Objects, buildFileGetter(client, args.Bucket))
}

func buildFileGetter(client *ds3.Client, bucketName string) bulkHandler {
    return func(objectList []models.Object, finish chan error) {
        // Get each object.
        for _, obj := range objectList {
            // Perform the request.
            response, requestErr := client.GetObject(models.NewGetObjectRequest(bucketName, obj.Key))

            // Send an error along and stop the function if we couldn't submit the request.
            if requestErr != nil {
                finish <- requestErr
                return
            }
            defer response.Content.Close()

            // Create the directory where we're going to store the file.
            currentDir, dirErr := os.Stat(".")
            if dirErr != nil {
                finish <- dirErr
                return
            }
            mkdirErr := os.MkdirAll(path.Dir(obj.Key), currentDir.Mode())

            // Send an error along and stop the function if we couldn't create the directory.
            if mkdirErr != nil {
                finish <- mkdirErr
                return
            }

            // Open the file to write.
            file, fileErr := os.Create(obj.Key)

            // Send an error along and stop the function if we couldn't open the file.
            if fileErr != nil {
                finish <- fileErr
                return
            }
            defer file.Close()

            // Copy the request stream to the file.
            io.Copy(file, response.Content)
        }

        // Signal that we're done.
        finish <- nil
    }
}

