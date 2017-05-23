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

type PutS3DataReplicationRuleSpectraS3Request struct {
    dataPolicyId string
    dataReplicationRuleType DataReplicationRuleType
    initialDataPlacement S3InitialDataPlacementPolicy
    maxBlobPartSizeInBytes int64
    replicateDeletes bool
    targetId string
    queryParams *url.Values
}

func NewPutS3DataReplicationRuleSpectraS3Request(dataPolicyId string, dataReplicationRuleType DataReplicationRuleType, targetId string) *PutS3DataReplicationRuleSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("data_policy_id", dataPolicyId)
    queryParams.Set("target_id", targetId)
    queryParams.Set("type", dataReplicationRuleType.String())

    return &PutS3DataReplicationRuleSpectraS3Request{
        dataPolicyId: dataPolicyId,
        targetId: targetId,
        dataReplicationRuleType: dataReplicationRuleType,
        queryParams: queryParams,
    }
}

func (putS3DataReplicationRuleSpectraS3Request *PutS3DataReplicationRuleSpectraS3Request) WithInitialDataPlacement(initialDataPlacement S3InitialDataPlacementPolicy) *PutS3DataReplicationRuleSpectraS3Request {
    putS3DataReplicationRuleSpectraS3Request.initialDataPlacement = initialDataPlacement
    putS3DataReplicationRuleSpectraS3Request.queryParams.Set("initial_data_placement", initialDataPlacement.String())
    return putS3DataReplicationRuleSpectraS3Request
}
func (putS3DataReplicationRuleSpectraS3Request *PutS3DataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *PutS3DataReplicationRuleSpectraS3Request {
    putS3DataReplicationRuleSpectraS3Request.maxBlobPartSizeInBytes = maxBlobPartSizeInBytes
    putS3DataReplicationRuleSpectraS3Request.queryParams.Set("max_blob_part_size_in_bytes", strconv.FormatInt(maxBlobPartSizeInBytes, 10))
    return putS3DataReplicationRuleSpectraS3Request
}
func (putS3DataReplicationRuleSpectraS3Request *PutS3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *PutS3DataReplicationRuleSpectraS3Request {
    putS3DataReplicationRuleSpectraS3Request.replicateDeletes = replicateDeletes
    putS3DataReplicationRuleSpectraS3Request.queryParams.Set("replicate_deletes", strconv.FormatBool(replicateDeletes))
    return putS3DataReplicationRuleSpectraS3Request
}



func (PutS3DataReplicationRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putS3DataReplicationRuleSpectraS3Request *PutS3DataReplicationRuleSpectraS3Request) Path() string {
    return "/_rest_/s3_data_replication_rule"
}

func (putS3DataReplicationRuleSpectraS3Request *PutS3DataReplicationRuleSpectraS3Request) QueryParams() *url.Values {
    return putS3DataReplicationRuleSpectraS3Request.queryParams
}

func (PutS3DataReplicationRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutS3DataReplicationRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutS3DataReplicationRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
