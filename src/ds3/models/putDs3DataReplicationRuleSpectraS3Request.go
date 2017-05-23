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

type PutDs3DataReplicationRuleSpectraS3Request struct {
    dataPolicyId string
    dataReplicationRuleType DataReplicationRuleType
    replicateDeletes bool
    targetDataPolicy *string
    targetId string
    queryParams *url.Values
}

func NewPutDs3DataReplicationRuleSpectraS3Request(dataPolicyId string, dataReplicationRuleType DataReplicationRuleType, targetId string) *PutDs3DataReplicationRuleSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("data_policy_id", dataPolicyId)
    queryParams.Set("target_id", targetId)
    queryParams.Set("type", dataReplicationRuleType.String())

    return &PutDs3DataReplicationRuleSpectraS3Request{
        dataPolicyId: dataPolicyId,
        targetId: targetId,
        dataReplicationRuleType: dataReplicationRuleType,
        queryParams: queryParams,
    }
}

func (putDs3DataReplicationRuleSpectraS3Request *PutDs3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *PutDs3DataReplicationRuleSpectraS3Request {
    putDs3DataReplicationRuleSpectraS3Request.replicateDeletes = replicateDeletes
    putDs3DataReplicationRuleSpectraS3Request.queryParams.Set("replicate_deletes", strconv.FormatBool(replicateDeletes))
    return putDs3DataReplicationRuleSpectraS3Request
}

func (putDs3DataReplicationRuleSpectraS3Request *PutDs3DataReplicationRuleSpectraS3Request) WithTargetDataPolicy(targetDataPolicy *string) *PutDs3DataReplicationRuleSpectraS3Request {
    putDs3DataReplicationRuleSpectraS3Request.targetDataPolicy = targetDataPolicy
    if targetDataPolicy != nil {
        putDs3DataReplicationRuleSpectraS3Request.queryParams.Set("target_data_policy", *targetDataPolicy)
    } else {
        putDs3DataReplicationRuleSpectraS3Request.queryParams.Set("target_data_policy", "")
    }
    return putDs3DataReplicationRuleSpectraS3Request
}


func (PutDs3DataReplicationRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putDs3DataReplicationRuleSpectraS3Request *PutDs3DataReplicationRuleSpectraS3Request) Path() string {
    return "/_rest_/ds3_data_replication_rule"
}

func (putDs3DataReplicationRuleSpectraS3Request *PutDs3DataReplicationRuleSpectraS3Request) QueryParams() *url.Values {
    return putDs3DataReplicationRuleSpectraS3Request.queryParams
}

func (PutDs3DataReplicationRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutDs3DataReplicationRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutDs3DataReplicationRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
