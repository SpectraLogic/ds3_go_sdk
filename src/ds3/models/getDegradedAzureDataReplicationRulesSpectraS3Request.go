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

type GetDegradedAzureDataReplicationRulesSpectraS3Request struct {
    DataPolicyId *string
    DataReplicationRuleType DataReplicationRuleType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    State DataPlacementRuleState
    TargetId *string
}

func NewGetDegradedAzureDataReplicationRulesSpectraS3Request() *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    return &GetDegradedAzureDataReplicationRulesSpectraS3Request{
    }
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithLastPage() *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.LastPage = true
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.State = state
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

