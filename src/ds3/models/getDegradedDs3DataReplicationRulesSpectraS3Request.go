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

type GetDegradedDs3DataReplicationRulesSpectraS3Request struct {
    dataPolicyId string
    dataReplicationRuleType DataReplicationRuleType
    pageLength int
    pageOffset int
    pageStartMarker string
    state DataPlacementRuleState
    targetId string
    queryParams *url.Values
}

func NewGetDegradedDs3DataReplicationRulesSpectraS3Request() *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDegradedDs3DataReplicationRulesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.dataPolicyId = dataPolicyId
    getDegradedDs3DataReplicationRulesSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}
func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.pageLength = pageLength
    getDegradedDs3DataReplicationRulesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}
func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.pageOffset = pageOffset
    getDegradedDs3DataReplicationRulesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}
func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.pageStartMarker = pageStartMarker
    getDegradedDs3DataReplicationRulesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}
func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.state = state
    getDegradedDs3DataReplicationRulesSpectraS3Request.queryParams.Set("state", state.String())
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}
func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.targetId = targetId
    getDegradedDs3DataReplicationRulesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}
func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.dataReplicationRuleType = dataReplicationRuleType
    getDegradedDs3DataReplicationRulesSpectraS3Request.queryParams.Set("type", dataReplicationRuleType.String())
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}


func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithLastPage() *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.queryParams.Set("last_page", "")
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (GetDegradedDs3DataReplicationRulesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) Path() string {
    return "/_rest_/degraded_ds3_data_replication_rule"
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) QueryParams() *url.Values {
    return getDegradedDs3DataReplicationRulesSpectraS3Request.queryParams
}

func (GetDegradedDs3DataReplicationRulesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDegradedDs3DataReplicationRulesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDegradedDs3DataReplicationRulesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
