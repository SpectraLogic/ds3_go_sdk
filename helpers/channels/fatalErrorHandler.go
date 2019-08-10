package channels

import "sync"

// Implements the fatal error handling in channel builders.
// This is used to track if a file has encountered a fatal error such that
// future blobs for a given file should not be processed.
type FatalErrorHandler struct {
    errLock sync.RWMutex
    fatalError error
}

func (errHandler *FatalErrorHandler) HasFatalError() bool {
    errHandler.errLock.RLock()
    defer errHandler.errLock.RUnlock()

    return errHandler.fatalError != nil
}

func (errHandler *FatalErrorHandler) SetFatalError(err error) {
    errHandler.errLock.Lock()
    defer errHandler.errLock.Unlock()

    errHandler.fatalError = err
}

func (errHandler *FatalErrorHandler) GetFatalError() error {
    errHandler.errLock.RLock()
    defer errHandler.errLock.RUnlock()

    return errHandler.fatalError
}