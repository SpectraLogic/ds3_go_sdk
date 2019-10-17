package helpers

import "time"

// Strategy for how to blob objects, used both in writing and reading blob strategies
type BlobStrategy interface {
    // Determines the maximum amount to delay between calls to getAvailableJobChunk.
    // If blobs are finishing processing, then we will query for more ready blobs earlier.
    // The recommended duration is five minutes.
    delay() time.Duration

    // determines the maximum number of go routines to be created when transferring objects to/from BP
    maxConcurrentTransfers() uint

    // determines the maximum number of blobs that are waiting to be transferred in the operations queue
    // i.e. the size of the operation queue
    maxWaitingTransfers() uint
}

type SimpleBlobStrategy struct {
    Delay                  time.Duration
    MaxConcurrentTransfers uint
    MaxWaitingTransfers    uint
}

func (strategy *SimpleBlobStrategy) delay() time.Duration {
    return strategy.Delay
}

func (strategy *SimpleBlobStrategy) maxConcurrentTransfers() uint {
    return strategy.MaxConcurrentTransfers
}

func (strategy *SimpleBlobStrategy) maxWaitingTransfers() uint {
    return strategy.MaxWaitingTransfers
}