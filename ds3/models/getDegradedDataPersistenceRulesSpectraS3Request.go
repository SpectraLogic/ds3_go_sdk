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

type GetDegradedDataPersistenceRulesSpectraS3Request struct {
    DataPersistenceRuleType DataPersistenceRuleType
    DataPolicyId *string
    IsolationLevel DataIsolationLevel
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    State DataPlacementRuleState
    StorageDomainId *string
}

func NewGetDegradedDataPersistenceRulesSpectraS3Request() *GetDegradedDataPersistenceRulesSpectraS3Request {
    return &GetDegradedDataPersistenceRulesSpectraS3Request{
    }
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithIsolationLevel(isolationLevel DataIsolationLevel) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.IsolationLevel = isolationLevel
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithLastPage() *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.LastPage = true
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.PageLength = &pageLength
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.PageOffset = &pageOffset
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.State = state
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.StorageDomainId = &storageDomainId
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithDataPersistenceRuleType(dataPersistenceRuleType DataPersistenceRuleType) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.DataPersistenceRuleType = dataPersistenceRuleType
    return getDegradedDataPersistenceRulesSpectraS3Request
}

