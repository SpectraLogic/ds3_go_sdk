package helpers

type Transceiver interface {
    transfer() error
    //cancel()
}
