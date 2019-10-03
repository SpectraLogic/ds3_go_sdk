package helpers

type ListenerStrategy struct {
    // Called when an error occurred during transfer of an object.
    // This must be a thread safe function.
    ErrorCallback func(objectName string, err error)
}

func (listener *ListenerStrategy) Errored(objectName string, err error) {
    if listener.ErrorCallback != nil {
        listener.ErrorCallback(objectName, err)
    }
}
