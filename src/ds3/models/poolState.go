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

type PoolState Enum

const (
    POOL_STATE_BLANK PoolState = 1 + iota
    POOL_STATE_NORMAL PoolState = 1 + iota
    POOL_STATE_LOST PoolState = 1 + iota
    POOL_STATE_FOREIGN PoolState = 1 + iota
    POOL_STATE_IMPORT_PENDING PoolState = 1 + iota
    POOL_STATE_IMPORT_IN_PROGRESS PoolState = 1 + iota
)

func (poolState *PoolState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "BLANK": *poolState = POOL_STATE_BLANK
        case "NORMAL": *poolState = POOL_STATE_NORMAL
        case "LOST": *poolState = POOL_STATE_LOST
        case "FOREIGN": *poolState = POOL_STATE_FOREIGN
        case "IMPORT_PENDING": *poolState = POOL_STATE_IMPORT_PENDING
        case "IMPORT_IN_PROGRESS": *poolState = POOL_STATE_IMPORT_IN_PROGRESS
        default:
            *poolState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into PoolState", str))
    }
    return nil
}

func (poolState PoolState) String() string {
    switch poolState {
        case POOL_STATE_BLANK: return "BLANK"
        case POOL_STATE_NORMAL: return "NORMAL"
        case POOL_STATE_LOST: return "LOST"
        case POOL_STATE_FOREIGN: return "FOREIGN"
        case POOL_STATE_IMPORT_PENDING: return "IMPORT_PENDING"
        case POOL_STATE_IMPORT_IN_PROGRESS: return "IMPORT_IN_PROGRESS"
        default:
            log.Printf("Error: invalid PoolState represented by '%d'", poolState)
            return ""
    }
}
