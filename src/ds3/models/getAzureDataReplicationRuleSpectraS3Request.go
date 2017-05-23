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

type GetAzureDataReplicationRuleSpectraS3Request struct {
    azureDataReplicationRule string
    queryParams *url.Values
}

func NewGetAzureDataReplicationRuleSpectraS3Request(azureDataReplicationRule string) *GetAzureDataReplicationRuleSpectraS3Request {
    queryParams := &url.Values{}

    return &GetAzureDataReplicationRuleSpectraS3Request{
        azureDataReplicationRule: azureDataReplicationRule,
        queryParams: queryParams,
    }
}




func (GetAzureDataReplicationRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getAzureDataReplicationRuleSpectraS3Request *GetAzureDataReplicationRuleSpectraS3Request) Path() string {
    return "/_rest_/azure_data_replication_rule/" + getAzureDataReplicationRuleSpectraS3Request.azureDataReplicationRule
}

func (getAzureDataReplicationRuleSpectraS3Request *GetAzureDataReplicationRuleSpectraS3Request) QueryParams() *url.Values {
    return getAzureDataReplicationRuleSpectraS3Request.queryParams
}

func (GetAzureDataReplicationRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetAzureDataReplicationRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetAzureDataReplicationRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
