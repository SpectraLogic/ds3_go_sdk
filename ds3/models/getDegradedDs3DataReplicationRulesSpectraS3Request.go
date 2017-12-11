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

type GetDegradedDs3DataReplicationRulesSpectraS3Request struct {
    DataPolicyId *string
    DataReplicationRuleType DataReplicationRuleType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    State DataPlacementRuleState
    TargetId *string
}

func NewGetDegradedDs3DataReplicationRulesSpectraS3Request() *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    return &GetDegradedDs3DataReplicationRulesSpectraS3Request{
    }
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithLastPage() *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.LastPage = true
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.State = state
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

