package helpers

import (
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
)

type HelperInterface interface {
    //ListObjectsFromBucket(bucketName string) []ds3Models.S3Object
    //ListObjectsFromDirectory(directoryName string) []helperModels.PutObject
    PutObjects(bucketName string, objects []helperModels.PutObject, strategy WriteTransferStrategy) (error)
    GetObjects(bucketName string, objects []helperModels.GetObject, strategy ReadTransferStrategy) (error)
}

type HelperImpl struct {
     client *ds3.Client
}

func NewHelpers(client *ds3.Client) HelperInterface {
    return &HelperImpl{client:client}
}

/*
func (helper *HelperImpl) ListObjectsFromBucket(bucketName string) []ds3Models.S3Object {
    panic("not implemented yet")
}

func (helper *HelperImpl) ListObjectsFromDirectory(directoryName string) []helperModels.PutObject {
    panic("not implemented yet")
}
*/

func (helper *HelperImpl) PutObjects(bucketName string, objects []helperModels.PutObject, strategy WriteTransferStrategy) (error) {
    transceiver := newPutTransceiver(bucketName, &objects, &strategy, helper.client)
    return transceiver.transfer()
}

func (helper *HelperImpl) GetObjects(bucketName string, objects []helperModels.GetObject, strategy ReadTransferStrategy) (error) {
    transceiver := newGetTransceiver(bucketName, &objects, &strategy, helper.client)
    return transceiver.transfer()
}
