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

type DataReplicationRuleType Enum

const (
    DATA_REPLICATION_RULE_TYPE_PERMANENT DataReplicationRuleType = 1 + iota
    DATA_REPLICATION_RULE_TYPE_RETIRED DataReplicationRuleType = 1 + iota
)

func (dataReplicationRuleType *DataReplicationRuleType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "PERMANENT": *dataReplicationRuleType = DATA_REPLICATION_RULE_TYPE_PERMANENT
        case "RETIRED": *dataReplicationRuleType = DATA_REPLICATION_RULE_TYPE_RETIRED
        default:
            *dataReplicationRuleType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into DataReplicationRuleType", str))
    }
    return nil
}

func (dataReplicationRuleType DataReplicationRuleType) String() string {
    switch dataReplicationRuleType {
        case DATA_REPLICATION_RULE_TYPE_PERMANENT: return "PERMANENT"
        case DATA_REPLICATION_RULE_TYPE_RETIRED: return "RETIRED"
        default:
            log.Printf("Error: invalid DataReplicationRuleType represented by '%d'", dataReplicationRuleType)
            return ""
    }
}
