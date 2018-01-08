package helpers

import (
    "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/ds3"
    "io"
)

type WriteObject struct {
    PutObject      models.Ds3PutObject
    ChannelBuilder ChannelBuilder
}

type ReadObject struct {
    GetObject      models.Ds3GetObject
    ChannelBuilder ChannelBuilder
}

type ChannelBuilder interface {
    IsChannelAvailable(offset int64) bool
    GetChannel(offset int64) (reader io.ReadCloser, err error)
    OnDone() // Determines what a given blob does when it finishes transferring
}

type HelperInterface interface {
    ListObjectsFromBucket(bucketName string) []models.S3Object
    ListObjectsFromDirectory(directoryName string) []WriteObject
    WriteObjects(objects []WriteObject, bucketName string, strategy WriteTransferStrategy) (error) //todo return future
    ReadObjects(bucketName string, objects[]ReadObject, strategy ReadTransferStrategy) (error)     //todo return future
}

type HelperImpl struct {
     client *ds3.Client
}

func NewHelpers(client *ds3.Client) HelperInterface {
    return &HelperImpl{client:client}
}

func (h *HelperImpl) ListObjectsFromBucket(bucketName string) []models.S3Object {
    //todo
    panic("not implemented yet")
}

func (h *HelperImpl) ListObjectsFromDirectory(directoryName string) []WriteObject {
    //todo
    panic("not implemented yet")
}

func (h *HelperImpl) WriteObjects(writeObjects []WriteObject, bucketName string, strategy WriteTransferStrategy) (error) {
    transfernator := newPutTransfernator(bucketName, &writeObjects, &strategy, h.client)
    err := transfernator.transfer()
    return err
}

func (h *HelperImpl) ReadObjects(bucketName string, objects[]ReadObject, strategy ReadTransferStrategy) (error) {
    //todo
    panic("not implemented yet")
}