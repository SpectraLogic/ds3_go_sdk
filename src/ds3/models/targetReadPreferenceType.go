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

type TargetReadPreferenceType Enum

const (
    TARGET_READ_PREFERENCE_TYPE_MINIMUM_LATENCY TargetReadPreferenceType = 1 + iota
    TARGET_READ_PREFERENCE_TYPE_AFTER_ONLINE_POOL TargetReadPreferenceType = 1 + iota
    TARGET_READ_PREFERENCE_TYPE_AFTER_NEARLINE_POOL TargetReadPreferenceType = 1 + iota
    TARGET_READ_PREFERENCE_TYPE_AFTER_NON_EJECTABLE_TAPE TargetReadPreferenceType = 1 + iota
    TARGET_READ_PREFERENCE_TYPE_LAST_RESORT TargetReadPreferenceType = 1 + iota
    TARGET_READ_PREFERENCE_TYPE_NEVER TargetReadPreferenceType = 1 + iota
)

func (targetReadPreferenceType *TargetReadPreferenceType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *targetReadPreferenceType = UNDEFINED
        case "MINIMUM_LATENCY": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_MINIMUM_LATENCY
        case "AFTER_ONLINE_POOL": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_AFTER_ONLINE_POOL
        case "AFTER_NEARLINE_POOL": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_AFTER_NEARLINE_POOL
        case "AFTER_NON_EJECTABLE_TAPE": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_AFTER_NON_EJECTABLE_TAPE
        case "LAST_RESORT": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_LAST_RESORT
        case "NEVER": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_NEVER
        default:
            *targetReadPreferenceType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TargetReadPreferenceType", str))
    }
    return nil
}

func (targetReadPreferenceType TargetReadPreferenceType) String() string {
    switch targetReadPreferenceType {
        case TARGET_READ_PREFERENCE_TYPE_MINIMUM_LATENCY: return "MINIMUM_LATENCY"
        case TARGET_READ_PREFERENCE_TYPE_AFTER_ONLINE_POOL: return "AFTER_ONLINE_POOL"
        case TARGET_READ_PREFERENCE_TYPE_AFTER_NEARLINE_POOL: return "AFTER_NEARLINE_POOL"
        case TARGET_READ_PREFERENCE_TYPE_AFTER_NON_EJECTABLE_TAPE: return "AFTER_NON_EJECTABLE_TAPE"
        case TARGET_READ_PREFERENCE_TYPE_LAST_RESORT: return "LAST_RESORT"
        case TARGET_READ_PREFERENCE_TYPE_NEVER: return "NEVER"
        default:
            log.Printf("Error: invalid TargetReadPreferenceType represented by '%d'", targetReadPreferenceType)
            return ""
    }
}

func (targetReadPreferenceType TargetReadPreferenceType) StringPtr() *string {
    if targetReadPreferenceType == UNDEFINED {
        return nil
    }
    result := targetReadPreferenceType.String()
    return &result
}