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

type GetDs3DataReplicationRulesSpectraS3Request struct {
    DataPolicyId *string
    DataReplicationRuleType DataReplicationRuleType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    ReplicateDeletes *bool
    State DataPlacementRuleState
    TargetId *string
}

func NewGetDs3DataReplicationRulesSpectraS3Request() *GetDs3DataReplicationRulesSpectraS3Request {
    return &GetDs3DataReplicationRulesSpectraS3Request{
    }
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithLastPage() *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.LastPage = true
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.State = state
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getDs3DataReplicationRulesSpectraS3Request
}

