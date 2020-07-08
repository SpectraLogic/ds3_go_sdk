package helpers_integration

import (
    "os"
    "testing"
    "time"
    "github.com/SpectraLogic/ds3_go_sdk/helpers"
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_integration/utils"
    "github.com/SpectraLogic/ds3_go_sdk/helpers/channels"
)

func getTestWriteObjectStreamAccess(objectName string, path string) (*helperModels.PutObject, error) {
    fileInfo, err := os.Stat(path)
    if err != nil {
        return nil, err
    }
    size := fileInfo.Size()
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    curWriteObj := helperModels.PutObject{
        PutObject:      ds3Models.Ds3PutObject{Name:objectName,Size:size},
        ChannelBuilder: &testStreamAccessReadChannelBuilder{f:f},
    }
    return &curWriteObj, nil
}

// Retrieves a file with the specified path, and creates a PutObject using the specified name
func getTestWriteObjectRandomAccess(objectName string, path string) (*helperModels.PutObject, error) {
    channelBuilder := channels.NewReadChannelBuilder(path)
    fileInfo, err := os.Stat(path)
    if err != nil {
        return nil, err
    }
    size := fileInfo.Size()
    curWriteObj := helperModels.PutObject{
        PutObject:      ds3Models.Ds3PutObject{Name:objectName,Size:size},
        ChannelBuilder: channelBuilder,
    }
    return &curWriteObj, nil
}

// Retrieves the test books as write objects with random access
func getTestBooksAsWriteObjects() (*[]helperModels.PutObject, error) {
    var writeObjects []helperModels.PutObject
    for _, book := range testutils.BookTitles {
        curWriteObj, err := getTestWriteObjectRandomAccess(book, testutils.BookPath+book)
        if err != nil {
            return nil, err
        }

        writeObjects = append(writeObjects, *curWriteObj)
    }
    return &writeObjects, nil
}

// Creates a simple transfer strategy for testing
func newTestTransferStrategy(t * testing.T) helpers.WriteTransferStrategy {
    return helpers.WriteTransferStrategy{
        BlobStrategy: newTestBlobStrategy(),
        Options:      helpers.WriteBulkJobOptions{MaxUploadSize: &helpers.MinUploadSize},
        Listeners:    newErrorOnErrorListenerStrategy(t),
    }
}

func newErrorOnErrorListenerStrategy(t *testing.T) helpers.ListenerStrategy {
    return helpers.ListenerStrategy{
        ErrorCallback: func(objectName string, err error) {
            t.Errorf("unexpected error on '%s': %v", objectName, err)
        },
    }
}

// Creates a simple blob strategy for testing
func newTestBlobStrategy() *helpers.SimpleBlobStrategy {
    var delay time.Duration = time.Second * 5
    var maxTransferGoroutines uint = 5
    return &helpers.SimpleBlobStrategy{Delay:delay, MaxConcurrentTransfers:maxTransferGoroutines}
}
