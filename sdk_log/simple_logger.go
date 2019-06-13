package sdk_log

import (
    "log"
)

// A logger which prints everything to log.Printf with the level as a prefix.
func NewSimpleLogger() Logger {
    return &SimpleLogger{}
}

type SimpleLogger struct {}

func (SimpleLogger) Infof(format string, args ...interface{}) {
    log.Printf("INFO " + format, args...)
}

func (SimpleLogger) Debugf(format string, args ...interface{}) {
    log.Printf("DEBUG " + format, args...)
}

func (SimpleLogger) Warningf(format string, args ...interface{}) {
    log.Printf("WARNING " + format, args...)
}

func (SimpleLogger) Errorf(format string, args ...interface{}) {
    log.Printf("ERROR " + format, args...)
}
