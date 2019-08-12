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

// Captures the first fatal error that occurs.
// In a multi-threaded environment, all subsequent errors will be ignored.
func (errHandler *FatalErrorHandler) SetFatalError(err error) {
    errHandler.errLock.Lock()
    defer errHandler.errLock.Unlock()

    if errHandler.fatalError == nil {
        errHandler.fatalError = err
    }
}

func (errHandler *FatalErrorHandler) GetFatalError() error {
    errHandler.errLock.RLock()
    defer errHandler.errLock.RUnlock()

    return errHandler.fatalError
}
