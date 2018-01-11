package helpers

import "time"

// Strategy for how to blob objects, used both in writing and reading blob strategies
type BlobStrategy interface {
    delay() // performs delay when no job chunks are available
    maxTransferGoroutines() int // determines the maximum number of go routines to be created when transferring objects to/from BP
}

type SimpleBlobStrategy struct {
    Delay time.Duration
    MaxTransferGoroutines int
}

func (sbs *SimpleBlobStrategy) delay() {
    time.Sleep(sbs.Delay)
}

func (sbs *SimpleBlobStrategy) maxTransferGoroutines() int {
    return sbs.MaxTransferGoroutines
}
