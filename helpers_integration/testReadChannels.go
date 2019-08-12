package helpers_integration

import (
    "github.com/SpectraLogic/ds3_go_sdk/helpers/channels"
    "sync"
    "io"
    "log"
    "os"
    "errors"
)

// Channel builder for a test file simulating stream access
type testStreamAccessReadChannelBuilder struct {
    f *os.File
    channelCounter channelCounter
    channels.FatalErrorHandler
}

// Used as a thread safe counter to keep track of the current number of channels in use for a given ReadChannelBuilder
type channelCounter struct {
    mux sync.RWMutex
    numChan int
}

func (counter *channelCounter) currentCount() int {
    counter.mux.RLock()
    defer counter.mux.RUnlock()
    return counter.numChan
}

func (counter *channelCounter) increment() {
    counter.mux.Lock()
    defer counter.mux.Unlock()
    counter.numChan++
}

func (counter *channelCounter) decrement() {
    counter.mux.Lock()
    defer counter.mux.Unlock()
    counter.numChan--
}

func (builder *testStreamAccessReadChannelBuilder) GetChannel(offset int64) (io.ReadCloser, error) {
    if builder.channelCounter.currentCount() > 0 {
        return nil, errors.New("ERROR: channel is not currently available")
    }
    log.Printf("DEBUG: incrementing channel semaphore with length %d", builder.channelCounter.currentCount())
    builder.channelCounter.increment()
    return builder.f, nil
}

func (builder *testStreamAccessReadChannelBuilder) IsChannelAvailable(offset int64) bool {
    curOffset, _ := builder.f.Seek(0, io.SeekCurrent)
    if curOffset != offset || builder.channelCounter.currentCount() > 0 {
        log.Printf("DEBUG: channel is not available. Offset expected=%d actual=%d. Semaphore expected=0 actual=%d", offset, curOffset, builder.channelCounter.currentCount())
        return false
    }
    log.Print("DEBUG: channel deamed available\n")
    return true
}

func (builder *testStreamAccessReadChannelBuilder) OnDone(reader io.ReadCloser) {
    if builder.channelCounter.currentCount() == 0 {
        log.Printf("ERROR: cannot perform OnDone as no channels currently allocated")
        return
    }
    log.Printf("DEBUG: Decrementing semaphore with current usage %d", builder.channelCounter.currentCount())
    builder.channelCounter.decrement() // decrement semaphore
}