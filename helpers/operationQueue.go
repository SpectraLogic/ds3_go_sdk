package helpers

type TransferOperation func() // transfer operation that sends/gets stuff from BP

const MinQueueSize int = 1
const MaxQueueSize int = 5

func newOperationQueue(size int) chan TransferOperation {
    var queue chan TransferOperation

    if size > MaxQueueSize {
        queue = make(chan TransferOperation, MaxQueueSize)
    } else if size < MinQueueSize {
        queue = make(chan TransferOperation, MinQueueSize)
    } else {
        queue = make(chan TransferOperation, size)
    }

    return queue
}
