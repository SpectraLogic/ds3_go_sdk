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

    // make the blob done channel larger than the number of transfer operations queued.
    blobDoneChannel := make(chan struct{}, numOperations+1)

    consumer := newConsumer(&queue, blobDoneChannel, &wg, 5)

    go producer(&queue)
    go consumer.run()

    wg.Wait()

    ds3Testing.AssertInt(t, "Executed Transfer Operations", numOperations, resultCount)

    // verify that 10 done messages were sent
    ds3Testing.AssertInt(t, "Done signals sent", numOperations, len(blobDoneChannel))
    for len(blobDoneChannel) > 0 {
        _, ok := <-blobDoneChannel
        ds3Testing.AssertBool(t, "expected channel not to be closed", true, ok)
    }
    _, ok := <- blobDoneChannel
    ds3Testing.AssertBool(t, "expected channel to be closed", false, ok)
}
