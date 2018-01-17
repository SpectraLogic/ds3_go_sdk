package helpers

import (
    "errors"
    helperModels "spectra/ds3_go_sdk/helpers/models"
)

// A queue that manages descriptions of blobs
// Used to track blobs that are waiting to be transferred
type BlobDescriptionQueue interface {
    Push(description *helperModels.BlobDescription)
    Pop() (*helperModels.BlobDescription, error)
    Size() int
}

// Implements BlobDescriptionQueue using a slice
// NOT thread safe
type blobDescriptionQueueImpl struct {
    queue []*helperModels.BlobDescription
}

func NewBlobDescriptionQueue() BlobDescriptionQueue {
    queue :=make([]*helperModels.BlobDescription, 0)
    return &blobDescriptionQueueImpl{queue:queue}
}

func (q *blobDescriptionQueueImpl) Push(description *helperModels.BlobDescription) {
    q.queue = append(q.queue, description)
}

func (q *blobDescriptionQueueImpl) Pop() (*helperModels.BlobDescription, error) {
    if q.Size() == 0 {
        return nil, errors.New("Cannot perform Pop() from blobDescriptionQueueImpl as queue is empty")
    }
    descriptor := q.queue[0]
    q.queue = q.queue[1:]
    return descriptor, nil
}

func (q *blobDescriptionQueueImpl) Size() int {
    return len(q.queue)
}
