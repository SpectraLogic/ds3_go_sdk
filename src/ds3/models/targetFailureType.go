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

type TargetFailureType Enum

const (
    TARGET_FAILURE_TYPE_IMPORT_FAILED TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_NOT_ONLINE TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_WRITE_FAILED TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_WRITE_INITIATE_FAILED TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_READ_FAILED TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_READ_INITIATE_FAILED TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_VERIFY_FAILED TargetFailureType = 1 + iota
)

func (targetFailureType *TargetFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *targetFailureType = UNDEFINED
        case "IMPORT_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_IMPORT_FAILED
        case "NOT_ONLINE": *targetFailureType = TARGET_FAILURE_TYPE_NOT_ONLINE
        case "WRITE_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_WRITE_FAILED
        case "WRITE_INITIATE_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_WRITE_INITIATE_FAILED
        case "READ_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_READ_FAILED
        case "READ_INITIATE_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_READ_INITIATE_FAILED
        case "VERIFY_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_VERIFY_FAILED
        default:
            *targetFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TargetFailureType", str))
    }
    return nil
}

func (targetFailureType TargetFailureType) String() string {
    switch targetFailureType {
        case TARGET_FAILURE_TYPE_IMPORT_FAILED: return "IMPORT_FAILED"
        case TARGET_FAILURE_TYPE_NOT_ONLINE: return "NOT_ONLINE"
        case TARGET_FAILURE_TYPE_WRITE_FAILED: return "WRITE_FAILED"
        case TARGET_FAILURE_TYPE_WRITE_INITIATE_FAILED: return "WRITE_INITIATE_FAILED"
        case TARGET_FAILURE_TYPE_READ_FAILED: return "READ_FAILED"
        case TARGET_FAILURE_TYPE_READ_INITIATE_FAILED: return "READ_INITIATE_FAILED"
        case TARGET_FAILURE_TYPE_VERIFY_FAILED: return "VERIFY_FAILED"
        default:
            log.Printf("Error: invalid TargetFailureType represented by '%d'", targetFailureType)
            return ""
    }
}
