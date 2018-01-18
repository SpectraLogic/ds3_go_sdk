package helpers

import "sync"

// TODO delete whole file
// Tracks which chunks have been processed by the producer.
type ProcessedChunksTracker interface {
    IsProcessed(chunkId string) bool // determines if a chunk has been processed yet
    SetProcessed(chunkId string) // marks a chunk as having been processed
    NumberOfProcessedChunks() int // returns the total number of chunks that have been processed
}

type processedChunksTrackerImpl struct {
    processedMap map[string]bool
    mux sync.RWMutex
}

func NewProccessedChunksTracker() ProcessedChunksTracker {
    processedMap := make(map[string]bool)
    return &processedChunksTrackerImpl{processedMap:processedMap,}
}

func (t *processedChunksTrackerImpl) IsProcessed(chunkId string) bool {
    t.mux.RLock()
    defer t.mux.RUnlock()
    return t.processedMap[chunkId]
}

func (t *processedChunksTrackerImpl) SetProcessed(chunkId string) {
    t.mux.Lock()
    defer t.mux.Unlock()
    t.processedMap[chunkId] = true
}

func (t *processedChunksTrackerImpl) NumberOfProcessedChunks() int {
    t.mux.RLock()
    defer t.mux.RUnlock()
    return len(t.processedMap)
}
