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
    "net/url"
    "net/http"
    "ds3/networking"
    "strconv"
)

type GetDataPersistenceRulesSpectraS3Request struct {
    dataPersistenceRuleType DataPersistenceRuleType
    dataPolicyId string
    isolationLevel DataIsolationLevel
    pageLength int
    pageOffset int
    pageStartMarker string
    state DataPlacementRuleState
    storageDomainId string
    queryParams *url.Values
}

func NewGetDataPersistenceRulesSpectraS3Request() *GetDataPersistenceRulesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDataPersistenceRulesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.dataPolicyId = dataPolicyId
    getDataPersistenceRulesSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getDataPersistenceRulesSpectraS3Request
}
func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithIsolationLevel(isolationLevel DataIsolationLevel) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.isolationLevel = isolationLevel
    getDataPersistenceRulesSpectraS3Request.queryParams.Set("isolation_level", isolationLevel.String())
    return getDataPersistenceRulesSpectraS3Request
}
func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithPageLength(pageLength int) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.pageLength = pageLength
    getDataPersistenceRulesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDataPersistenceRulesSpectraS3Request
}
func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.pageOffset = pageOffset
    getDataPersistenceRulesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDataPersistenceRulesSpectraS3Request
}
func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.pageStartMarker = pageStartMarker
    getDataPersistenceRulesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDataPersistenceRulesSpectraS3Request
}
func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.state = state
    getDataPersistenceRulesSpectraS3Request.queryParams.Set("state", state.String())
    return getDataPersistenceRulesSpectraS3Request
}
func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.storageDomainId = storageDomainId
    getDataPersistenceRulesSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return getDataPersistenceRulesSpectraS3Request
}
func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithDataPersistenceRuleType(dataPersistenceRuleType DataPersistenceRuleType) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.dataPersistenceRuleType = dataPersistenceRuleType
    getDataPersistenceRulesSpectraS3Request.queryParams.Set("type", dataPersistenceRuleType.String())
    return getDataPersistenceRulesSpectraS3Request
}


func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithLastPage() *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.queryParams.Set("last_page", "")
    return getDataPersistenceRulesSpectraS3Request
}

func (GetDataPersistenceRulesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) Path() string {
    return "/_rest_/data_persistence_rule"
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) QueryParams() *url.Values {
    return getDataPersistenceRulesSpectraS3Request.queryParams
}

func (GetDataPersistenceRulesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDataPersistenceRulesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDataPersistenceRulesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
