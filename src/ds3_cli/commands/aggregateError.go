package commands

import "bytes"

// Provides a way to combine error messages together. This is useful for
// displaying errors that occur during parallel processing.
type aggregateError struct {
    errors []error
}

// Adds an error to the aggregate.
func (self *aggregateError) Add(err error) {
    self.errors = append(self.errors, err)
}

// Implements the error interface.
func (self *aggregateError) Error() string {
    var buffer bytes.Buffer
    for _, err := range self.errors {
        buffer.WriteString(err.Error())
        buffer.WriteString("\n")
    }
    return buffer.String()
}

