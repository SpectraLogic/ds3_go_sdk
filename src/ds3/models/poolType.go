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

type PoolType Enum

const (
    POOL_TYPE_NEARLINE PoolType = 1 + iota
    POOL_TYPE_ONLINE PoolType = 1 + iota
)

func (poolType *PoolType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *poolType = UNDEFINED
        case "NEARLINE": *poolType = POOL_TYPE_NEARLINE
        case "ONLINE": *poolType = POOL_TYPE_ONLINE
        default:
            *poolType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into PoolType", str))
    }
    return nil
}

func (poolType PoolType) String() string {
    switch poolType {
        case POOL_TYPE_NEARLINE: return "NEARLINE"
        case POOL_TYPE_ONLINE: return "ONLINE"
        default:
            log.Printf("Error: invalid PoolType represented by '%d'", poolType)
            return ""
    }
}

func (poolType PoolType) StringPtr() *string {
    if poolType == UNDEFINED {
        return nil
    }
    result := poolType.String()
    return &result
}