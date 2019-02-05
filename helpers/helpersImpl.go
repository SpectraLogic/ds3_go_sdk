package helpers

import (
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
)

type HelperInterface interface {
    //ListObjectsFromBucket(bucketName string) []ds3Models.S3Object
    //ListObjectsFromDirectory(directoryName string) []helperModels.PutObject

    // Puts the specified list of objects on the Black Pearl in the specified bucket.
    // Returns the Black Pearl bulk put Job ID and any errors that may have occurred.
    // A job ID will be returned if a BP job was successfully created, regardless of
    // whether additional errors occur.
    PutObjects(bucketName string, objects []helperModels.PutObject, strategy WriteTransferStrategy) (string, error)

    // Retrieves the list of objects from the specified bucket on the Black Pearl.
    // Returns the Black Pearl bulk get Job ID and any errors that may have occurred.
    // A job ID will be returned if a BP job was successfully created, regardless of
    // whether additional errors occur.
    GetObjects(bucketName string, objects []helperModels.GetObject, strategy ReadTransferStrategy) (string, error)
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

func (helper *HelperImpl) PutObjects(bucketName string, objects []helperModels.PutObject, strategy WriteTransferStrategy) (string, error) {
    transceiver := newPutTransceiver(bucketName, &objects, &strategy, helper.client)
    return transceiver.transfer()
}

func (helper *HelperImpl) GetObjects(bucketName string, objects []helperModels.GetObject, strategy ReadTransferStrategy) (string, error) {
    transceiver := newGetTransceiver(bucketName, &objects, &strategy, helper.client)
    return transceiver.transfer()
}
