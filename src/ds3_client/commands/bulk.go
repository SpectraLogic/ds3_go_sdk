package commands

import "ds3/models"

type bulkHandler func([]models.Object, chan error)

func handleBulkResponse(objectLists [][]models.Object, objectHandler bulkHandler) error {
    // Make a channel to keep track of the errors and how many threads have finished.
    finish := make(chan error)

    // Spawn put threads.
    for _, objectList := range objectLists {
        go objectHandler(objectList, finish)
    }

    // Wait for every thread to stop and accumulate their errors.
    errors := &aggregateError{}
    for i := len(objectLists); i > 0; i-- {
        err := <-finish
        if err != nil {
            errors.Add(err)
        }
    }

    return errors
}

