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

type TapeFailureType Enum

const (
    TAPE_FAILURE_TYPE_BAR_CODE_CHANGED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_BAR_CODE_DUPLICATE TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_BLOB_READ_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DATA_CHECKPOINT_MISSING TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DELAYED_OWNERSHIP_FAILURE TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DRIVE_CLEAN_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DRIVE_CLEANED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_FORMAT_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_GET_TAPE_INFORMATION_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_IMPORT_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_INSPECT_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_READ_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_REIMPORT_REQUIRED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_SERIAL_NUMBER_MISMATCH TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_VERIFY_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_WRITE_FAILED TapeFailureType = 1 + iota
)

func (tapeFailureType *TapeFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapeFailureType = UNDEFINED
        case "BAR_CODE_CHANGED": *tapeFailureType = TAPE_FAILURE_TYPE_BAR_CODE_CHANGED
        case "BAR_CODE_DUPLICATE": *tapeFailureType = TAPE_FAILURE_TYPE_BAR_CODE_DUPLICATE
        case "BLOB_READ_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_BLOB_READ_FAILED
        case "DATA_CHECKPOINT_FAILURE": *tapeFailureType = TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE
        case "DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY": *tapeFailureType = TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY
        case "DATA_CHECKPOINT_MISSING": *tapeFailureType = TAPE_FAILURE_TYPE_DATA_CHECKPOINT_MISSING
        case "DELAYED_OWNERSHIP_FAILURE": *tapeFailureType = TAPE_FAILURE_TYPE_DELAYED_OWNERSHIP_FAILURE
        case "DRIVE_CLEAN_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_DRIVE_CLEAN_FAILED
        case "DRIVE_CLEANED": *tapeFailureType = TAPE_FAILURE_TYPE_DRIVE_CLEANED
        case "FORMAT_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_FORMAT_FAILED
        case "GET_TAPE_INFORMATION_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_GET_TAPE_INFORMATION_FAILED
        case "IMPORT_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_IMPORT_FAILED
        case "IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE": *tapeFailureType = TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE
        case "IMPORT_FAILED_DUE_TO_DATA_INTEGRITY": *tapeFailureType = TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY
        case "INSPECT_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_INSPECT_FAILED
        case "READ_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_READ_FAILED
        case "REIMPORT_REQUIRED": *tapeFailureType = TAPE_FAILURE_TYPE_REIMPORT_REQUIRED
        case "SERIAL_NUMBER_MISMATCH": *tapeFailureType = TAPE_FAILURE_TYPE_SERIAL_NUMBER_MISMATCH
        case "VERIFY_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_VERIFY_FAILED
        case "WRITE_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_WRITE_FAILED
        default:
            *tapeFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapeFailureType", str))
    }
    return nil
}

func (tapeFailureType TapeFailureType) String() string {
    switch tapeFailureType {
        case TAPE_FAILURE_TYPE_BAR_CODE_CHANGED: return "BAR_CODE_CHANGED"
        case TAPE_FAILURE_TYPE_BAR_CODE_DUPLICATE: return "BAR_CODE_DUPLICATE"
        case TAPE_FAILURE_TYPE_BLOB_READ_FAILED: return "BLOB_READ_FAILED"
        case TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE: return "DATA_CHECKPOINT_FAILURE"
        case TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY: return "DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY"
        case TAPE_FAILURE_TYPE_DATA_CHECKPOINT_MISSING: return "DATA_CHECKPOINT_MISSING"
        case TAPE_FAILURE_TYPE_DELAYED_OWNERSHIP_FAILURE: return "DELAYED_OWNERSHIP_FAILURE"
        case TAPE_FAILURE_TYPE_DRIVE_CLEAN_FAILED: return "DRIVE_CLEAN_FAILED"
        case TAPE_FAILURE_TYPE_DRIVE_CLEANED: return "DRIVE_CLEANED"
        case TAPE_FAILURE_TYPE_FORMAT_FAILED: return "FORMAT_FAILED"
        case TAPE_FAILURE_TYPE_GET_TAPE_INFORMATION_FAILED: return "GET_TAPE_INFORMATION_FAILED"
        case TAPE_FAILURE_TYPE_IMPORT_FAILED: return "IMPORT_FAILED"
        case TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE: return "IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE"
        case TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY: return "IMPORT_FAILED_DUE_TO_DATA_INTEGRITY"
        case TAPE_FAILURE_TYPE_INSPECT_FAILED: return "INSPECT_FAILED"
        case TAPE_FAILURE_TYPE_READ_FAILED: return "READ_FAILED"
        case TAPE_FAILURE_TYPE_REIMPORT_REQUIRED: return "REIMPORT_REQUIRED"
        case TAPE_FAILURE_TYPE_SERIAL_NUMBER_MISMATCH: return "SERIAL_NUMBER_MISMATCH"
        case TAPE_FAILURE_TYPE_VERIFY_FAILED: return "VERIFY_FAILED"
        case TAPE_FAILURE_TYPE_WRITE_FAILED: return "WRITE_FAILED"
        default:
            log.Printf("Error: invalid TapeFailureType represented by '%d'", tapeFailureType)
            return ""
    }
}
