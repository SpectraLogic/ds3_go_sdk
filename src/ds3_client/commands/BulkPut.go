package commands

import (
    "errors"
    "io/ioutil"
    "ds3"
    "ds3/models"
)

func bulkPut(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket name when doing a bulk_put.")
    }
    if args.KeyPrefix == "" {
        return errors.New("Must specify a key prefix when doing a bulk_put.")
    }

    // List the files.
    files, filesErr := recursivelyListFiles(args.KeyPrefix)
    if filesErr != nil {
        return filesErr
    }

    // Create object entities that we can query with for each file.
    objects := make([]models.Object, len(files))
    for i, file := range(files) {
        objects[i] = models.Object{Key: file.path, Size: file.size}
    }

    // Run request.
    response, err := client.BulkPut(models.NewBulkPutRequest(args.Bucket, objects))
    if err != nil {
        return err
    }

    // Handle the responses in parallel
    return handleBulkResponse(response.Objects, buildFilePutter(client, args.Bucket))
}

func buildFilePutter(client *ds3.Client, bucketName string) bulkHandler {
    return func(objectList []models.Object, finish chan error) {
        // Put each object.
        for _, obj := range objectList {
            // Read the file.
            sizeReadCloser, fileErr := readFile(obj.Key)

            // Send an error along and stop the function if we couldn't open the file.
            if fileErr != nil {
                finish <- fileErr
                return
            }

            // Submit the put object request.
            _, putErr := client.PutObject(models.NewPutObjectRequest(
                bucketName,
                obj.Key,
                sizeReadCloser,
            ))

            // Send an error along and stop the function if we couldn't submit the request.
            if putErr != nil {
                finish <- putErr
                return
            }
        }

        // Signal that we're done.
        finish <- nil
    }
}

type file struct {
    path string
    size int64
}

func recursivelyListFiles(path string) ([]file, error) {
    infos, err := ioutil.ReadDir(path)
    if err != nil {
        return nil, err
    }

    var results []file
    for _, info := range(infos) {
        name := path + "/" + info.Name()
        if info.IsDir() {
            files, err := recursivelyListFiles(name)
            if err != nil {
                return nil, err
            }
            results = append(results, files...)
        } else {
            results = append(results, file{name, info.Size()})
        }
    }
    return results, nil
}

