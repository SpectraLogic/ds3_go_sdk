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
    blobDoneChannel         chan<- struct{}
}

func newConsumer(queue *chan TransferOperation, blobDoneChannel chan<- struct{}, waitGroup *sync.WaitGroup, maxConcurrentOperations uint) Consumer {
    return &consumerImpl{
        queue:                   queue,
        waitGroup:               waitGroup,
        maxConcurrentOperations: maxConcurrentOperations,
        blobDoneChannel:         blobDoneChannel,
    }
}

func performTransfer(operation TransferOperation, semaphore *chan int, blobDoneChannel chan<- struct{}, jobWaitGroup *sync.WaitGroup, childWaitGroup *sync.WaitGroup) {
    defer func() {
        // per operation that finishes, send a done message to the producer
        blobDoneChannel <-struct {}{}
        jobWaitGroup.Done()
        childWaitGroup.Done()
    }()
    operation()
    <- *semaphore
}

func (consumer *consumerImpl) run() {
    // Defer closing the blob done channel. This will signal to the producer that it can shut down.
    defer func() {close(consumer.blobDoneChannel)}()

    // semaphore for controlling max number of transfer operations in flight per job
    semaphore := make(chan int, consumer.maxConcurrentOperations + 1)

    var childWaitGroup sync.WaitGroup
    for {
        nextOp, ok := <- *consumer.queue
        if ok {
            semaphore <- 1
            childWaitGroup.Add(1)
            go performTransfer(nextOp, &semaphore, consumer.blobDoneChannel, consumer.waitGroup, &childWaitGroup)
        } else {
            consumer.waitGroup.Done()
            break
        }
    }

    // Wait for all child transfer operations to finish before shutting down.
    // This is to stop the done channel from being close prematurely
    childWaitGroup.Wait()
}
