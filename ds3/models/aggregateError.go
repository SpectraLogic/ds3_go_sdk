// Copyright 2014-2017 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package models

import (
    "fmt"
    "sync"
)

// AggregateError is an error that aggregates multiple errors and is multi thread safe.
type AggregateError struct {
    Errors []error
    mux sync.RWMutex
}

func (aggregateError *AggregateError) Error() string {
    aggregateError.mux.RLock()
    defer aggregateError.mux.RUnlock()

    msg := fmt.Sprintf("Multiple errors occured: %d\n", len(aggregateError.Errors))

    for i, err := range aggregateError.Errors {
        msg += fmt.Sprintf("%d) %s\n", i + 1, err.Error())
    }

    return msg
}

// Returns the aggregate error if at least one error exists,
// else returns nil
func (aggregateError *AggregateError) GetErrors() error {
    aggregateError.mux.RLock()
    defer aggregateError.mux.RUnlock()

    if len (aggregateError.Errors) == 0 {
        return nil
    }
    return aggregateError
}

func (aggregateError *AggregateError) Append(err error) {
    aggregateError.mux.Lock()
    defer aggregateError.mux.Unlock()

    if err != nil {
        aggregateError.Errors = append(aggregateError.Errors, err)
    }
}