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

type GetAzureDataReplicationRulesSpectraS3Request struct {
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

func NewGetAzureDataReplicationRulesSpectraS3Request() *GetAzureDataReplicationRulesSpectraS3Request {
    return &GetAzureDataReplicationRulesSpectraS3Request{
    }
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithLastPage() *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.LastPage = true
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.State = state
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getAzureDataReplicationRulesSpectraS3Request
}

