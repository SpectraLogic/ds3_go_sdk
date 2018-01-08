package helpers

import "errors"

// Describes a blob
type BlobDescription struct {
    objectName string
    offset     int64
    length     int64
}

func (d *BlobDescription) Name() string {
    return d.objectName
}

func (d *BlobDescription) Offset() int64 {
    return d.offset
}

func (d *BlobDescription) Length() int64 {
    return d.length
}

// A queue that manages descriptions of blobs
// Used to track blobs that are waiting to be transferred
type BlobDescriptionQueue interface {
    Push(description *BlobDescription)
    Pop() (*BlobDescription, error)
    Size() int
}

// Implements BlobDescriptionQueue using a slice
// NOT thread safe
type blobDescriptionQueueImpl struct {
    queue []*BlobDescription
}

func NewBlobDescriptionQueue() BlobDescriptionQueue {
    queue :=make([]*BlobDescription, 0)
    return &blobDescriptionQueueImpl{queue:queue}
}

func (q *blobDescriptionQueueImpl) Push(description *BlobDescription) {
    q.queue = append(q.queue, description)
}

func (q *blobDescriptionQueueImpl) Pop() (*BlobDescription, error) {
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
