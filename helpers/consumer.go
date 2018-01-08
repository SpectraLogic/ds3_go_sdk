package helpers

import (
    "sync"
)

type Consumer interface {
    run()
}

type putConsumer struct {
    queue *chan TransferOperation
    wg *sync.WaitGroup
    maxPutsPerJob int
}

func newPutConsumer(queue *chan TransferOperation, wg *sync.WaitGroup, maxPutsPerJob int) *putConsumer {
    return &putConsumer{
        queue:queue,
        wg:wg,
        maxPutsPerJob:maxPutsPerJob,
    }
}

func performTransfer(operation *TransferOperation, sem *chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    (*operation)()
    <- *sem
}

func (pc *putConsumer) run() {
    sem := make(chan int, pc.maxPutsPerJob) // semaphore for controlling max number of transfer operations in flight per job
    for {
        nextOp, ok := <- *pc.queue
        if ok {
            sem <- 1
            go performTransfer(&nextOp, &sem, pc.wg)
        } else {
            pc.wg.Done()
            return
        }
    }
}
