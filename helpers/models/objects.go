package models

import "spectra/ds3_go_sdk/ds3/models"

type GetObject struct {
    Name string
    Ranges []models.Range
    ChannelBuilder WriteChannelBuilder
}

type PutObject struct {
    PutObject      models.Ds3PutObject
    ChannelBuilder ReadChannelBuilder
}