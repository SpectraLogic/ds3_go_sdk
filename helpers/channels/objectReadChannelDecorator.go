package channels

import (
	"io"
	"github.com/SpectraLogic/ds3_go_sdk/helpers/models"
)

type ObjectReadChannelDecorator struct {
	readCloser io.ReadCloser
	FatalErrorHandler
}

func NewObjectReadChannelDecorator(reader io.Reader) models.ReadChannelBuilder {
	return &ObjectReadChannelDecorator { readCloser: NewDs3ObjectReadCloserDecorator(reader) }
}

func (readChannelDecorator *ObjectReadChannelDecorator) IsChannelAvailable(offset int64) bool {
	return true
}

func (readChannelDecorator *ObjectReadChannelDecorator) GetChannel(offset int64) (io.ReadCloser, error) {
	return readChannelDecorator.readCloser, nil
}

func (readChannelDecorator *ObjectReadChannelDecorator) OnDone(reader io.ReadCloser) {
	// Intentionally not implemented
}
