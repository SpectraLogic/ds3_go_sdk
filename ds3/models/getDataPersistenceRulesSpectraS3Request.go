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

type GetDataPersistenceRulesSpectraS3Request struct {
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

func NewGetDataPersistenceRulesSpectraS3Request() *GetDataPersistenceRulesSpectraS3Request {
    return &GetDataPersistenceRulesSpectraS3Request{
    }
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithIsolationLevel(isolationLevel DataIsolationLevel) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.IsolationLevel = isolationLevel
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithLastPage() *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.LastPage = true
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithPageLength(pageLength int) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.PageLength = &pageLength
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.PageOffset = &pageOffset
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.State = state
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.StorageDomainId = &storageDomainId
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithDataPersistenceRuleType(dataPersistenceRuleType DataPersistenceRuleType) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.DataPersistenceRuleType = dataPersistenceRuleType
    return getDataPersistenceRulesSpectraS3Request
}

