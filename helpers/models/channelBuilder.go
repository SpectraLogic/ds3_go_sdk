package models

import "io"

type ChannelErrorHandler interface {
    HasFatalError() bool // Determines if a fatal error has happened while transferring a file, i.e. don't transfer more blobs for this file.
    SetFatalError(err error) // Sets a fatal error for the associated file. This will stop future blobs from being sent from this channel.
}

type ReadChannelBuilder interface {
    IsChannelAvailable(offset int64) bool
    GetChannel(offset int64) (io.ReadCloser, error)
    OnDone(reader io.ReadCloser) // Determines what a given blob does when it finishes transferring
    ChannelErrorHandler
}

type WriteChannelBuilder interface {
    IsChannelAvailable(offset int64) bool
    GetChannel(offset int64) (io.WriteCloser, error)
    OnDone(writer io.WriteCloser) // Determines what a given blob does when it finishes transferring
    ChannelErrorHandler
}

