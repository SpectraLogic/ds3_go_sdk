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

type WritePreferenceLevel Enum

const (
    WRITE_PREFERENCE_LEVEL_HIGH WritePreferenceLevel = 1 + iota
    WRITE_PREFERENCE_LEVEL_NORMAL WritePreferenceLevel = 1 + iota
    WRITE_PREFERENCE_LEVEL_LOW WritePreferenceLevel = 1 + iota
    WRITE_PREFERENCE_LEVEL_NEVER_SELECT WritePreferenceLevel = 1 + iota
)

func (writePreferenceLevel *WritePreferenceLevel) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "HIGH": *writePreferenceLevel = WRITE_PREFERENCE_LEVEL_HIGH
        case "NORMAL": *writePreferenceLevel = WRITE_PREFERENCE_LEVEL_NORMAL
        case "LOW": *writePreferenceLevel = WRITE_PREFERENCE_LEVEL_LOW
        case "NEVER_SELECT": *writePreferenceLevel = WRITE_PREFERENCE_LEVEL_NEVER_SELECT
        default:
            *writePreferenceLevel = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into WritePreferenceLevel", str))
    }
    return nil
}

func (writePreferenceLevel WritePreferenceLevel) String() string {
    switch writePreferenceLevel {
        case WRITE_PREFERENCE_LEVEL_HIGH: return "HIGH"
        case WRITE_PREFERENCE_LEVEL_NORMAL: return "NORMAL"
        case WRITE_PREFERENCE_LEVEL_LOW: return "LOW"
        case WRITE_PREFERENCE_LEVEL_NEVER_SELECT: return "NEVER_SELECT"
        default:
            log.Printf("Error: invalid WritePreferenceLevel represented by '%d'", writePreferenceLevel)
            return ""
    }
}
