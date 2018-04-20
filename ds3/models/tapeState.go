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

type TapeState Enum

const (
    TAPE_STATE_NORMAL TapeState = 1 + iota
    TAPE_STATE_OFFLINE TapeState = 1 + iota
    TAPE_STATE_ONLINE_PENDING TapeState = 1 + iota
    TAPE_STATE_ONLINE_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_PENDING_INSPECTION TapeState = 1 + iota
    TAPE_STATE_UNKNOWN TapeState = 1 + iota
    TAPE_STATE_DATA_CHECKPOINT_FAILURE TapeState = 1 + iota
    TAPE_STATE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY TapeState = 1 + iota
    TAPE_STATE_DATA_CHECKPOINT_MISSING TapeState = 1 + iota
    TAPE_STATE_LTFS_WITH_FOREIGN_DATA TapeState = 1 + iota
    TAPE_STATE_RAW_IMPORT_PENDING TapeState = 1 + iota
    TAPE_STATE_RAW_IMPORT_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_FOREIGN TapeState = 1 + iota
    TAPE_STATE_IMPORT_PENDING TapeState = 1 + iota
    TAPE_STATE_IMPORT_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_INCOMPATIBLE TapeState = 1 + iota
    TAPE_STATE_LOST TapeState = 1 + iota
    TAPE_STATE_BAD TapeState = 1 + iota
    TAPE_STATE_CANNOT_FORMAT_DUE_TO_WRITE_PROTECTION TapeState = 1 + iota
    TAPE_STATE_SERIAL_NUMBER_MISMATCH TapeState = 1 + iota
    TAPE_STATE_BAR_CODE_MISSING TapeState = 1 + iota
    TAPE_STATE_AUTO_COMPACTION_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_FORMAT_PENDING TapeState = 1 + iota
    TAPE_STATE_FORMAT_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_EJECT_TO_EE_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_EJECT_FROM_EE_PENDING TapeState = 1 + iota
    TAPE_STATE_EJECTED TapeState = 1 + iota
)

func (tapeState *TapeState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapeState = UNDEFINED
        case "NORMAL": *tapeState = TAPE_STATE_NORMAL
        case "OFFLINE": *tapeState = TAPE_STATE_OFFLINE
        case "ONLINE_PENDING": *tapeState = TAPE_STATE_ONLINE_PENDING
        case "ONLINE_IN_PROGRESS": *tapeState = TAPE_STATE_ONLINE_IN_PROGRESS
        case "PENDING_INSPECTION": *tapeState = TAPE_STATE_PENDING_INSPECTION
        case "UNKNOWN": *tapeState = TAPE_STATE_UNKNOWN
        case "DATA_CHECKPOINT_FAILURE": *tapeState = TAPE_STATE_DATA_CHECKPOINT_FAILURE
        case "DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY": *tapeState = TAPE_STATE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY
        case "DATA_CHECKPOINT_MISSING": *tapeState = TAPE_STATE_DATA_CHECKPOINT_MISSING
        case "LTFS_WITH_FOREIGN_DATA": *tapeState = TAPE_STATE_LTFS_WITH_FOREIGN_DATA
        case "RAW_IMPORT_PENDING": *tapeState = TAPE_STATE_RAW_IMPORT_PENDING
        case "RAW_IMPORT_IN_PROGRESS": *tapeState = TAPE_STATE_RAW_IMPORT_IN_PROGRESS
        case "FOREIGN": *tapeState = TAPE_STATE_FOREIGN
        case "IMPORT_PENDING": *tapeState = TAPE_STATE_IMPORT_PENDING
        case "IMPORT_IN_PROGRESS": *tapeState = TAPE_STATE_IMPORT_IN_PROGRESS
        case "INCOMPATIBLE": *tapeState = TAPE_STATE_INCOMPATIBLE
        case "LOST": *tapeState = TAPE_STATE_LOST
        case "BAD": *tapeState = TAPE_STATE_BAD
        case "CANNOT_FORMAT_DUE_TO_WRITE_PROTECTION": *tapeState = TAPE_STATE_CANNOT_FORMAT_DUE_TO_WRITE_PROTECTION
        case "SERIAL_NUMBER_MISMATCH": *tapeState = TAPE_STATE_SERIAL_NUMBER_MISMATCH
        case "BAR_CODE_MISSING": *tapeState = TAPE_STATE_BAR_CODE_MISSING
        case "AUTO_COMPACTION_IN_PROGRESS": *tapeState = TAPE_STATE_AUTO_COMPACTION_IN_PROGRESS
        case "FORMAT_PENDING": *tapeState = TAPE_STATE_FORMAT_PENDING
        case "FORMAT_IN_PROGRESS": *tapeState = TAPE_STATE_FORMAT_IN_PROGRESS
        case "EJECT_TO_EE_IN_PROGRESS": *tapeState = TAPE_STATE_EJECT_TO_EE_IN_PROGRESS
        case "EJECT_FROM_EE_PENDING": *tapeState = TAPE_STATE_EJECT_FROM_EE_PENDING
        case "EJECTED": *tapeState = TAPE_STATE_EJECTED
        default:
            *tapeState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapeState", str))
    }
    return nil
}

