package helpers_integration

import (
    "io"
    "os"
    "time"
    "log"
    "errors"
    "sync"
    "spectra/ds3_go_sdk/helpers"
    "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/ds3_integration/utils"
)

// TODO abstract random access and stream channel and move abstraction to networking layer. Leave test implementation here.

// Channel builder for a test file with random access
type testRandomAccessChannelBuilder struct {
    name string
}

func (cb *testRandomAccessChannelBuilder) GetChannel(offset int64) (io.ReadCloser, error) {
    f, err := os.Open(cb.name)
    if err != nil {
        return nil, err
    }
    f.Seek(offset, io.SeekStart)
    return f, nil
}

func (cb *testRandomAccessChannelBuilder) IsChannelAvailable(offset int64) bool {
    return true
}

func (testRandomAccessChannelBuilder) OnDone() {
    //do nothing
}

// Channel builder for a test file simulating stream access
type testStreamAccessChannelBuilder struct {
    f *os.File
    channelCounter channelCounter
}

// Used as a thread safe counter to keep track of the current number of channels in use for a given ChannelBuilder
type channelCounter struct {
    mux sync.RWMutex
    numChan int
}

func (cc *channelCounter) currentCount() int {
    cc.mux.RLock()
    defer cc.mux.RUnlock()
    return cc.numChan
}

func (cc *channelCounter) increment() {
    cc.mux.Lock()
    defer cc.mux.Unlock()
    cc.numChan++
}

func (cc *channelCounter) decrement() {
    cc.mux.Lock()
    defer cc.mux.Unlock()
    cc.numChan--
}

func (sb *testStreamAccessChannelBuilder) GetChannel(offset int64) (io.ReadCloser, error) {
    if sb.channelCounter.currentCount() > 0 {
        return nil, errors.New("ERROR: channel is not currently available")
    }
    log.Printf("DEBUG: incrementing channel semaphore with length %d", sb.channelCounter.currentCount()) //todo delete
    sb.channelCounter.increment()
    return sb.f, nil
}

func (sb *testStreamAccessChannelBuilder) IsChannelAvailable(offset int64) bool {
    curOffset, _ := sb.f.Seek(0, io.SeekCurrent)
    if curOffset != offset || sb.channelCounter.currentCount() > 0 {
        log.Printf("DEBUG: channel is not available. Offset expected=%d actual=%d. Semaphore expected=0 actual=%d", offset, curOffset, sb.channelCounter.currentCount()) //todo delete
        return false
    }
    log.Print("DEBUG: channel deamed available\n") //todo delete
    return true
}

func (sb *testStreamAccessChannelBuilder) OnDone() {
    if sb.channelCounter.currentCount() == 0 {
        log.Printf("ERROR: cannot perform OnDone as no channels currently allocated")
        return
    }
    log.Printf("DEBUG: Decrementing semaphore with current usage %d", sb.channelCounter.currentCount())
    sb.channelCounter.decrement() // decrement semaphore
}

func getTestWriteObjectStreamAccess(objectName string, path string) (*helpers.WriteObject, error) {
    fileInfo, err := os.Stat(path)
    if err != nil {
        return nil, err
    }
    size := fileInfo.Size()
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    curWriteObj := helpers.WriteObject{
        PutObject:models.Ds3PutObject{Name:objectName,Size:size},
        ChannelBuilder:&testStreamAccessChannelBuilder{f:f},
    }
    return &curWriteObj, nil
}

// Retrieves a file with the specified path, and creates a WriteObject using the specified name
func getTestWriteObjectRandomAccess(objectName string, path string) (*helpers.WriteObject, error) {
    channelBuilder := testRandomAccessChannelBuilder{name:path}
    fileInfo, err := os.Stat(path)
    if err != nil {
        return nil, err
    }
    size := fileInfo.Size()
    curWriteObj := helpers.WriteObject{
        PutObject:models.Ds3PutObject{Name:objectName,Size:size},
        ChannelBuilder:&channelBuilder,
    }
    return &curWriteObj, nil
}

// Retrieves the test books as write objects with random access
func getTestBooksAsWriteObjects() (*[]helpers.WriteObject, error) {
    var writeObjects []helpers.WriteObject
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
func newTestTransferStrategy() helpers.WriteTransferStrategy {
    return helpers.WriteTransferStrategy{
        BlobStrategy: newTestBlobStrategy(),
        Options:      helpers.WriteBulkJobOptions{MaxUploadSize:&helpers.MinUploadSize},
    }
}

// Creates a simple blob strategy for testing
func newTestBlobStrategy() *helpers.SimpleWriteBlobStrategy {
    var delay time.Duration = time.Second * 5
    var maxTransferGoroutines = 5
    return &helpers.SimpleWriteBlobStrategy{Delay:delay, MaxTransferGoroutines:maxTransferGoroutines}
}
