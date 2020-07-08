package helpers

import (
    "sync"
)

type Consumer interface {
    run()
}

type consumerImpl struct {
    queue                   *chan TransferOperation
    waitGroup               *sync.WaitGroup
    maxConcurrentOperations uint

    // Conditional value that gets triggered when a blob has finished being transferred
    doneNotifier NotifyBlobDone
}

func newConsumer(queue *chan TransferOperation, waitGroup *sync.WaitGroup, maxConcurrentOperations uint, doneNotifier NotifyBlobDone) Consumer {
    return &consumerImpl{
        queue:                   queue,
        waitGroup:               waitGroup,
        maxConcurrentOperations: maxConcurrentOperations,
        doneNotifier:            doneNotifier,
    }
}

func performTransfer(operation TransferOperation, semaphore *chan int, waitGroup *sync.WaitGroup, doneNotifier NotifyBlobDone) {
    defer waitGroup.Done()
    operation()

    // send done signal
    doneNotifier.SignalDone()

    <- *semaphore
}

func (consumer *consumerImpl) run() {
    // semaphore for controlling max number of transfer operations in flight per job
    semaphore := make(chan int, consumer.maxConcurrentOperations + 1)

    for {
        nextOp, ok := <- *consumer.queue
        if ok {
            semaphore <- 1
            go performTransfer(nextOp, &semaphore, consumer.waitGroup, consumer.doneNotifier)
        } else {
            consumer.waitGroup.Done()
            return
        }
    }
}
