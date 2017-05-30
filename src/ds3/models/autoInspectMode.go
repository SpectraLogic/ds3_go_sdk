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

type AutoInspectMode Enum

const (
    AUTO_INSPECT_MODE_NEVER AutoInspectMode = 1 + iota
    AUTO_INSPECT_MODE_MINIMAL AutoInspectMode = 1 + iota
    AUTO_INSPECT_MODE_FULL AutoInspectMode = 1 + iota
)

func (autoInspectMode *AutoInspectMode) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *autoInspectMode = UNDEFINED
        case "NEVER": *autoInspectMode = AUTO_INSPECT_MODE_NEVER
        case "MINIMAL": *autoInspectMode = AUTO_INSPECT_MODE_MINIMAL
        case "FULL": *autoInspectMode = AUTO_INSPECT_MODE_FULL
        default:
            *autoInspectMode = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into AutoInspectMode", str))
    }
    return nil
}

func (autoInspectMode AutoInspectMode) String() string {
    switch autoInspectMode {
        case AUTO_INSPECT_MODE_NEVER: return "NEVER"
        case AUTO_INSPECT_MODE_MINIMAL: return "MINIMAL"
        case AUTO_INSPECT_MODE_FULL: return "FULL"
        default:
            log.Printf("Error: invalid AutoInspectMode represented by '%d'", autoInspectMode)
            return ""
    }
}
