package helpers

import (
    "github.com/SpectraLogic/ds3_go_sdk/ds3_utils/ds3Testing"
    "sync"
    "testing"
)

func testTransferBuilder(t *testing.T, i int, resultCount *int, resultMux *sync.Mutex) TransferOperation {
    return func() {
        resultMux.Lock()
        *resultCount++
        resultMux.Unlock()

        t.Logf("Transfer Op: '%d'\n", i)
    }
}

func TestProducerConsumerModel(t *testing.T) {
    const numOperations = 10
    var wg sync.WaitGroup
    wg.Add(2)

    resultCount := 0
    var resultMux sync.Mutex

    var producer = func(queue *chan TransferOperation) {
        defer wg.Done()
        for i := 0; i < numOperations; i++ {
            wg.Add(1)

            var transferOf = testTransferBuilder(t, i, &resultCount, &resultMux)

            t.Logf("Producer: '%d'\n", i)

            *queue <- transferOf
        }
        close(*queue)
    }

    queue := make(chan TransferOperation, 5)

    doneNotifier := NewConditionalBool()

    consumer := newConsumer(&queue, &wg, 5, doneNotifier)

    go producer(&queue)
    go consumer.run()

    wg.Wait()

    ds3Testing.AssertInt(t, "Executed Transfer Operations", numOperations, resultCount)
    ds3Testing.AssertBool(t, "received done notification", true, doneNotifier.Done)
}
