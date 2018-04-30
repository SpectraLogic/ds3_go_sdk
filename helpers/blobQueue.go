package helpers

import (
    "errors"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
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

func (queue *blobDescriptionQueueImpl) Push(description *helperModels.BlobDescription) {
    queue.queue = append(queue.queue, description)
}

func (queue *blobDescriptionQueueImpl) Pop() (*helperModels.BlobDescription, error) {
    if queue.Size() == 0 {
        return nil, errors.New("Cannot perform Pop() from blobDescriptionQueueImpl as queue is empty")
    }
    descriptor := queue.queue[0]
    queue.queue = queue.queue[1:]
    return descriptor, nil
}

func (queue *blobDescriptionQueueImpl) Size() int {
    return len(queue.queue)
}
