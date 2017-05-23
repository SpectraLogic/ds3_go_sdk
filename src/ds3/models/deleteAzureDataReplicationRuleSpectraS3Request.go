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

type DeleteAzureDataReplicationRuleSpectraS3Request struct {
    azureDataReplicationRule string
    queryParams *url.Values
}

func NewDeleteAzureDataReplicationRuleSpectraS3Request(azureDataReplicationRule string) *DeleteAzureDataReplicationRuleSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteAzureDataReplicationRuleSpectraS3Request{
        azureDataReplicationRule: azureDataReplicationRule,
        queryParams: queryParams,
    }
}




func (DeleteAzureDataReplicationRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteAzureDataReplicationRuleSpectraS3Request *DeleteAzureDataReplicationRuleSpectraS3Request) Path() string {
    return "/_rest_/azure_data_replication_rule/" + deleteAzureDataReplicationRuleSpectraS3Request.azureDataReplicationRule
}

func (deleteAzureDataReplicationRuleSpectraS3Request *DeleteAzureDataReplicationRuleSpectraS3Request) QueryParams() *url.Values {
    return deleteAzureDataReplicationRuleSpectraS3Request.queryParams
}

func (DeleteAzureDataReplicationRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteAzureDataReplicationRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteAzureDataReplicationRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
