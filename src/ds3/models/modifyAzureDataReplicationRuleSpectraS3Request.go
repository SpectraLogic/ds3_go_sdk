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

type ModifyAzureDataReplicationRuleSpectraS3Request struct {
    azureDataReplicationRule string
    dataReplicationRuleType DataReplicationRuleType
    maxBlobPartSizeInBytes int64
    replicateDeletes bool
    queryParams *url.Values
}

func NewModifyAzureDataReplicationRuleSpectraS3Request(azureDataReplicationRule string) *ModifyAzureDataReplicationRuleSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyAzureDataReplicationRuleSpectraS3Request{
        azureDataReplicationRule: azureDataReplicationRule,
        queryParams: queryParams,
    }
}

func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *ModifyAzureDataReplicationRuleSpectraS3Request {
    modifyAzureDataReplicationRuleSpectraS3Request.maxBlobPartSizeInBytes = maxBlobPartSizeInBytes
    modifyAzureDataReplicationRuleSpectraS3Request.queryParams.Set("max_blob_part_size_in_bytes", strconv.FormatInt(maxBlobPartSizeInBytes, 10))
    return modifyAzureDataReplicationRuleSpectraS3Request
}
func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *ModifyAzureDataReplicationRuleSpectraS3Request {
    modifyAzureDataReplicationRuleSpectraS3Request.replicateDeletes = replicateDeletes
    modifyAzureDataReplicationRuleSpectraS3Request.queryParams.Set("replicate_deletes", strconv.FormatBool(replicateDeletes))
    return modifyAzureDataReplicationRuleSpectraS3Request
}
func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *ModifyAzureDataReplicationRuleSpectraS3Request {
    modifyAzureDataReplicationRuleSpectraS3Request.dataReplicationRuleType = dataReplicationRuleType
    modifyAzureDataReplicationRuleSpectraS3Request.queryParams.Set("type", dataReplicationRuleType.String())
    return modifyAzureDataReplicationRuleSpectraS3Request
}



func (ModifyAzureDataReplicationRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) Path() string {
    return "/_rest_/azure_data_replication_rule/" + modifyAzureDataReplicationRuleSpectraS3Request.azureDataReplicationRule
}

func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) QueryParams() *url.Values {
    return modifyAzureDataReplicationRuleSpectraS3Request.queryParams
}

func (ModifyAzureDataReplicationRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyAzureDataReplicationRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyAzureDataReplicationRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
