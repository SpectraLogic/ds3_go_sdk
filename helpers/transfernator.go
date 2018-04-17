package helpers

type Transfernator interface {
    transfer() error
    cancel()
}
