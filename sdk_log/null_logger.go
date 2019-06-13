package sdk_log

func NewNullLogger() Logger {
    return &NullLogger{}
}

// A logger which does nothing.
type NullLogger struct {}

func (NullLogger) Infof(format string, args ...interface{}) {
    // do nothing
}

func (NullLogger) Debugf(format string, args ...interface{}) {
    // do nothing
}

func (NullLogger) Warningf(format string, args ...interface{}) {
    // do nothing
}

func (NullLogger) Errorf(format string, args ...interface{}) {
    // do nothing
}
