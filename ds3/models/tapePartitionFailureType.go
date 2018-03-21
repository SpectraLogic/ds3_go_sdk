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

type TapePartitionFailureType Enum

const (
    TAPE_PARTITION_FAILURE_TYPE_CLEANING_TAPE_REQUIRED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_DUPLICATE_TAPE_BAR_CODES_DETECTED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_EJECT_STALLED_DUE_TO_OFFLINE_TAPES TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_MINIMUM_DRIVE_COUNT_NOT_MET TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED_DUE_TO_PREPARE_TAPE_FOR_REMOVAL_FAILURE TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_NO_USABLE_DRIVES TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_ONLINE_STALLED_DUE_TO_NO_STORAGE_SLOTS TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_IN_ERROR TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_MISSING TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_QUIESCED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_TYPE_MISMATCH TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_EJECTION_BY_OPERATOR_REQUIRED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_MEDIA_TYPE_INCOMPATIBLE TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_REMOVAL_UNEXPECTED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_IN_INVALID_PARTITION TapePartitionFailureType = 1 + iota
)

func (tapePartitionFailureType *TapePartitionFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapePartitionFailureType = UNDEFINED
        case "CLEANING_TAPE_REQUIRED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_CLEANING_TAPE_REQUIRED
        case "DUPLICATE_TAPE_BAR_CODES_DETECTED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_DUPLICATE_TAPE_BAR_CODES_DETECTED
        case "EJECT_STALLED_DUE_TO_OFFLINE_TAPES": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_EJECT_STALLED_DUE_TO_OFFLINE_TAPES
        case "MINIMUM_DRIVE_COUNT_NOT_MET": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_MINIMUM_DRIVE_COUNT_NOT_MET
        case "MOVE_FAILED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED
        case "MOVE_FAILED_DUE_TO_PREPARE_TAPE_FOR_REMOVAL_FAILURE": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED_DUE_TO_PREPARE_TAPE_FOR_REMOVAL_FAILURE
        case "NO_USABLE_DRIVES": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_NO_USABLE_DRIVES
        case "ONLINE_STALLED_DUE_TO_NO_STORAGE_SLOTS": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_ONLINE_STALLED_DUE_TO_NO_STORAGE_SLOTS
        case "TAPE_DRIVE_IN_ERROR": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_IN_ERROR
        case "TAPE_DRIVE_MISSING": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_MISSING
        case "TAPE_DRIVE_QUIESCED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_QUIESCED
        case "TAPE_DRIVE_TYPE_MISMATCH": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_TYPE_MISMATCH
        case "TAPE_EJECTION_BY_OPERATOR_REQUIRED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_EJECTION_BY_OPERATOR_REQUIRED
        case "TAPE_MEDIA_TYPE_INCOMPATIBLE": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_MEDIA_TYPE_INCOMPATIBLE
        case "TAPE_REMOVAL_UNEXPECTED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_REMOVAL_UNEXPECTED
        case "TAPE_IN_INVALID_PARTITION": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_IN_INVALID_PARTITION
        default:
            *tapePartitionFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapePartitionFailureType", str))
    }
    return nil
}

func (tapePartitionFailureType TapePartitionFailureType) String() string {
    switch tapePartitionFailureType {
        case TAPE_PARTITION_FAILURE_TYPE_CLEANING_TAPE_REQUIRED: return "CLEANING_TAPE_REQUIRED"
        case TAPE_PARTITION_FAILURE_TYPE_DUPLICATE_TAPE_BAR_CODES_DETECTED: return "DUPLICATE_TAPE_BAR_CODES_DETECTED"
        case TAPE_PARTITION_FAILURE_TYPE_EJECT_STALLED_DUE_TO_OFFLINE_TAPES: return "EJECT_STALLED_DUE_TO_OFFLINE_TAPES"
        case TAPE_PARTITION_FAILURE_TYPE_MINIMUM_DRIVE_COUNT_NOT_MET: return "MINIMUM_DRIVE_COUNT_NOT_MET"
        case TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED: return "MOVE_FAILED"
        case TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED_DUE_TO_PREPARE_TAPE_FOR_REMOVAL_FAILURE: return "MOVE_FAILED_DUE_TO_PREPARE_TAPE_FOR_REMOVAL_FAILURE"
        case TAPE_PARTITION_FAILURE_TYPE_NO_USABLE_DRIVES: return "NO_USABLE_DRIVES"
        case TAPE_PARTITION_FAILURE_TYPE_ONLINE_STALLED_DUE_TO_NO_STORAGE_SLOTS: return "ONLINE_STALLED_DUE_TO_NO_STORAGE_SLOTS"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_IN_ERROR: return "TAPE_DRIVE_IN_ERROR"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_MISSING: return "TAPE_DRIVE_MISSING"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_QUIESCED: return "TAPE_DRIVE_QUIESCED"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_TYPE_MISMATCH: return "TAPE_DRIVE_TYPE_MISMATCH"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_EJECTION_BY_OPERATOR_REQUIRED: return "TAPE_EJECTION_BY_OPERATOR_REQUIRED"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_MEDIA_TYPE_INCOMPATIBLE: return "TAPE_MEDIA_TYPE_INCOMPATIBLE"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_REMOVAL_UNEXPECTED: return "TAPE_REMOVAL_UNEXPECTED"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_IN_INVALID_PARTITION: return "TAPE_IN_INVALID_PARTITION"
        default:
            log.Printf("Error: invalid TapePartitionFailureType represented by '%d'", tapePartitionFailureType)
            return ""
    }
}

func (tapePartitionFailureType TapePartitionFailureType) StringPtr() *string {
    if tapePartitionFailureType == UNDEFINED {
        return nil
    }
    result := tapePartitionFailureType.String()
    return &result
}