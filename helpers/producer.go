package helpers

import "spectra/ds3_go_sdk/ds3/models"

type Producer interface {
    run(error *models.AggregateError)
}
