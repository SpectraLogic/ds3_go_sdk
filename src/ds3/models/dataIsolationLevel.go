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

// This code is auto-generated, do not modify

package models

import (
    "errors"
    "fmt"
    "bytes"
    "log"
)

type DataIsolationLevel Enum

const (
    DATA_ISOLATION_LEVEL_STANDARD DataIsolationLevel = 1 + iota
    DATA_ISOLATION_LEVEL_BUCKET_ISOLATED DataIsolationLevel = 1 + iota
)

func (dataIsolationLevel *DataIsolationLevel) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "STANDARD": *dataIsolationLevel = DATA_ISOLATION_LEVEL_STANDARD
        case "BUCKET_ISOLATED": *dataIsolationLevel = DATA_ISOLATION_LEVEL_BUCKET_ISOLATED
        default:
            *dataIsolationLevel = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into DataIsolationLevel", str))
    }
    return nil
}

func (dataIsolationLevel DataIsolationLevel) String() string {
    switch dataIsolationLevel {
        case DATA_ISOLATION_LEVEL_STANDARD: return "STANDARD"
        case DATA_ISOLATION_LEVEL_BUCKET_ISOLATED: return "BUCKET_ISOLATED"
        default:
            log.Printf("Error: invalid DataIsolationLevel represented by '%d'", dataIsolationLevel)
            return ""
    }
}
