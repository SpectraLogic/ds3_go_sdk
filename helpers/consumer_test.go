package helpers

import (
    "testing"
    "sync"
    "fmt"
    "spectra/ds3_go_sdk/ds3_utils/ds3Testing"
)

func testTransferBuilder(i int, waitGroup *sync.WaitGroup, resultCount *int, resultMux *sync.Mutex) TransferOperation {
    return func() {
        //defer waitGroup.Done()

        resultMux.Lock()
        *resultCount++
        resultMux.Unlock()

        fmt.Printf("Transfer Op: '%d'\n", i)
    }
}

func TestProducerConsumerModel(t *testing.T) {
    var wg sync.WaitGroup
    wg.Add(2)

    resultCount := 0
    var resultMux sync.Mutex

    var producer = func(queue *chan TransferOperation) {
        defer wg.Done()
        for i := 0; i < 10; i++ {
            wg.Add(1)

            var transferOf = testTransferBuilder(i, &wg, &resultCount, &resultMux)

            fmt.Printf("Producer: '%d'\n", i)

            *queue <- transferOf
        }
        close(*queue)
    }

    queue := make(chan TransferOperation, 5)

    consumer := newConsumer(&queue, &wg, 5)

    go producer(&queue)
    go consumer.run()

    wg.Wait()

    ds3Testing.AssertInt(t, "Executed Transfer Operations", 10, resultCount)
}
