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

type ModifyDs3DataReplicationRuleSpectraS3Request struct {
    dataReplicationRuleType DataReplicationRuleType
    ds3DataReplicationRule string
    replicateDeletes bool
    targetDataPolicy *string
    queryParams *url.Values
}

func NewModifyDs3DataReplicationRuleSpectraS3Request(ds3DataReplicationRule string) *ModifyDs3DataReplicationRuleSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyDs3DataReplicationRuleSpectraS3Request{
        ds3DataReplicationRule: ds3DataReplicationRule,
        queryParams: queryParams,
    }
}

func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *ModifyDs3DataReplicationRuleSpectraS3Request {
    modifyDs3DataReplicationRuleSpectraS3Request.replicateDeletes = replicateDeletes
    modifyDs3DataReplicationRuleSpectraS3Request.queryParams.Set("replicate_deletes", strconv.FormatBool(replicateDeletes))
    return modifyDs3DataReplicationRuleSpectraS3Request
}
func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *ModifyDs3DataReplicationRuleSpectraS3Request {
    modifyDs3DataReplicationRuleSpectraS3Request.dataReplicationRuleType = dataReplicationRuleType
    modifyDs3DataReplicationRuleSpectraS3Request.queryParams.Set("type", dataReplicationRuleType.String())
    return modifyDs3DataReplicationRuleSpectraS3Request
}

func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) WithTargetDataPolicy(targetDataPolicy *string) *ModifyDs3DataReplicationRuleSpectraS3Request {
    modifyDs3DataReplicationRuleSpectraS3Request.targetDataPolicy = targetDataPolicy
    if targetDataPolicy != nil {
        modifyDs3DataReplicationRuleSpectraS3Request.queryParams.Set("target_data_policy", *targetDataPolicy)
    } else {
        modifyDs3DataReplicationRuleSpectraS3Request.queryParams.Set("target_data_policy", "")
    }
    return modifyDs3DataReplicationRuleSpectraS3Request
}


func (ModifyDs3DataReplicationRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) Path() string {
    return "/_rest_/ds3_data_replication_rule/" + modifyDs3DataReplicationRuleSpectraS3Request.ds3DataReplicationRule
}

func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) QueryParams() *url.Values {
    return modifyDs3DataReplicationRuleSpectraS3Request.queryParams
}

func (ModifyDs3DataReplicationRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyDs3DataReplicationRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyDs3DataReplicationRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
