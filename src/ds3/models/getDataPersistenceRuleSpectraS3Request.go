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

type GetDataPersistenceRuleSpectraS3Request struct {
    dataPersistenceRule string
    queryParams *url.Values
}

func NewGetDataPersistenceRuleSpectraS3Request(dataPersistenceRule string) *GetDataPersistenceRuleSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDataPersistenceRuleSpectraS3Request{
        dataPersistenceRule: dataPersistenceRule,
        queryParams: queryParams,
    }
}




func (GetDataPersistenceRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDataPersistenceRuleSpectraS3Request *GetDataPersistenceRuleSpectraS3Request) Path() string {
    return "/_rest_/data_persistence_rule/" + getDataPersistenceRuleSpectraS3Request.dataPersistenceRule
}

func (getDataPersistenceRuleSpectraS3Request *GetDataPersistenceRuleSpectraS3Request) QueryParams() *url.Values {
    return getDataPersistenceRuleSpectraS3Request.queryParams
}

func (GetDataPersistenceRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDataPersistenceRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDataPersistenceRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
