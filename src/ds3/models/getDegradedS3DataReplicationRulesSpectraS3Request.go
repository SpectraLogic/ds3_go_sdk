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

type GetDegradedS3DataReplicationRulesSpectraS3Request struct {
    dataPolicyId string
    dataReplicationRuleType DataReplicationRuleType
    pageLength int
    pageOffset int
    pageStartMarker string
    state DataPlacementRuleState
    targetId string
    queryParams *url.Values
}

func NewGetDegradedS3DataReplicationRulesSpectraS3Request() *GetDegradedS3DataReplicationRulesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDegradedS3DataReplicationRulesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.dataPolicyId = dataPolicyId
    getDegradedS3DataReplicationRulesSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getDegradedS3DataReplicationRulesSpectraS3Request
}
func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.pageLength = pageLength
    getDegradedS3DataReplicationRulesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDegradedS3DataReplicationRulesSpectraS3Request
}
func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.pageOffset = pageOffset
    getDegradedS3DataReplicationRulesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDegradedS3DataReplicationRulesSpectraS3Request
}
func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.pageStartMarker = pageStartMarker
    getDegradedS3DataReplicationRulesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDegradedS3DataReplicationRulesSpectraS3Request
}
func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.state = state
    getDegradedS3DataReplicationRulesSpectraS3Request.queryParams.Set("state", state.String())
    return getDegradedS3DataReplicationRulesSpectraS3Request
}
func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.targetId = targetId
    getDegradedS3DataReplicationRulesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getDegradedS3DataReplicationRulesSpectraS3Request
}
func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.dataReplicationRuleType = dataReplicationRuleType
    getDegradedS3DataReplicationRulesSpectraS3Request.queryParams.Set("type", dataReplicationRuleType.String())
    return getDegradedS3DataReplicationRulesSpectraS3Request
}


func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithLastPage() *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.queryParams.Set("last_page", "")
    return getDegradedS3DataReplicationRulesSpectraS3Request
}

func (GetDegradedS3DataReplicationRulesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) Path() string {
    return "/_rest_/degraded_s3_data_replication_rule"
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) QueryParams() *url.Values {
    return getDegradedS3DataReplicationRulesSpectraS3Request.queryParams
}

func (GetDegradedS3DataReplicationRulesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDegradedS3DataReplicationRulesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDegradedS3DataReplicationRulesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
