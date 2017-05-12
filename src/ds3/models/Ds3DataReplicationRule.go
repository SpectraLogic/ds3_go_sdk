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

type Ds3DataReplicationRule struct {
    DataPolicyId string `xml:"DataPolicyId"`
    Id string `xml:"Id"`
    ReplicateDeletes bool `xml:"ReplicateDeletes"`
    State DataPlacementRuleState `xml:"State"`
    TargetDataPolicy *string `xml:"TargetDataPolicy"`
    TargetId string `xml:"TargetId"`
    Type DataReplicationRuleType `xml:"Type"`
}