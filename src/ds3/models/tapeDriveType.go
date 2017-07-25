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

type TapeDriveType Enum

const (
    TAPE_DRIVE_TYPE_UNKNOWN TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_LTO5 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_LTO6 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_LTO7 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_LTO8 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_TS1140 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_TS1150 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_TS1155 TapeDriveType = 1 + iota
)

func (tapeDriveType *TapeDriveType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapeDriveType = UNDEFINED
        case "UNKNOWN": *tapeDriveType = TAPE_DRIVE_TYPE_UNKNOWN
        case "LTO5": *tapeDriveType = TAPE_DRIVE_TYPE_LTO5
        case "LTO6": *tapeDriveType = TAPE_DRIVE_TYPE_LTO6
        case "LTO7": *tapeDriveType = TAPE_DRIVE_TYPE_LTO7
        case "LTO8": *tapeDriveType = TAPE_DRIVE_TYPE_LTO8
        case "TS1140": *tapeDriveType = TAPE_DRIVE_TYPE_TS1140
        case "TS1150": *tapeDriveType = TAPE_DRIVE_TYPE_TS1150
        case "TS1155": *tapeDriveType = TAPE_DRIVE_TYPE_TS1155
        default:
            *tapeDriveType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapeDriveType", str))
    }
    return nil
}

func (tapeDriveType TapeDriveType) String() string {
    switch tapeDriveType {
        case TAPE_DRIVE_TYPE_UNKNOWN: return "UNKNOWN"
        case TAPE_DRIVE_TYPE_LTO5: return "LTO5"
        case TAPE_DRIVE_TYPE_LTO6: return "LTO6"
        case TAPE_DRIVE_TYPE_LTO7: return "LTO7"
        case TAPE_DRIVE_TYPE_LTO8: return "LTO8"
        case TAPE_DRIVE_TYPE_TS1140: return "TS1140"
        case TAPE_DRIVE_TYPE_TS1150: return "TS1150"
        case TAPE_DRIVE_TYPE_TS1155: return "TS1155"
        default:
            log.Printf("Error: invalid TapeDriveType represented by '%d'", tapeDriveType)
            return ""
    }
}
