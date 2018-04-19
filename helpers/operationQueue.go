package helpers

import "log"

type TransferOperation func() // transfer operation that sends/gets stuff from BP

const MinQueueSize uint = 1
const MaxQueueSize uint = 100

func newOperationQueue(size uint) chan TransferOperation {
    var queue chan TransferOperation

    if size > MaxQueueSize {
        log.Printf("WARNING Invalid operation queue size: specified value '%d' which exceeds the maximum, defaulting to '%d'\n", size, MaxQueueSize)
        queue = make(chan TransferOperation, MaxQueueSize)
    } else if size < MinQueueSize {
        log.Printf("WARNING Invalid operation queue size: specified value '%d' which is below the minimum, defaulting to '%d'\n", size, MinQueueSize)
        queue = make(chan TransferOperation, MinQueueSize)
    } else {
        queue = make(chan TransferOperation, size)
    }

    return queue
}
