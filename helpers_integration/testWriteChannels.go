package helpers_integration

import (
    "io"
    "os"
    "spectra/ds3_go_sdk/ds3/models"
)

type testRandomAccessWriteChannelBuilder struct {
    name string
}

func (builder *testRandomAccessWriteChannelBuilder) IsChannelAvailable(offset int64) bool {
    return true
}

func (builder *testRandomAccessWriteChannelBuilder) GetChannel(offset int64) (io.WriteCloser, error) {
    f, err := os.Open(builder.name)
    if err != nil {
        return nil, err
    }
    f.Seek(offset, io.SeekStart)
    return f, nil
}

func (builder *testRandomAccessWriteChannelBuilder) OnDone(writer io.WriteCloser) {
    writer.Close()
}



type testPartialObjectRandomAccessWriteChannelBuilder struct {
    name string
    ranges []models.Range //todo
}

func (builder *testPartialObjectRandomAccessWriteChannelBuilder) IsChannelAvailable(offset int64) bool {
    return true
}

func (builder *testPartialObjectRandomAccessWriteChannelBuilder) GetChannel(offset int64) (io.WriteCloser, error) {
    f, err := os.Open(builder.name)
    if err != nil {
        return nil, err
    }
    f.Seek(offset, io.SeekStart)
    return f, nil
}

func (builder *testPartialObjectRandomAccessWriteChannelBuilder) OnDone(writer io.WriteCloser) {
    writer.Close()
}