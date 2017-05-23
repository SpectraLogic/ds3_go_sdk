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

type ModifyDataPersistenceRuleSpectraS3Request struct {
    dataPersistenceRuleId string
    dataPersistenceRuleType DataPersistenceRuleType
    isolationLevel DataIsolationLevel
    minimumDaysToRetain *int
    queryParams *url.Values
}

func NewModifyDataPersistenceRuleSpectraS3Request(dataPersistenceRuleId string) *ModifyDataPersistenceRuleSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyDataPersistenceRuleSpectraS3Request{
        dataPersistenceRuleId: dataPersistenceRuleId,
        queryParams: queryParams,
    }
}

func (modifyDataPersistenceRuleSpectraS3Request *ModifyDataPersistenceRuleSpectraS3Request) WithIsolationLevel(isolationLevel DataIsolationLevel) *ModifyDataPersistenceRuleSpectraS3Request {
    modifyDataPersistenceRuleSpectraS3Request.isolationLevel = isolationLevel
    modifyDataPersistenceRuleSpectraS3Request.queryParams.Set("isolation_level", isolationLevel.String())
    return modifyDataPersistenceRuleSpectraS3Request
}
func (modifyDataPersistenceRuleSpectraS3Request *ModifyDataPersistenceRuleSpectraS3Request) WithDataPersistenceRuleType(dataPersistenceRuleType DataPersistenceRuleType) *ModifyDataPersistenceRuleSpectraS3Request {
    modifyDataPersistenceRuleSpectraS3Request.dataPersistenceRuleType = dataPersistenceRuleType
    modifyDataPersistenceRuleSpectraS3Request.queryParams.Set("type", dataPersistenceRuleType.String())
    return modifyDataPersistenceRuleSpectraS3Request
}

func (modifyDataPersistenceRuleSpectraS3Request *ModifyDataPersistenceRuleSpectraS3Request) WithMinimumDaysToRetain(minimumDaysToRetain *int) *ModifyDataPersistenceRuleSpectraS3Request {
    modifyDataPersistenceRuleSpectraS3Request.minimumDaysToRetain = minimumDaysToRetain
    if minimumDaysToRetain != nil {
        modifyDataPersistenceRuleSpectraS3Request.queryParams.Set("minimum_days_to_retain", strconv.Itoa(*minimumDaysToRetain))
    } else {
        modifyDataPersistenceRuleSpectraS3Request.queryParams.Set("minimum_days_to_retain", "")
    }
    return modifyDataPersistenceRuleSpectraS3Request
}


func (ModifyDataPersistenceRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyDataPersistenceRuleSpectraS3Request *ModifyDataPersistenceRuleSpectraS3Request) Path() string {
    return "/_rest_/data_persistence_rule/" + modifyDataPersistenceRuleSpectraS3Request.dataPersistenceRuleId
}

func (modifyDataPersistenceRuleSpectraS3Request *ModifyDataPersistenceRuleSpectraS3Request) QueryParams() *url.Values {
    return modifyDataPersistenceRuleSpectraS3Request.queryParams
}

func (ModifyDataPersistenceRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyDataPersistenceRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyDataPersistenceRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
