package helpers

import (
    helperModels"spectra/ds3_go_sdk/helpers/models"
)

type blobTracker struct {
    blobMap map[helperModels.BlobDescription]bool
    blobCount int64
}

func newProcessedBlobTracker() blobTracker {
    return blobTracker{
        blobMap: make(map[helperModels.BlobDescription]bool),
        blobCount: 0,
    }
}

// Determines if a blob has been processed
func (tracker *blobTracker) IsProcessed(blob helperModels.BlobDescription) bool {
    return tracker.blobMap[blob]
}

func (tracker *blobTracker) MarkProcessed(blob helperModels.BlobDescription) {
    tracker.blobMap[blob] = true
    tracker.blobCount++
}

func (tracker *blobTracker) NumberOfProcessedBlobs() int64 {
    return tracker.blobCount
}
