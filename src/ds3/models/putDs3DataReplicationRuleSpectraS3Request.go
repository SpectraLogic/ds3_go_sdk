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

type PutDs3DataReplicationRuleSpectraS3Request struct {
    DataPolicyId string
    DataReplicationRuleType DataReplicationRuleType
    ReplicateDeletes *bool
    TargetDataPolicy *string
    TargetId string
}

func NewPutDs3DataReplicationRuleSpectraS3Request(dataPolicyId string, dataReplicationRuleType DataReplicationRuleType, targetId string) *PutDs3DataReplicationRuleSpectraS3Request {
    return &PutDs3DataReplicationRuleSpectraS3Request{
        DataPolicyId: dataPolicyId,
        TargetId: targetId,
        DataReplicationRuleType: dataReplicationRuleType,
    }
}

func (putDs3DataReplicationRuleSpectraS3Request *PutDs3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *PutDs3DataReplicationRuleSpectraS3Request {
    putDs3DataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return putDs3DataReplicationRuleSpectraS3Request
}

func (putDs3DataReplicationRuleSpectraS3Request *PutDs3DataReplicationRuleSpectraS3Request) WithTargetDataPolicy(targetDataPolicy string) *PutDs3DataReplicationRuleSpectraS3Request {
    putDs3DataReplicationRuleSpectraS3Request.TargetDataPolicy = &targetDataPolicy
    return putDs3DataReplicationRuleSpectraS3Request
}

