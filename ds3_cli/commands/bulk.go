package commands

import "github.com/SpectraLogic/ds3_go_sdk/ds3/models"

type bulkHandler func(object models.BulkObject) error

func handleBulkResponse(objectLists []models.Objects, objectHandler bulkHandler) error {
    // Make a channel to keep track of the errors and how many goroutines have finished.
    finish := make(chan error)

    // Spawn goroutines to perform the actual gets or puts.
    for _, objectList := range objectLists {
        go func() {
            // Loop over each object and call the bulk handler.
            for _, obj := range objectList.Objects {
                // If there was an error, pass it onto the error channel and
                // terminate the goroutine.
                if err := objectHandler(obj); err != nil {
                    finish <- err
                    return
                }
            }

            // Signal that the object list is done.
            finish <- nil
        }()
    }

    // Wait for every goroutine to stop and accumulate their errors.
    errors := &aggregateError{}
    for i := len(objectLists); i > 0; i-- {
        err := <-finish
        if err != nil {
            errors.Add(err)
        }
    }

    return errors
}