func (tapeState TapeState) String() string {
    switch tapeState {
        case TAPE_STATE_NORMAL: return "NORMAL"
        case TAPE_STATE_OFFLINE: return "OFFLINE"
        case TAPE_STATE_ONLINE_PENDING: return "ONLINE_PENDING"
        case TAPE_STATE_ONLINE_IN_PROGRESS: return "ONLINE_IN_PROGRESS"
        case TAPE_STATE_PENDING_INSPECTION: return "PENDING_INSPECTION"
        case TAPE_STATE_UNKNOWN: return "UNKNOWN"
        case TAPE_STATE_DATA_CHECKPOINT_FAILURE: return "DATA_CHECKPOINT_FAILURE"
        case TAPE_STATE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY: return "DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY"
        case TAPE_STATE_DATA_CHECKPOINT_MISSING: return "DATA_CHECKPOINT_MISSING"
        case TAPE_STATE_LTFS_WITH_FOREIGN_DATA: return "LTFS_WITH_FOREIGN_DATA"
        case TAPE_STATE_RAW_IMPORT_PENDING: return "RAW_IMPORT_PENDING"
        case TAPE_STATE_RAW_IMPORT_IN_PROGRESS: return "RAW_IMPORT_IN_PROGRESS"
        case TAPE_STATE_FOREIGN: return "FOREIGN"
        case TAPE_STATE_IMPORT_PENDING: return "IMPORT_PENDING"
        case TAPE_STATE_IMPORT_IN_PROGRESS: return "IMPORT_IN_PROGRESS"
        case TAPE_STATE_INCOMPATIBLE: return "INCOMPATIBLE"
        case TAPE_STATE_LOST: return "LOST"
        case TAPE_STATE_BAD: return "BAD"
        case TAPE_STATE_CANNOT_FORMAT_DUE_TO_WRITE_PROTECTION: return "CANNOT_FORMAT_DUE_TO_WRITE_PROTECTION"
        case TAPE_STATE_SERIAL_NUMBER_MISMATCH: return "SERIAL_NUMBER_MISMATCH"
        case TAPE_STATE_BAR_CODE_MISSING: return "BAR_CODE_MISSING"
        case TAPE_STATE_AUTO_COMPACTION_IN_PROGRESS: return "AUTO_COMPACTION_IN_PROGRESS"
        case TAPE_STATE_FORMAT_PENDING: return "FORMAT_PENDING"
        case TAPE_STATE_FORMAT_IN_PROGRESS: return "FORMAT_IN_PROGRESS"
        case TAPE_STATE_EJECT_TO_EE_IN_PROGRESS: return "EJECT_TO_EE_IN_PROGRESS"
        case TAPE_STATE_EJECT_FROM_EE_PENDING: return "EJECT_FROM_EE_PENDING"
        case TAPE_STATE_EJECTED: return "EJECTED"
        default:
            log.Printf("Error: invalid TapeState represented by '%d'", tapeState)
            return ""
    }
}

func (tapeState TapeState) StringPtr() *string {
    if tapeState == UNDEFINED {
        return nil
    }
    result := tapeState.String()
    return &result
}