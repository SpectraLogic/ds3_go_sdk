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

type ModifyDs3DataReplicationRuleSpectraS3Request struct {
    DataReplicationRuleType DataReplicationRuleType
    Ds3DataReplicationRule string
    ReplicateDeletes *bool
    TargetDataPolicy *string
}

func NewModifyDs3DataReplicationRuleSpectraS3Request(ds3DataReplicationRule string) *ModifyDs3DataReplicationRuleSpectraS3Request {
    return &ModifyDs3DataReplicationRuleSpectraS3Request{
        Ds3DataReplicationRule: ds3DataReplicationRule,
    }
}

func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *ModifyDs3DataReplicationRuleSpectraS3Request {
    modifyDs3DataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return modifyDs3DataReplicationRuleSpectraS3Request
}

func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) WithTargetDataPolicy(targetDataPolicy string) *ModifyDs3DataReplicationRuleSpectraS3Request {
    modifyDs3DataReplicationRuleSpectraS3Request.TargetDataPolicy = &targetDataPolicy
    return modifyDs3DataReplicationRuleSpectraS3Request
}

func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *ModifyDs3DataReplicationRuleSpectraS3Request {
    modifyDs3DataReplicationRuleSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return modifyDs3DataReplicationRuleSpectraS3Request
}

