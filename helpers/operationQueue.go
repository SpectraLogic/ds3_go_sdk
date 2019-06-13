package helpers

import (
    "github.com/SpectraLogic/ds3_go_sdk/sdk_log"
)

type TransferOperation func() // transfer operation that sends/gets stuff from BP

const MinQueueSize uint = 1
const MaxQueueSize uint = 100

func newOperationQueue(size uint, logger sdk_log.Logger) chan TransferOperation {
    var queue chan TransferOperation

    if size > MaxQueueSize {
        logger.Warningf("invalid operation queue size: specified value '%d' which exceeds the maximum, defaulting to '%d'", size, MaxQueueSize)
        queue = make(chan TransferOperation, MaxQueueSize)
    } else if size < MinQueueSize {
        logger.Warningf("invalid operation queue size: specified value '%d' which is below the minimum, defaulting to '%d'", size, MinQueueSize)
        queue = make(chan TransferOperation, MinQueueSize)
    } else {
        queue = make(chan TransferOperation, size)
    }

    return queue
}
