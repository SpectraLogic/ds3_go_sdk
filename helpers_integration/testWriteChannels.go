package helpers_integration

import (
    "io"
    "os"
)

type testRandomAccessWriteChannelBuilder struct {
    name string
}

func (builder *testRandomAccessWriteChannelBuilder) IsChannelAvailable(offset int64) bool {
    return true
}

func (builder *testRandomAccessWriteChannelBuilder) GetChannel(offset int64) (io.WriteCloser, error) {
    f, err := os.OpenFile(builder.name, os.O_WRONLY, 0)
    if err != nil {
        return nil, err
    }
    f.Seek(offset, io.SeekStart)
    return f, nil
}

func (builder *testRandomAccessWriteChannelBuilder) OnDone(writer io.WriteCloser) {
    writer.Close()
}
