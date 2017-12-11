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

type GetS3DataReplicationRulesSpectraS3Request struct {
    DataPolicyId *string
    DataReplicationRuleType DataReplicationRuleType
    InitialDataPlacement S3InitialDataPlacementPolicy
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    ReplicateDeletes *bool
    State DataPlacementRuleState
    TargetId *string
}

func NewGetS3DataReplicationRulesSpectraS3Request() *GetS3DataReplicationRulesSpectraS3Request {
    return &GetS3DataReplicationRulesSpectraS3Request{
    }
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithInitialDataPlacement(initialDataPlacement S3InitialDataPlacementPolicy) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.InitialDataPlacement = initialDataPlacement
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithLastPage() *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.LastPage = true
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.State = state
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getS3DataReplicationRulesSpectraS3Request
}

