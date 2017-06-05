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

type DataPlacementRuleState Enum

const (
    DATA_PLACEMENT_RULE_STATE_NORMAL DataPlacementRuleState = 1 + iota
    DATA_PLACEMENT_RULE_STATE_INCLUSION_IN_PROGRESS DataPlacementRuleState = 1 + iota
)

func (dataPlacementRuleState *DataPlacementRuleState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *dataPlacementRuleState = UNDEFINED
        case "NORMAL": *dataPlacementRuleState = DATA_PLACEMENT_RULE_STATE_NORMAL
        case "INCLUSION_IN_PROGRESS": *dataPlacementRuleState = DATA_PLACEMENT_RULE_STATE_INCLUSION_IN_PROGRESS
        default:
            *dataPlacementRuleState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into DataPlacementRuleState", str))
    }
    return nil
}

func (dataPlacementRuleState DataPlacementRuleState) String() string {
    switch dataPlacementRuleState {
        case DATA_PLACEMENT_RULE_STATE_NORMAL: return "NORMAL"
        case DATA_PLACEMENT_RULE_STATE_INCLUSION_IN_PROGRESS: return "INCLUSION_IN_PROGRESS"
        default:
            log.Printf("Error: invalid DataPlacementRuleState represented by '%d'", dataPlacementRuleState)
            return ""
    }
}
