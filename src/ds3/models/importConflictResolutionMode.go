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

type ImportConflictResolutionMode Enum

const (
    IMPORT_CONFLICT_RESOLUTION_MODE_CANCEL ImportConflictResolutionMode = 1 + iota
    IMPORT_CONFLICT_RESOLUTION_MODE_ACCEPT_MOST_RECENT ImportConflictResolutionMode = 1 + iota
    IMPORT_CONFLICT_RESOLUTION_MODE_ACCEPT_EXISTING ImportConflictResolutionMode = 1 + iota
)

func (importConflictResolutionMode *ImportConflictResolutionMode) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *importConflictResolutionMode = UNDEFINED
        case "CANCEL": *importConflictResolutionMode = IMPORT_CONFLICT_RESOLUTION_MODE_CANCEL
        case "ACCEPT_MOST_RECENT": *importConflictResolutionMode = IMPORT_CONFLICT_RESOLUTION_MODE_ACCEPT_MOST_RECENT
        case "ACCEPT_EXISTING": *importConflictResolutionMode = IMPORT_CONFLICT_RESOLUTION_MODE_ACCEPT_EXISTING
        default:
            *importConflictResolutionMode = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into ImportConflictResolutionMode", str))
    }
    return nil
}

func (importConflictResolutionMode ImportConflictResolutionMode) String() string {
    switch importConflictResolutionMode {
        case IMPORT_CONFLICT_RESOLUTION_MODE_CANCEL: return "CANCEL"
        case IMPORT_CONFLICT_RESOLUTION_MODE_ACCEPT_MOST_RECENT: return "ACCEPT_MOST_RECENT"
        case IMPORT_CONFLICT_RESOLUTION_MODE_ACCEPT_EXISTING: return "ACCEPT_EXISTING"
        default:
            log.Printf("Error: invalid ImportConflictResolutionMode represented by '%d'", importConflictResolutionMode)
            return ""
    }
}
