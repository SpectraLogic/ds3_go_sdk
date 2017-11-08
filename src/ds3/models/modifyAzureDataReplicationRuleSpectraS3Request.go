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

type ModifyAzureDataReplicationRuleSpectraS3Request struct {
    AzureDataReplicationRule string
    DataReplicationRuleType DataReplicationRuleType
    MaxBlobPartSizeInBytes *int64
    ReplicateDeletes *bool
}

func NewModifyAzureDataReplicationRuleSpectraS3Request(azureDataReplicationRule string) *ModifyAzureDataReplicationRuleSpectraS3Request {
    return &ModifyAzureDataReplicationRuleSpectraS3Request{
        AzureDataReplicationRule: azureDataReplicationRule,
    }
}

func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *ModifyAzureDataReplicationRuleSpectraS3Request {
    modifyAzureDataReplicationRuleSpectraS3Request.MaxBlobPartSizeInBytes = &maxBlobPartSizeInBytes
    return modifyAzureDataReplicationRuleSpectraS3Request
}

func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *ModifyAzureDataReplicationRuleSpectraS3Request {
    modifyAzureDataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return modifyAzureDataReplicationRuleSpectraS3Request
}

func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *ModifyAzureDataReplicationRuleSpectraS3Request {
    modifyAzureDataReplicationRuleSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return modifyAzureDataReplicationRuleSpectraS3Request
}

