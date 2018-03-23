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
    maxConcurrentOperations int
}

func newConsumer(queue *chan TransferOperation, waitGroup *sync.WaitGroup, maxConcurrentOperations int) Consumer {
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

func (pc *consumerImpl) run() {
    semaphore := make(chan int, pc.maxConcurrentOperations) // semaphore for controlling max number of transfer operations in flight per job
    for {
        nextOp, ok := <- *pc.queue
        if ok {
            semaphore <- 1
            go performTransfer(&nextOp, &semaphore, pc.waitGroup)
        } else {
            pc.waitGroup.Done()
            return
        }
    }
}
