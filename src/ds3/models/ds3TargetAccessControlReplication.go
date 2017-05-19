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

type Ds3TargetAccessControlReplication Enum

const (
    DS3_TARGET_ACCESS_CONTROL_REPLICATION_NONE Ds3TargetAccessControlReplication = 1 + iota
    DS3_TARGET_ACCESS_CONTROL_REPLICATION_USERS Ds3TargetAccessControlReplication = 1 + iota
)

func (ds3TargetAccessControlReplication *Ds3TargetAccessControlReplication) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "NONE": *ds3TargetAccessControlReplication = DS3_TARGET_ACCESS_CONTROL_REPLICATION_NONE
        case "USERS": *ds3TargetAccessControlReplication = DS3_TARGET_ACCESS_CONTROL_REPLICATION_USERS
        default:
            *ds3TargetAccessControlReplication = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into Ds3TargetAccessControlReplication", str))
    }
    return nil
}

func (ds3TargetAccessControlReplication Ds3TargetAccessControlReplication) String() string {
    switch ds3TargetAccessControlReplication {
        case DS3_TARGET_ACCESS_CONTROL_REPLICATION_NONE: return "NONE"
        case DS3_TARGET_ACCESS_CONTROL_REPLICATION_USERS: return "USERS"
        default:
            log.Printf("Error: invalid Ds3TargetAccessControlReplication represented by '%d'", ds3TargetAccessControlReplication)
            return ""
    }
}
