package commands

import "bytes"

// Provides a way to combine error messages together. This is useful for
// displaying errors that occur during parallel processing.
type aggregateError struct {
    errors []error
}

// Adds an error to the aggregate.
func (aggregateError *aggregateError) Add(err error) {
    aggregateError.errors = append(aggregateError.errors, err)
}

// Implements the error interface.
func (aggregateError *aggregateError) Error() string {
    var buffer bytes.Buffer
    for _, err := range aggregateError.errors {
        buffer.WriteString(err.Error())
        buffer.WriteString("\n")
    }
    return buffer.String()
}

