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

type GetDegradedDataPersistenceRulesSpectraS3Request struct {
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

func NewGetDegradedDataPersistenceRulesSpectraS3Request() *GetDegradedDataPersistenceRulesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDegradedDataPersistenceRulesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.dataPolicyId = dataPolicyId
    getDegradedDataPersistenceRulesSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getDegradedDataPersistenceRulesSpectraS3Request
}
func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithIsolationLevel(isolationLevel DataIsolationLevel) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.isolationLevel = isolationLevel
    getDegradedDataPersistenceRulesSpectraS3Request.queryParams.Set("isolation_level", isolationLevel.String())
    return getDegradedDataPersistenceRulesSpectraS3Request
}
func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.pageLength = pageLength
    getDegradedDataPersistenceRulesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDegradedDataPersistenceRulesSpectraS3Request
}
func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.pageOffset = pageOffset
    getDegradedDataPersistenceRulesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDegradedDataPersistenceRulesSpectraS3Request
}
func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.pageStartMarker = pageStartMarker
    getDegradedDataPersistenceRulesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDegradedDataPersistenceRulesSpectraS3Request
}
func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.state = state
    getDegradedDataPersistenceRulesSpectraS3Request.queryParams.Set("state", state.String())
    return getDegradedDataPersistenceRulesSpectraS3Request
}
func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.storageDomainId = storageDomainId
    getDegradedDataPersistenceRulesSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return getDegradedDataPersistenceRulesSpectraS3Request
}
func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithDataPersistenceRuleType(dataPersistenceRuleType DataPersistenceRuleType) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.dataPersistenceRuleType = dataPersistenceRuleType
    getDegradedDataPersistenceRulesSpectraS3Request.queryParams.Set("type", dataPersistenceRuleType.String())
    return getDegradedDataPersistenceRulesSpectraS3Request
}


func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithLastPage() *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.queryParams.Set("last_page", "")
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (GetDegradedDataPersistenceRulesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) Path() string {
    return "/_rest_/degraded_data_persistence_rule"
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) QueryParams() *url.Values {
    return getDegradedDataPersistenceRulesSpectraS3Request.queryParams
}

func (GetDegradedDataPersistenceRulesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDegradedDataPersistenceRulesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDegradedDataPersistenceRulesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
