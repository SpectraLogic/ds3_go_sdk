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

type GetS3DataReplicationRulesSpectraS3Request struct {
    dataPolicyId string
    dataReplicationRuleType DataReplicationRuleType
    initialDataPlacement S3InitialDataPlacementPolicy
    pageLength int
    pageOffset int
    pageStartMarker string
    replicateDeletes bool
    state DataPlacementRuleState
    targetId string
    queryParams *url.Values
}

func NewGetS3DataReplicationRulesSpectraS3Request() *GetS3DataReplicationRulesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetS3DataReplicationRulesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.dataPolicyId = dataPolicyId
    getS3DataReplicationRulesSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getS3DataReplicationRulesSpectraS3Request
}
func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithInitialDataPlacement(initialDataPlacement S3InitialDataPlacementPolicy) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.initialDataPlacement = initialDataPlacement
    getS3DataReplicationRulesSpectraS3Request.queryParams.Set("initial_data_placement", initialDataPlacement.String())
    return getS3DataReplicationRulesSpectraS3Request
}
func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.pageLength = pageLength
    getS3DataReplicationRulesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getS3DataReplicationRulesSpectraS3Request
}
func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.pageOffset = pageOffset
    getS3DataReplicationRulesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getS3DataReplicationRulesSpectraS3Request
}
func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.pageStartMarker = pageStartMarker
    getS3DataReplicationRulesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getS3DataReplicationRulesSpectraS3Request
}
func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.replicateDeletes = replicateDeletes
    getS3DataReplicationRulesSpectraS3Request.queryParams.Set("replicate_deletes", strconv.FormatBool(replicateDeletes))
    return getS3DataReplicationRulesSpectraS3Request
}
func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.state = state
    getS3DataReplicationRulesSpectraS3Request.queryParams.Set("state", state.String())
    return getS3DataReplicationRulesSpectraS3Request
}
func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.targetId = targetId
    getS3DataReplicationRulesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getS3DataReplicationRulesSpectraS3Request
}
func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.dataReplicationRuleType = dataReplicationRuleType
    getS3DataReplicationRulesSpectraS3Request.queryParams.Set("type", dataReplicationRuleType.String())
    return getS3DataReplicationRulesSpectraS3Request
}


func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithLastPage() *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.queryParams.Set("last_page", "")
    return getS3DataReplicationRulesSpectraS3Request
}

func (GetS3DataReplicationRulesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) Path() string {
    return "/_rest_/s3_data_replication_rule"
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) QueryParams() *url.Values {
    return getS3DataReplicationRulesSpectraS3Request.queryParams
}

func (GetS3DataReplicationRulesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetS3DataReplicationRulesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetS3DataReplicationRulesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
