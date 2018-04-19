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
}

func newConsumer(queue *chan TransferOperation, waitGroup *sync.WaitGroup, maxConcurrentOperations uint) Consumer {
    return &consumerImpl{
        queue:                   queue,
        waitGroup:               waitGroup,
        maxConcurrentOperations: maxConcurrentOperations,
    }
}

func performTransfer(operation *TransferOperation, semaphore *chan int, waitGroup *sync.WaitGroup) {
    defer waitGroup.Done()
    (*operation)()
    <- *semaphore
}

func (consumer *consumerImpl) run() {
    // semaphore for controlling max number of transfer operations in flight per job
    semaphore := make(chan int, consumer.maxConcurrentOperations + 1)
    for {
        nextOp, ok := <- *consumer.queue
        if ok {
            semaphore <- 1
            go performTransfer(&nextOp, &semaphore, consumer.waitGroup)
        } else {
            consumer.waitGroup.Done()
            return
        }
    }
}
