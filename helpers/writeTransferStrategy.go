package helpers

import (
    "time"
    "spectra/ds3_go_sdk/ds3/models"
)

var MinUploadSize int64 = 10485760 //todo move to bulk put??

// Defines strategy for how to perform write transfers
type WriteTransferStrategy struct {
    BlobStrategy WriteBlobStrategy
    Options      WriteBulkJobOptions
}

// Defines the options to use on the put bulk job
type WriteBulkJobOptions struct {
    Aggregating *bool
    ImplicitJobIdResolution *bool
    MaxUploadSize *int64
    MinimizeSpanningAcrossMedia *bool
    Priority *models.Priority
    VerifyAfterWrite *bool
    Name *string
    Force *bool
    IgnoreNamingConflicts *bool
}

// Strategy for how to blob objects for writing
type WriteBlobStrategy interface {
    delay() // performs delay when no job chunks are available
    maxTransferGoroutines() int // determines the maximum number of go routines to be created when transferring objects to/from BP
}

type SimpleWriteBlobStrategy struct {
    Delay time.Duration
    MaxTransferGoroutines int
}

func (sbs *SimpleWriteBlobStrategy) delay() {
    time.Sleep(sbs.Delay)
}

func (sbs *SimpleWriteBlobStrategy) maxTransferGoroutines() int {
    return sbs.MaxTransferGoroutines
}
