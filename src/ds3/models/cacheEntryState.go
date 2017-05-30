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

type CacheEntryState Enum

const (
    CACHE_ENTRY_STATE_ALLOCATED CacheEntryState = 1 + iota
    CACHE_ENTRY_STATE_IN_CACHE CacheEntryState = 1 + iota
)

func (cacheEntryState *CacheEntryState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *cacheEntryState = UNDEFINED
        case "ALLOCATED": *cacheEntryState = CACHE_ENTRY_STATE_ALLOCATED
        case "IN_CACHE": *cacheEntryState = CACHE_ENTRY_STATE_IN_CACHE
        default:
            *cacheEntryState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into CacheEntryState", str))
    }
    return nil
}

func (cacheEntryState CacheEntryState) String() string {
    switch cacheEntryState {
        case CACHE_ENTRY_STATE_ALLOCATED: return "ALLOCATED"
        case CACHE_ENTRY_STATE_IN_CACHE: return "IN_CACHE"
        default:
            log.Printf("Error: invalid CacheEntryState represented by '%d'", cacheEntryState)
            return ""
    }
}
