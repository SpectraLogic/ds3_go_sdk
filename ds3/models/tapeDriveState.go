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

type TapeDriveState Enum

const (
    TAPE_DRIVE_STATE_OFFLINE TapeDriveState = 1 + iota
    TAPE_DRIVE_STATE_NORMAL TapeDriveState = 1 + iota
    TAPE_DRIVE_STATE_ERROR TapeDriveState = 1 + iota
    TAPE_DRIVE_STATE_NOT_COMPATIBLE_IN_PARTITION_DUE_TO_NEWER_TAPE_DRIVES TapeDriveState = 1 + iota
)

func (tapeDriveState *TapeDriveState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapeDriveState = UNDEFINED
        case "OFFLINE": *tapeDriveState = TAPE_DRIVE_STATE_OFFLINE
        case "NORMAL": *tapeDriveState = TAPE_DRIVE_STATE_NORMAL
        case "ERROR": *tapeDriveState = TAPE_DRIVE_STATE_ERROR
        case "NOT_COMPATIBLE_IN_PARTITION_DUE_TO_NEWER_TAPE_DRIVES": *tapeDriveState = TAPE_DRIVE_STATE_NOT_COMPATIBLE_IN_PARTITION_DUE_TO_NEWER_TAPE_DRIVES
        default:
            *tapeDriveState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapeDriveState", str))
    }
    return nil
}

func (tapeDriveState TapeDriveState) String() string {
    switch tapeDriveState {
        case TAPE_DRIVE_STATE_OFFLINE: return "OFFLINE"
        case TAPE_DRIVE_STATE_NORMAL: return "NORMAL"
        case TAPE_DRIVE_STATE_ERROR: return "ERROR"
        case TAPE_DRIVE_STATE_NOT_COMPATIBLE_IN_PARTITION_DUE_TO_NEWER_TAPE_DRIVES: return "NOT_COMPATIBLE_IN_PARTITION_DUE_TO_NEWER_TAPE_DRIVES"
        default:
            log.Printf("Error: invalid TapeDriveState represented by '%d'", tapeDriveState)
            return ""
    }
}

func (tapeDriveState TapeDriveState) StringPtr() *string {
    if tapeDriveState == UNDEFINED {
        return nil
    }
    result := tapeDriveState.String()
    return &result
}