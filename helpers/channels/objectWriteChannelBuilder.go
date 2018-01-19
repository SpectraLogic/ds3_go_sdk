package channels

import (
    "io"
    "os"
    helperModels "spectra/ds3_go_sdk/helpers/models"
)

// Implements the WriteChannelBuilder interface and uses a file as the WriteCloser implementation.
// This channel functions as a random-access and can be used concurrently so long as two writers
// are not writing to the same location within the file.
type ObjectWriteChannelBuilder struct {
    name string
}

func NewWriteChannelBuilder(name string) helperModels.WriteChannelBuilder {
    return &ObjectWriteChannelBuilder{name:name}
}

func (builder *ObjectWriteChannelBuilder) IsChannelAvailable(offset int64) bool {
    return true
}

func (builder *ObjectWriteChannelBuilder) GetChannel(offset int64) (io.WriteCloser, error) {
    f, err := os.OpenFile(builder.name, os.O_WRONLY, 0)
    if err != nil {
        return nil, err
    }
    f.Seek(offset, io.SeekStart)
    return f, nil
}

func (builder *ObjectWriteChannelBuilder) OnDone(writer io.WriteCloser) {
    writer.Close()
}
