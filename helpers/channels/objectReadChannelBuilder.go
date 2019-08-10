package channels

import (
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "io"
    "os"
)

const defaultPermissions = 0

// Implements the ReadChannelBuilder interface and uses a file as the ReadCloser implementation.
// This channel functions as a random-access and can be used concurrently.
type ObjectReadChannelBuilder struct {
    name string
    FatalErrorHandler
}

func NewReadChannelBuilder(name string) helperModels.ReadChannelBuilder {
    return &ObjectReadChannelBuilder{name:name}
}

func (builder *ObjectReadChannelBuilder) GetChannel(offset int64) (io.ReadCloser, error) {
    f, err := os.OpenFile(builder.name, os.O_RDONLY, defaultPermissions)
    if err != nil {
        return nil, err
    }
    f.Seek(offset, io.SeekStart)
    return f, nil
}

func (builder *ObjectReadChannelBuilder) IsChannelAvailable(offset int64) bool {
    return true
}

func (ObjectReadChannelBuilder) OnDone(reader io.ReadCloser) {
    reader.Close()
}
