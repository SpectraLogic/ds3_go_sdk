package channels

import (
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "io"
    "os"
)

type PartialObjectWriteChannelBuilder struct {
    name string
    // map of object ranges to destination offset for start of range
    rangeMap partialObjectRangeMap
    FatalErrorHandler
}

func NewPartialObjectChannelBuilder(name string, objRanges []ds3Models.Range) helperModels.WriteChannelBuilder {

    return &PartialObjectWriteChannelBuilder{
        name: name,
        rangeMap: newPartialObjectRangeMap(objRanges),
    }
}

func (builder *PartialObjectWriteChannelBuilder) IsChannelAvailable(objectOffset int64) bool {
    return true
}

func (builder *PartialObjectWriteChannelBuilder) GetChannel(objectOffset int64) (io.WriteCloser, error) {
    destinationOffset, err := builder.rangeMap.getDestinationOffset(objectOffset)
    if err != nil {
        return nil, err
    }

    f, err := os.OpenFile(builder.name, os.O_WRONLY, defaultPermissions)
    if err != nil {
        return nil, err
    }

    _, err = f.Seek(destinationOffset, io.SeekStart)
    if err != nil {
        f.Close()
        return nil, err
    }
    return f, nil
}

func (builder *PartialObjectWriteChannelBuilder) OnDone(writer io.WriteCloser) {
    writer.Close()
}
