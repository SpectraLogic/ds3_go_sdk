package helpers_integration

import (
    "errors"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/helpers/channels"
    "io"
    "log"
    "os"
)

type testStreamAccessWriteChannelBuilder struct {
    f *os.File
    channelCounter channelCounter
    channels.FatalErrorHandler
}

func (builder *testStreamAccessWriteChannelBuilder) IsChannelAvailable(offset int64) bool {
    curOffset, _ := builder.f.Seek(0, io.SeekCurrent)
    if curOffset != offset || builder.channelCounter.currentCount() > 0 {
        log.Printf("DEBUG: channel is not available. Offset expected Offset expected=%d actual=%d. Semaphore expected=0 actual=%d\n", offset, curOffset, builder.channelCounter.currentCount())
        return false
    }
    log.Printf("DEBUG: channel is available\n")
    return true
}

func (builder *testStreamAccessWriteChannelBuilder) GetChannel(offset int64) (io.WriteCloser, error) {
    if builder.channelCounter.currentCount() > 0 {
        return nil, errors.New("ERROR: channel is not currently available")
    }
    builder.channelCounter.increment()
    return builder.f, nil
}

func (builder *testStreamAccessWriteChannelBuilder) OnDone(writer io.WriteCloser) {
    if builder.channelCounter.currentCount() == 0 {
        log.Printf("ERROR: cannot perform OnDone as no channels currently allocated")
        return
    }
    builder.channelCounter.decrement()
}

// test channel that will fail to get channel
type testStreamAccessWriteFailOnFirstBlob struct {
    channelCounter
    channels.FatalErrorHandler
}

func (builder *testStreamAccessWriteFailOnFirstBlob) IsChannelAvailable(offset int64) bool {
    if offset != 0 {
        log.Printf("DEBUG: channel is not available at offset=%d. Semaphore expected=0 actual=%d\n", offset, builder.channelCounter.currentCount())
        return false
    }
    log.Printf("DEBUG: channel is available\n")
    return true
}

func (builder testStreamAccessWriteFailOnFirstBlob) GetChannel(offset int64) (io.WriteCloser, error) {
    return nil, fmt.Errorf("unable to get channel")
}

func (builder *testStreamAccessWriteFailOnFirstBlob) OnDone(writer io.WriteCloser) {
    if builder.channelCounter.currentCount() == 0 {
        log.Printf("ERROR: cannot perform OnDone as no channels currently allocated")
        return
    }
    builder.channelCounter.decrement()
}
