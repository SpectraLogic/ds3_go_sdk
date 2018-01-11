package helpers

import (
    "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/ds3"
    "io"
)

type PutObject struct {
    PutObject      models.Ds3PutObject
    ChannelBuilder ReadChannelBuilder
}

type GetObject struct {
    GetObject      models.Ds3GetObject
    ChannelBuilder WriteChannelBuilder
}

type ReadChannelBuilder interface {
    IsChannelAvailable(offset int64) bool
    GetChannel(offset int64) (reader io.ReadCloser, err error)
    OnDone() // Determines what a given blob does when it finishes transferring
}

type WriteChannelBuilder interface {
    IsChannelAvailable(offset int64) bool
    GetChannel(offset int64) (reader io.WriteCloser, err error)
    OnDone() // Determines what a given blob does when it finishes transferring
}

type HelperInterface interface {
    ListObjectsFromBucket(bucketName string) []models.S3Object
    ListObjectsFromDirectory(directoryName string) []PutObject
    PutObjects(objects []PutObject, bucketName string, strategy WriteTransferStrategy) (error) //todo return future
    GetObjects(bucketName string, objects []GetObject, strategy ReadTransferStrategy) (error)   //todo return future
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

func (h *HelperImpl) ListObjectsFromDirectory(directoryName string) []PutObject {
    //todo
    panic("not implemented yet")
}

func (h *HelperImpl) PutObjects(objects []PutObject, bucketName string, strategy WriteTransferStrategy) (error) {
    transfernator := newPutTransfernator(bucketName, &objects, &strategy, h.client)
    err := transfernator.transfer()
    return err
}

func (h *HelperImpl) GetObjects(bucketName string, objects []GetObject, strategy ReadTransferStrategy) (error) {
    transfernator := newGetTransfernator(bucketName, &objects, &strategy, h.client)
    err := transfernator.transfer()
    return err
}
