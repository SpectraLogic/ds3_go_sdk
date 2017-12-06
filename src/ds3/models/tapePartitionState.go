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

type TapePartitionState Enum

const (
    TAPE_PARTITION_STATE_ONLINE TapePartitionState = 1 + iota
    TAPE_PARTITION_STATE_OFFLINE TapePartitionState = 1 + iota
    TAPE_PARTITION_STATE_ERROR TapePartitionState = 1 + iota
)

func (tapePartitionState *TapePartitionState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapePartitionState = UNDEFINED
        case "ONLINE": *tapePartitionState = TAPE_PARTITION_STATE_ONLINE
        case "OFFLINE": *tapePartitionState = TAPE_PARTITION_STATE_OFFLINE
        case "ERROR": *tapePartitionState = TAPE_PARTITION_STATE_ERROR
        default:
            *tapePartitionState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapePartitionState", str))
    }
    return nil
}

func (tapePartitionState TapePartitionState) String() string {
    switch tapePartitionState {
        case TAPE_PARTITION_STATE_ONLINE: return "ONLINE"
        case TAPE_PARTITION_STATE_OFFLINE: return "OFFLINE"
        case TAPE_PARTITION_STATE_ERROR: return "ERROR"
        default:
            log.Printf("Error: invalid TapePartitionState represented by '%d'", tapePartitionState)
            return ""
    }
}

func (tapePartitionState TapePartitionState) StringPtr() *string {
    if tapePartitionState == UNDEFINED {
        return nil
    }
    result := tapePartitionState.String()
    return &result
}