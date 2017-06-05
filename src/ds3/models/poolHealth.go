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

type PoolHealth Enum

const (
    POOL_HEALTH_OK PoolHealth = 1 + iota
    POOL_HEALTH_DEGRADED PoolHealth = 1 + iota
)

func (poolHealth *PoolHealth) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *poolHealth = UNDEFINED
        case "OK": *poolHealth = POOL_HEALTH_OK
        case "DEGRADED": *poolHealth = POOL_HEALTH_DEGRADED
        default:
            *poolHealth = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into PoolHealth", str))
    }
    return nil
}

func (poolHealth PoolHealth) String() string {
    switch poolHealth {
        case POOL_HEALTH_OK: return "OK"
        case POOL_HEALTH_DEGRADED: return "DEGRADED"
        default:
            log.Printf("Error: invalid PoolHealth represented by '%d'", poolHealth)
            return ""
    }
}
