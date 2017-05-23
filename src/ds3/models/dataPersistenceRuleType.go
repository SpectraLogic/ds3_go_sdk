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

type DataPersistenceRuleType Enum

const (
    DATA_PERSISTENCE_RULE_TYPE_PERMANENT DataPersistenceRuleType = 1 + iota
    DATA_PERSISTENCE_RULE_TYPE_TEMPORARY DataPersistenceRuleType = 1 + iota
    DATA_PERSISTENCE_RULE_TYPE_RETIRED DataPersistenceRuleType = 1 + iota
)

func (dataPersistenceRuleType *DataPersistenceRuleType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "PERMANENT": *dataPersistenceRuleType = DATA_PERSISTENCE_RULE_TYPE_PERMANENT
        case "TEMPORARY": *dataPersistenceRuleType = DATA_PERSISTENCE_RULE_TYPE_TEMPORARY
        case "RETIRED": *dataPersistenceRuleType = DATA_PERSISTENCE_RULE_TYPE_RETIRED
        default:
            *dataPersistenceRuleType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into DataPersistenceRuleType", str))
    }
    return nil
}

func (dataPersistenceRuleType DataPersistenceRuleType) String() string {
    switch dataPersistenceRuleType {
        case DATA_PERSISTENCE_RULE_TYPE_PERMANENT: return "PERMANENT"
        case DATA_PERSISTENCE_RULE_TYPE_TEMPORARY: return "TEMPORARY"
        case DATA_PERSISTENCE_RULE_TYPE_RETIRED: return "RETIRED"
        default:
            log.Printf("Error: invalid DataPersistenceRuleType represented by '%d'", dataPersistenceRuleType)
            return ""
    }
}
