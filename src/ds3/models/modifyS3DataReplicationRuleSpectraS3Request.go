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

type ModifyS3DataReplicationRuleSpectraS3Request struct {
    dataReplicationRuleType DataReplicationRuleType
    initialDataPlacement S3InitialDataPlacementPolicy
    maxBlobPartSizeInBytes int64
    replicateDeletes bool
    s3DataReplicationRule string
    queryParams *url.Values
}

func NewModifyS3DataReplicationRuleSpectraS3Request(s3DataReplicationRule string) *ModifyS3DataReplicationRuleSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyS3DataReplicationRuleSpectraS3Request{
        s3DataReplicationRule: s3DataReplicationRule,
        queryParams: queryParams,
    }
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithInitialDataPlacement(initialDataPlacement S3InitialDataPlacementPolicy) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.initialDataPlacement = initialDataPlacement
    modifyS3DataReplicationRuleSpectraS3Request.queryParams.Set("initial_data_placement", initialDataPlacement.String())
    return modifyS3DataReplicationRuleSpectraS3Request
}
func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.maxBlobPartSizeInBytes = maxBlobPartSizeInBytes
    modifyS3DataReplicationRuleSpectraS3Request.queryParams.Set("max_blob_part_size_in_bytes", strconv.FormatInt(maxBlobPartSizeInBytes, 10))
    return modifyS3DataReplicationRuleSpectraS3Request
}
func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.replicateDeletes = replicateDeletes
    modifyS3DataReplicationRuleSpectraS3Request.queryParams.Set("replicate_deletes", strconv.FormatBool(replicateDeletes))
    return modifyS3DataReplicationRuleSpectraS3Request
}
func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.dataReplicationRuleType = dataReplicationRuleType
    modifyS3DataReplicationRuleSpectraS3Request.queryParams.Set("type", dataReplicationRuleType.String())
    return modifyS3DataReplicationRuleSpectraS3Request
}



func (ModifyS3DataReplicationRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) Path() string {
    return "/_rest_/s3_data_replication_rule/" + modifyS3DataReplicationRuleSpectraS3Request.s3DataReplicationRule
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) QueryParams() *url.Values {
    return modifyS3DataReplicationRuleSpectraS3Request.queryParams
}

func (ModifyS3DataReplicationRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyS3DataReplicationRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyS3DataReplicationRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
