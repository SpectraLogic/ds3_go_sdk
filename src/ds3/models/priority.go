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

type Priority Enum

const (
    PRIORITY_CRITICAL Priority = 1 + iota
    PRIORITY_URGENT Priority = 1 + iota
    PRIORITY_HIGH Priority = 1 + iota
    PRIORITY_NORMAL Priority = 1 + iota
    PRIORITY_LOW Priority = 1 + iota
    PRIORITY_BACKGROUND Priority = 1 + iota
)

func (priority *Priority) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *priority = UNDEFINED
        case "CRITICAL": *priority = PRIORITY_CRITICAL
        case "URGENT": *priority = PRIORITY_URGENT
        case "HIGH": *priority = PRIORITY_HIGH
        case "NORMAL": *priority = PRIORITY_NORMAL
        case "LOW": *priority = PRIORITY_LOW
        case "BACKGROUND": *priority = PRIORITY_BACKGROUND
        default:
            *priority = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into Priority", str))
    }
    return nil
}

func (priority Priority) String() string {
    switch priority {
        case PRIORITY_CRITICAL: return "CRITICAL"
        case PRIORITY_URGENT: return "URGENT"
        case PRIORITY_HIGH: return "HIGH"
        case PRIORITY_NORMAL: return "NORMAL"
        case PRIORITY_LOW: return "LOW"
        case PRIORITY_BACKGROUND: return "BACKGROUND"
        default:
            log.Printf("Error: invalid Priority represented by '%d'", priority)
            return ""
    }
}
