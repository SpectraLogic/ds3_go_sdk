package helpers_integration

import (
    "io"
    "bytes"
    "fmt"
    "io/ioutil"
    "errors"
    "spectra/ds3_go_sdk/ds3"
    "spectra/ds3_go_sdk/helpers"
    helperModels "spectra/ds3_go_sdk/helpers/models"
)

const LargeBookPath = "./resources/bigfiles/"

const LargeBookTitle = "lesmis-copies.txt"

// Verifies the content of a ReadCloser matches the content of the specified book.
func VerifyLargeBookContent(content io.ReadCloser) error {
    expected, err := ioutil.ReadFile(LargeBookPath + LargeBookTitle)
    if err != nil {
        return err
    }

    actual, err := ioutil.ReadAll(content)
    if err != nil {
        return err
    }

    if bytes.Compare(expected, actual) != 0 {
        errors.New(fmt.Sprintf("Mismatched content for expected and received object. Expected size %d, actual size %d.",
            len(expected),
            len(actual)))
    }
    return nil
}

func LoadLargeFile(bucket string, client *ds3.Client) error {
    helper := helpers.NewHelpers(client)

    writeObj, err := getTestWriteObjectRandomAccess(LargeBookTitle, LargeBookPath + LargeBookTitle)
    if err != nil {
        return err
    }

    var writeObjects []helperModels.PutObject
    writeObjects = append(writeObjects, *writeObj)

    return helper.PutObjects(bucket, writeObjects, newTestTransferStrategy())
}