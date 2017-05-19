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

type GetAzureDataReplicationRulesSpectraS3Request struct {
    dataPolicyId string
    dataReplicationRuleType DataReplicationRuleType
    pageLength int
    pageOffset int
    pageStartMarker string
    replicateDeletes bool
    state DataPlacementRuleState
    targetId string
    queryParams *url.Values
}

func NewGetAzureDataReplicationRulesSpectraS3Request() *GetAzureDataReplicationRulesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetAzureDataReplicationRulesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.dataPolicyId = dataPolicyId
    getAzureDataReplicationRulesSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getAzureDataReplicationRulesSpectraS3Request
}
func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.pageLength = pageLength
    getAzureDataReplicationRulesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getAzureDataReplicationRulesSpectraS3Request
}
func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.pageOffset = pageOffset
    getAzureDataReplicationRulesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getAzureDataReplicationRulesSpectraS3Request
}
func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.pageStartMarker = pageStartMarker
    getAzureDataReplicationRulesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getAzureDataReplicationRulesSpectraS3Request
}
func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.replicateDeletes = replicateDeletes
    getAzureDataReplicationRulesSpectraS3Request.queryParams.Set("replicate_deletes", strconv.FormatBool(replicateDeletes))
    return getAzureDataReplicationRulesSpectraS3Request
}
func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.state = state
    getAzureDataReplicationRulesSpectraS3Request.queryParams.Set("state", state.String())
    return getAzureDataReplicationRulesSpectraS3Request
}
func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.targetId = targetId
    getAzureDataReplicationRulesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getAzureDataReplicationRulesSpectraS3Request
}
func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.dataReplicationRuleType = dataReplicationRuleType
    getAzureDataReplicationRulesSpectraS3Request.queryParams.Set("type", dataReplicationRuleType.String())
    return getAzureDataReplicationRulesSpectraS3Request
}


func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithLastPage() *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.queryParams.Set("last_page", "")
    return getAzureDataReplicationRulesSpectraS3Request
}

func (GetAzureDataReplicationRulesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) Path() string {
    return "/_rest_/azure_data_replication_rule"
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) QueryParams() *url.Values {
    return getAzureDataReplicationRulesSpectraS3Request.queryParams
}

func (GetAzureDataReplicationRulesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetAzureDataReplicationRulesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetAzureDataReplicationRulesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
