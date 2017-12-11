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

type ModifyS3DataReplicationRuleSpectraS3Request struct {
    DataReplicationRuleType DataReplicationRuleType
    InitialDataPlacement S3InitialDataPlacementPolicy
    MaxBlobPartSizeInBytes *int64
    ReplicateDeletes *bool
    S3DataReplicationRule string
}

func NewModifyS3DataReplicationRuleSpectraS3Request(s3DataReplicationRule string) *ModifyS3DataReplicationRuleSpectraS3Request {
    return &ModifyS3DataReplicationRuleSpectraS3Request{
        S3DataReplicationRule: s3DataReplicationRule,
    }
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithInitialDataPlacement(initialDataPlacement S3InitialDataPlacementPolicy) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.InitialDataPlacement = initialDataPlacement
    return modifyS3DataReplicationRuleSpectraS3Request
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.MaxBlobPartSizeInBytes = &maxBlobPartSizeInBytes
    return modifyS3DataReplicationRuleSpectraS3Request
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return modifyS3DataReplicationRuleSpectraS3Request
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return modifyS3DataReplicationRuleSpectraS3Request
}

