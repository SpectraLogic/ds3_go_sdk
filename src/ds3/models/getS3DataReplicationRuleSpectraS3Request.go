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
)

type GetS3DataReplicationRuleSpectraS3Request struct {
    s3DataReplicationRule string
    queryParams *url.Values
}

func NewGetS3DataReplicationRuleSpectraS3Request(s3DataReplicationRule string) *GetS3DataReplicationRuleSpectraS3Request {
    queryParams := &url.Values{}

    return &GetS3DataReplicationRuleSpectraS3Request{
        s3DataReplicationRule: s3DataReplicationRule,
        queryParams: queryParams,
    }
}




func (GetS3DataReplicationRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getS3DataReplicationRuleSpectraS3Request *GetS3DataReplicationRuleSpectraS3Request) Path() string {
    return "/_rest_/s3_data_replication_rule/" + getS3DataReplicationRuleSpectraS3Request.s3DataReplicationRule
}

func (getS3DataReplicationRuleSpectraS3Request *GetS3DataReplicationRuleSpectraS3Request) QueryParams() *url.Values {
    return getS3DataReplicationRuleSpectraS3Request.queryParams
}

func (GetS3DataReplicationRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetS3DataReplicationRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetS3DataReplicationRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
