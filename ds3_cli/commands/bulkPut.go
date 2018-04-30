package commands

import (
    "errors"
    "io/ioutil"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
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
    objects := make([]models.Ds3PutObject, len(files))
    for i, file := range files {
        objects[i] = models.Ds3PutObject{Name: file.path, Size: file.size}
    }

    // Run request.
    response, err := client.PutBulkJobSpectraS3(models.NewPutBulkJobSpectraS3Request(args.Bucket, objects))
    if err != nil {
        return err
    }

    // Handle the responses in parallel
    return handleBulkResponse(response.MasterObjectList.Objects, buildFilePutter(client, args.Bucket))
}

// Returns a function that puts an object for a bulk put.
func buildFilePutter(client *ds3.Client, bucketName string) bulkHandler {
    return func(obj models.BulkObject) error {
        // Read the file.
        sizeReadCloser, fileErr := readFile(*obj.Name)
        if fileErr != nil {
            return fileErr
        }

        // Submit the put object request.
        _, putErr := client.PutObject(models.NewPutObjectRequest(
            bucketName,
            *obj.Name,
            sizeReadCloser,
        ))
        return putErr
    }
}

// Represent a file that we found in a recursive directory traverse. We need
// the relative path as well as the size. The FileInfo structure has the size
// and the filename, but not the relative path.
type file struct {
    path string
    size int64
}

func recursivelyListFiles(path string) ([]file, error) {
    // Read the specified directory.
    infos, err := ioutil.ReadDir(path)
    if err != nil {
        return nil, err
    }

    // For each file or directory...
    var results []file
    for _, info := range infos {
        // Build its relative path.
        name := path + "/" + info.Name()

        // If it's a directory...
        if info.IsDir() {
            // List all of the files in there too.
            files, err := recursivelyListFiles(name)
            if err != nil {
                return nil, err
            }

            // Append the results.
            results = append(results, files...)
        } else {
            // Append the current file.
            results = append(results, file{name, info.Size()})
        }
    }
    return results, nil
}

