package helpers

import "sync"

type NotifyBlobDone interface {
    // Waits for at least one done signal.
    Wait()

    // Sends a done signal. Multiple signals have no additional effect.
    SignalDone()
}


func NewConditionalBool() *ConditionalBool {
    conditional :=sync.NewCond(&sync.Mutex{})
    return &ConditionalBool{
        conditional: *conditional,
        Done: false,
    }
}

type ConditionalBool struct {
    conditional sync.Cond
    Done bool
}

func (conditionalBool *ConditionalBool) Wait() {
    conditionalBool.conditional.L.Lock()
    // wait for a done signal to be received
    for !conditionalBool.Done {
        conditionalBool.conditional.Wait()
    }
    // reset done notifier
    conditionalBool.Done = false
    conditionalBool.conditional.L.Unlock()
}

func (conditionalBool *ConditionalBool) SignalDone() {
    conditionalBool.conditional.L.Lock()
    conditionalBool.Done = true
    conditionalBool.conditional.Broadcast()
    conditionalBool.conditional.L.Unlock()
}