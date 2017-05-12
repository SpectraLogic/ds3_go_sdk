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

type TapeType Enum

const (
    TAPE_TYPE_LTO5 TapeType = 1 + iota
    TAPE_TYPE_LTO6 TapeType = 1 + iota
    TAPE_TYPE_LTO7 TapeType = 1 + iota
    TAPE_TYPE_LTO_CLEANING_TAPE TapeType = 1 + iota
    TAPE_TYPE_TS_JC TapeType = 1 + iota
    TAPE_TYPE_TS_JY TapeType = 1 + iota
    TAPE_TYPE_TS_JK TapeType = 1 + iota
    TAPE_TYPE_TS_JD TapeType = 1 + iota
    TAPE_TYPE_TS_JZ TapeType = 1 + iota
    TAPE_TYPE_TS_JL TapeType = 1 + iota
    TAPE_TYPE_TS_CLEANING_TAPE TapeType = 1 + iota
    TAPE_TYPE_UNKNOWN TapeType = 1 + iota
    TAPE_TYPE_FORBIDDEN TapeType = 1 + iota
)

func (tapeType *TapeType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "LTO5": *tapeType = TAPE_TYPE_LTO5
        case "LTO6": *tapeType = TAPE_TYPE_LTO6
        case "LTO7": *tapeType = TAPE_TYPE_LTO7
        case "LTO_CLEANING_TAPE": *tapeType = TAPE_TYPE_LTO_CLEANING_TAPE
        case "TS_JC": *tapeType = TAPE_TYPE_TS_JC
        case "TS_JY": *tapeType = TAPE_TYPE_TS_JY
        case "TS_JK": *tapeType = TAPE_TYPE_TS_JK
        case "TS_JD": *tapeType = TAPE_TYPE_TS_JD
        case "TS_JZ": *tapeType = TAPE_TYPE_TS_JZ
        case "TS_JL": *tapeType = TAPE_TYPE_TS_JL
        case "TS_CLEANING_TAPE": *tapeType = TAPE_TYPE_TS_CLEANING_TAPE
        case "UNKNOWN": *tapeType = TAPE_TYPE_UNKNOWN
        case "FORBIDDEN": *tapeType = TAPE_TYPE_FORBIDDEN
        default:
            *tapeType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapeType", str))
    }
    return nil
}

func (tapeType TapeType) String() string {
    switch tapeType {
        case TAPE_TYPE_LTO5: return "LTO5"
        case TAPE_TYPE_LTO6: return "LTO6"
        case TAPE_TYPE_LTO7: return "LTO7"
        case TAPE_TYPE_LTO_CLEANING_TAPE: return "LTO_CLEANING_TAPE"
        case TAPE_TYPE_TS_JC: return "TS_JC"
        case TAPE_TYPE_TS_JY: return "TS_JY"
        case TAPE_TYPE_TS_JK: return "TS_JK"
        case TAPE_TYPE_TS_JD: return "TS_JD"
        case TAPE_TYPE_TS_JZ: return "TS_JZ"
        case TAPE_TYPE_TS_JL: return "TS_JL"
        case TAPE_TYPE_TS_CLEANING_TAPE: return "TS_CLEANING_TAPE"
        case TAPE_TYPE_UNKNOWN: return "UNKNOWN"
        case TAPE_TYPE_FORBIDDEN: return "FORBIDDEN"
        default:
            log.Printf("Error: invalid TapeType represented by '%d'", tapeType)
            return ""
    }
}
