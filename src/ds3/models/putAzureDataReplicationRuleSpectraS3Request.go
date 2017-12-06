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

type PutAzureDataReplicationRuleSpectraS3Request struct {
    DataPolicyId string
    DataReplicationRuleType DataReplicationRuleType
    MaxBlobPartSizeInBytes *int64
    ReplicateDeletes *bool
    TargetId string
}

func NewPutAzureDataReplicationRuleSpectraS3Request(dataPolicyId string, dataReplicationRuleType DataReplicationRuleType, targetId string) *PutAzureDataReplicationRuleSpectraS3Request {
    return &PutAzureDataReplicationRuleSpectraS3Request{
        DataPolicyId: dataPolicyId,
        TargetId: targetId,
        DataReplicationRuleType: dataReplicationRuleType,
    }
}

func (putAzureDataReplicationRuleSpectraS3Request *PutAzureDataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *PutAzureDataReplicationRuleSpectraS3Request {
    putAzureDataReplicationRuleSpectraS3Request.MaxBlobPartSizeInBytes = &maxBlobPartSizeInBytes
    return putAzureDataReplicationRuleSpectraS3Request
}

func (putAzureDataReplicationRuleSpectraS3Request *PutAzureDataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *PutAzureDataReplicationRuleSpectraS3Request {
    putAzureDataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return putAzureDataReplicationRuleSpectraS3Request
}

