package commands

import "bytes"

type aggregateError struct {
    errors []error
}

func (self *aggregateError) Add(err error) {
    self.errors = append(self.errors, err)
}

func (self *aggregateError) Error() string {
    var buffer bytes.Buffer
    for _, err := range self.errors {
        buffer.WriteString(err.Error())
        buffer.WriteString("\n")
    }
    return buffer.String()
}

