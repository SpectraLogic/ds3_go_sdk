package models

import "io"

type ReadChannelBuilder interface {
    IsChannelAvailable(offset int64) bool
    GetChannel(offset int64) (io.ReadCloser, error)
    OnDone(reader io.ReadCloser) // Determines what a given blob does when it finishes transferring
}

type WriteChannelBuilder interface {
    IsChannelAvailable(offset int64) bool
    GetChannel(offset int64) (io.WriteCloser, error)
    OnDone(writer io.WriteCloser) // Determines what a given blob does when it finishes transferring
}

