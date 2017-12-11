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

type PutDataPersistenceRuleSpectraS3Request struct {
    DataPersistenceRuleType DataPersistenceRuleType
    DataPolicyId string
    IsolationLevel DataIsolationLevel
    MinimumDaysToRetain *int
    StorageDomainId string
}

func NewPutDataPersistenceRuleSpectraS3Request(dataPersistenceRuleType DataPersistenceRuleType, dataPolicyId string, isolationLevel DataIsolationLevel, storageDomainId string) *PutDataPersistenceRuleSpectraS3Request {
    return &PutDataPersistenceRuleSpectraS3Request{
        DataPolicyId: dataPolicyId,
        IsolationLevel: isolationLevel,
        StorageDomainId: storageDomainId,
        DataPersistenceRuleType: dataPersistenceRuleType,
    }
}

func (putDataPersistenceRuleSpectraS3Request *PutDataPersistenceRuleSpectraS3Request) WithMinimumDaysToRetain(minimumDaysToRetain int) *PutDataPersistenceRuleSpectraS3Request {
    putDataPersistenceRuleSpectraS3Request.MinimumDaysToRetain = &minimumDaysToRetain
    return putDataPersistenceRuleSpectraS3Request
}

