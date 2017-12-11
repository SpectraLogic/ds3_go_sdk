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

type ImportAllPoolsSpectraS3Request struct {
    ConflictResolutionMode ImportConflictResolutionMode
    DataPolicyId *string
    Priority Priority
    StorageDomainId *string
    UserId *string
    VerifyDataAfterImport Priority
    VerifyDataPriorToImport *bool
}

func NewImportAllPoolsSpectraS3Request() *ImportAllPoolsSpectraS3Request {
    return &ImportAllPoolsSpectraS3Request{
    }
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.ConflictResolutionMode = conflictResolutionMode
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.DataPolicyId = &dataPolicyId
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithPriority(priority Priority) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.Priority = priority
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.StorageDomainId = &storageDomainId
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithUserId(userId string) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.UserId = &userId
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.VerifyDataAfterImport = verifyDataAfterImport
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.VerifyDataPriorToImport = &verifyDataPriorToImport
    return importAllPoolsSpectraS3Request
}

