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

type DatabasePhysicalSpaceState Enum

const (
    DATABASE_PHYSICAL_SPACE_STATE_CRITICAL DatabasePhysicalSpaceState = 1 + iota
    DATABASE_PHYSICAL_SPACE_STATE_LOW DatabasePhysicalSpaceState = 1 + iota
    DATABASE_PHYSICAL_SPACE_STATE_NEAR_LOW DatabasePhysicalSpaceState = 1 + iota
    DATABASE_PHYSICAL_SPACE_STATE_NORMAL DatabasePhysicalSpaceState = 1 + iota
)

func (databasePhysicalSpaceState *DatabasePhysicalSpaceState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *databasePhysicalSpaceState = UNDEFINED
        case "CRITICAL": *databasePhysicalSpaceState = DATABASE_PHYSICAL_SPACE_STATE_CRITICAL
        case "LOW": *databasePhysicalSpaceState = DATABASE_PHYSICAL_SPACE_STATE_LOW
        case "NEAR_LOW": *databasePhysicalSpaceState = DATABASE_PHYSICAL_SPACE_STATE_NEAR_LOW
        case "NORMAL": *databasePhysicalSpaceState = DATABASE_PHYSICAL_SPACE_STATE_NORMAL
        default:
            *databasePhysicalSpaceState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into DatabasePhysicalSpaceState", str))
    }
    return nil
}

func (databasePhysicalSpaceState DatabasePhysicalSpaceState) String() string {
    switch databasePhysicalSpaceState {
        case DATABASE_PHYSICAL_SPACE_STATE_CRITICAL: return "CRITICAL"
        case DATABASE_PHYSICAL_SPACE_STATE_LOW: return "LOW"
        case DATABASE_PHYSICAL_SPACE_STATE_NEAR_LOW: return "NEAR_LOW"
        case DATABASE_PHYSICAL_SPACE_STATE_NORMAL: return "NORMAL"
        default:
            log.Printf("Error: invalid DatabasePhysicalSpaceState represented by '%d'", databasePhysicalSpaceState)
            return ""
    }
}

func (databasePhysicalSpaceState DatabasePhysicalSpaceState) StringPtr() *string {
    if databasePhysicalSpaceState == UNDEFINED {
        return nil
    }
    result := databasePhysicalSpaceState.String()
    return &result
}