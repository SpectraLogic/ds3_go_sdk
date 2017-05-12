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

type PutAzureDataReplicationRuleSpectraS3Request struct {
    dataPolicyId string
    dataReplicationRuleType DataReplicationRuleType
    maxBlobPartSizeInBytes int64
    replicateDeletes bool
    targetId string
    queryParams *url.Values
}

func NewPutAzureDataReplicationRuleSpectraS3Request(dataPolicyId string, dataReplicationRuleType DataReplicationRuleType, targetId string) *PutAzureDataReplicationRuleSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("data_policy_id", dataPolicyId)
    queryParams.Set("target_id", targetId)
    queryParams.Set("type", dataReplicationRuleType.String())

    return &PutAzureDataReplicationRuleSpectraS3Request{
        dataPolicyId: dataPolicyId,
        targetId: targetId,
        dataReplicationRuleType: dataReplicationRuleType,
        queryParams: queryParams,
    }
}

func (putAzureDataReplicationRuleSpectraS3Request *PutAzureDataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *PutAzureDataReplicationRuleSpectraS3Request {
    putAzureDataReplicationRuleSpectraS3Request.maxBlobPartSizeInBytes = maxBlobPartSizeInBytes
    putAzureDataReplicationRuleSpectraS3Request.queryParams.Set("max_blob_part_size_in_bytes", strconv.FormatInt(maxBlobPartSizeInBytes, 10))
    return putAzureDataReplicationRuleSpectraS3Request
}
func (putAzureDataReplicationRuleSpectraS3Request *PutAzureDataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *PutAzureDataReplicationRuleSpectraS3Request {
    putAzureDataReplicationRuleSpectraS3Request.replicateDeletes = replicateDeletes
    putAzureDataReplicationRuleSpectraS3Request.queryParams.Set("replicate_deletes", strconv.FormatBool(replicateDeletes))
    return putAzureDataReplicationRuleSpectraS3Request
}



func (PutAzureDataReplicationRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putAzureDataReplicationRuleSpectraS3Request *PutAzureDataReplicationRuleSpectraS3Request) Path() string {
    return "/_rest_/azure_data_replication_rule"
}

func (putAzureDataReplicationRuleSpectraS3Request *PutAzureDataReplicationRuleSpectraS3Request) QueryParams() *url.Values {
    return putAzureDataReplicationRuleSpectraS3Request.queryParams
}

func (PutAzureDataReplicationRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutAzureDataReplicationRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutAzureDataReplicationRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
