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

type GetDegradedAzureDataReplicationRulesSpectraS3Request struct {
    dataPolicyId string
    dataReplicationRuleType DataReplicationRuleType
    pageLength int
    pageOffset int
    pageStartMarker string
    state DataPlacementRuleState
    targetId string
    queryParams *url.Values
}

func NewGetDegradedAzureDataReplicationRulesSpectraS3Request() *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDegradedAzureDataReplicationRulesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.dataPolicyId = dataPolicyId
    getDegradedAzureDataReplicationRulesSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}
func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.pageLength = pageLength
    getDegradedAzureDataReplicationRulesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}
func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.pageOffset = pageOffset
    getDegradedAzureDataReplicationRulesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}
func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.pageStartMarker = pageStartMarker
    getDegradedAzureDataReplicationRulesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}
func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.state = state
    getDegradedAzureDataReplicationRulesSpectraS3Request.queryParams.Set("state", state.String())
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}
func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.targetId = targetId
    getDegradedAzureDataReplicationRulesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}
func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.dataReplicationRuleType = dataReplicationRuleType
    getDegradedAzureDataReplicationRulesSpectraS3Request.queryParams.Set("type", dataReplicationRuleType.String())
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}


func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithLastPage() *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.queryParams.Set("last_page", "")
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (GetDegradedAzureDataReplicationRulesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) Path() string {
    return "/_rest_/degraded_azure_data_replication_rule"
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) QueryParams() *url.Values {
    return getDegradedAzureDataReplicationRulesSpectraS3Request.queryParams
}

func (GetDegradedAzureDataReplicationRulesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDegradedAzureDataReplicationRulesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDegradedAzureDataReplicationRulesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
