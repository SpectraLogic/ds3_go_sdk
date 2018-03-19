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

type ReservedTaskType Enum

const (
    RESERVED_TASK_TYPE_ANY ReservedTaskType = 1 + iota
    RESERVED_TASK_TYPE_READ ReservedTaskType = 1 + iota
    RESERVED_TASK_TYPE_WRITE ReservedTaskType = 1 + iota
)

func (reservedTaskType *ReservedTaskType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *reservedTaskType = UNDEFINED
        case "ANY": *reservedTaskType = RESERVED_TASK_TYPE_ANY
        case "READ": *reservedTaskType = RESERVED_TASK_TYPE_READ
        case "WRITE": *reservedTaskType = RESERVED_TASK_TYPE_WRITE
        default:
            *reservedTaskType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into ReservedTaskType", str))
    }
    return nil
}

func (reservedTaskType ReservedTaskType) String() string {
    switch reservedTaskType {
        case RESERVED_TASK_TYPE_ANY: return "ANY"
        case RESERVED_TASK_TYPE_READ: return "READ"
        case RESERVED_TASK_TYPE_WRITE: return "WRITE"
        default:
            log.Printf("Error: invalid ReservedTaskType represented by '%d'", reservedTaskType)
            return ""
    }
}

func (reservedTaskType ReservedTaskType) StringPtr() *string {
    if reservedTaskType == UNDEFINED {
        return nil
    }
    result := reservedTaskType.String()
    return &result
}