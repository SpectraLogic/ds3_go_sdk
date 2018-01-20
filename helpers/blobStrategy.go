package helpers

import "time"

// Strategy for how to blob objects, used both in writing and reading blob strategies
type BlobStrategy interface {
    delay()                      // performs delay when no job chunks are available
    maxConcurrentTransfers() int // determines the maximum number of go routines to be created when transferring objects to/from BP
}

type SimpleBlobStrategy struct {
    Delay                  time.Duration
    MaxConcurrentTransfers int
}

func (strategy *SimpleBlobStrategy) delay() {
    time.Sleep(strategy.Delay)
}

func (strategy *SimpleBlobStrategy) maxConcurrentTransfers() int {
    return strategy.MaxConcurrentTransfers
}
