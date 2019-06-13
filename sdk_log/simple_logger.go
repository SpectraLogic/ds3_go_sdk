package sdk_log

import (
    "log"
)

func NewSimpleLogger() Logger {
    return &SimpleLogger{}
}

type SimpleLogger struct {}

func (logger *SimpleLogger) Infof(format string, args ...interface{}) {
    log.Printf("INFO " + format, args...)
}

func (logger *SimpleLogger) Debugf(format string, args ...interface{}) {
    log.Printf("DEBUG " + format, args...)
}

func (logger *SimpleLogger) Warningf(format string, args ...interface{}) {
    log.Printf("WARNING " + format, args...)
}

func (logger *SimpleLogger) Errorf(format string, args ...interface{}) {
    log.Printf("ERROR " + format, args...)
}
