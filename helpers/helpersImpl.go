package helpers

import (
    ds3Models "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/ds3"
    helperModels "spectra/ds3_go_sdk/helpers/models"
)

type HelperInterface interface {
    ListObjectsFromBucket(bucketName string) []ds3Models.S3Object
    ListObjectsFromDirectory(directoryName string) []helperModels.PutObject
    PutObjects(bucketName string, objects []helperModels.PutObject, strategy WriteTransferStrategy) (error) //todo return future
    GetObjects(bucketName string, objects []helperModels.GetObject, strategy ReadTransferStrategy) (error)   //todo return future
}

type HelperImpl struct {
     client *ds3.Client
}

func NewHelpers(client *ds3.Client) HelperInterface {
    return &HelperImpl{client:client}
}

func (helper *HelperImpl) ListObjectsFromBucket(bucketName string) []ds3Models.S3Object {
    //todo implement
    panic("not implemented yet")
}

func (helper *HelperImpl) ListObjectsFromDirectory(directoryName string) []helperModels.PutObject {
    //todo implement
    panic("not implemented yet")
}

func (helper *HelperImpl) PutObjects(bucketName string, objects []helperModels.PutObject, strategy WriteTransferStrategy) (error) {
    transfernator := newPutTransfernator(bucketName, &objects, &strategy, helper.client)
    err := transfernator.transfer()
    return err
}

func (helper *HelperImpl) GetObjects(bucketName string, objects []helperModels.GetObject, strategy ReadTransferStrategy) (error) {
    transfernator := newGetTransfernator(bucketName, &objects, &strategy, helper.client)
    err := transfernator.transfer()
    return err
}
