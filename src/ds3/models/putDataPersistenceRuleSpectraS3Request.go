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

type PutDataPersistenceRuleSpectraS3Request struct {
    dataPersistenceRuleType DataPersistenceRuleType
    dataPolicyId string
    isolationLevel DataIsolationLevel
    minimumDaysToRetain *int
    storageDomainId string
    queryParams *url.Values
}

func NewPutDataPersistenceRuleSpectraS3Request(dataPersistenceRuleType DataPersistenceRuleType, dataPolicyId string, isolationLevel DataIsolationLevel, storageDomainId string) *PutDataPersistenceRuleSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("data_policy_id", dataPolicyId)
    queryParams.Set("isolation_level", isolationLevel.String())
    queryParams.Set("storage_domain_id", storageDomainId)
    queryParams.Set("type", dataPersistenceRuleType.String())

    return &PutDataPersistenceRuleSpectraS3Request{
        dataPolicyId: dataPolicyId,
        isolationLevel: isolationLevel,
        storageDomainId: storageDomainId,
        dataPersistenceRuleType: dataPersistenceRuleType,
        queryParams: queryParams,
    }
}


func (putDataPersistenceRuleSpectraS3Request *PutDataPersistenceRuleSpectraS3Request) WithMinimumDaysToRetain(minimumDaysToRetain *int) *PutDataPersistenceRuleSpectraS3Request {
    putDataPersistenceRuleSpectraS3Request.minimumDaysToRetain = minimumDaysToRetain
    if minimumDaysToRetain != nil {
        putDataPersistenceRuleSpectraS3Request.queryParams.Set("minimum_days_to_retain", strconv.Itoa(*minimumDaysToRetain))
    } else {
        putDataPersistenceRuleSpectraS3Request.queryParams.Set("minimum_days_to_retain", "")
    }
    return putDataPersistenceRuleSpectraS3Request
}


func (PutDataPersistenceRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putDataPersistenceRuleSpectraS3Request *PutDataPersistenceRuleSpectraS3Request) Path() string {
    return "/_rest_/data_persistence_rule"
}

func (putDataPersistenceRuleSpectraS3Request *PutDataPersistenceRuleSpectraS3Request) QueryParams() *url.Values {
    return putDataPersistenceRuleSpectraS3Request.queryParams
}

func (PutDataPersistenceRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutDataPersistenceRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutDataPersistenceRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
