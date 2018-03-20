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

type ImportPoolSpectraS3Request struct {
    DataPolicyId *string
    Pool string
    Priority Priority
    StorageDomainId *string
    UserId *string
    VerifyDataAfterImport Priority
    VerifyDataPriorToImport *bool
}

func NewImportPoolSpectraS3Request(pool string) *ImportPoolSpectraS3Request {
    return &ImportPoolSpectraS3Request{
        Pool: pool,
    }
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.DataPolicyId = &dataPolicyId
    return importPoolSpectraS3Request
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithPriority(priority Priority) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.Priority = priority
    return importPoolSpectraS3Request
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.StorageDomainId = &storageDomainId
    return importPoolSpectraS3Request
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithUserId(userId string) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.UserId = &userId
    return importPoolSpectraS3Request
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.VerifyDataAfterImport = verifyDataAfterImport
    return importPoolSpectraS3Request
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.VerifyDataPriorToImport = &verifyDataPriorToImport
    return importPoolSpectraS3Request
}

