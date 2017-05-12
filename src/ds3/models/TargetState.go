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

type TargetState Enum

const (
    TARGET_STATE_ONLINE TargetState = 1 + iota
    TARGET_STATE_OFFLINE TargetState = 1 + iota
)

func (targetState *TargetState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "ONLINE": *targetState = TARGET_STATE_ONLINE
        case "OFFLINE": *targetState = TARGET_STATE_OFFLINE
        default:
            *targetState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TargetState", str))
    }
    return nil
}

func (targetState TargetState) String() string {
    switch targetState {
        case TARGET_STATE_ONLINE: return "ONLINE"
        case TARGET_STATE_OFFLINE: return "OFFLINE"
        default:
            log.Printf("Error: invalid TargetState represented by '%d'", targetState)
            return ""
    }
}
