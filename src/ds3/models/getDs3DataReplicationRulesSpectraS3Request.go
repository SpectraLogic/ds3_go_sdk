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

type GetDs3DataReplicationRulesSpectraS3Request struct {
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

func NewGetDs3DataReplicationRulesSpectraS3Request() *GetDs3DataReplicationRulesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDs3DataReplicationRulesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.dataPolicyId = dataPolicyId
    getDs3DataReplicationRulesSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getDs3DataReplicationRulesSpectraS3Request
}
func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.pageLength = pageLength
    getDs3DataReplicationRulesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDs3DataReplicationRulesSpectraS3Request
}
func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.pageOffset = pageOffset
    getDs3DataReplicationRulesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDs3DataReplicationRulesSpectraS3Request
}
func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.pageStartMarker = pageStartMarker
    getDs3DataReplicationRulesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDs3DataReplicationRulesSpectraS3Request
}
func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.replicateDeletes = replicateDeletes
    getDs3DataReplicationRulesSpectraS3Request.queryParams.Set("replicate_deletes", strconv.FormatBool(replicateDeletes))
    return getDs3DataReplicationRulesSpectraS3Request
}
func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.state = state
    getDs3DataReplicationRulesSpectraS3Request.queryParams.Set("state", state.String())
    return getDs3DataReplicationRulesSpectraS3Request
}
func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.targetId = targetId
    getDs3DataReplicationRulesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getDs3DataReplicationRulesSpectraS3Request
}
func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.dataReplicationRuleType = dataReplicationRuleType
    getDs3DataReplicationRulesSpectraS3Request.queryParams.Set("type", dataReplicationRuleType.String())
    return getDs3DataReplicationRulesSpectraS3Request
}


func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithLastPage() *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.queryParams.Set("last_page", "")
    return getDs3DataReplicationRulesSpectraS3Request
}

func (GetDs3DataReplicationRulesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) Path() string {
    return "/_rest_/ds3_data_replication_rule"
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) QueryParams() *url.Values {
    return getDs3DataReplicationRulesSpectraS3Request.queryParams
}

func (GetDs3DataReplicationRulesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDs3DataReplicationRulesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDs3DataReplicationRulesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
