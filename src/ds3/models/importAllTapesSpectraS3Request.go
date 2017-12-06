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

type ImportAllTapesSpectraS3Request struct {
    ConflictResolutionMode ImportConflictResolutionMode
    DataPolicyId *string
    Priority Priority
    StorageDomainId *string
    UserId *string
    VerifyDataAfterImport Priority
    VerifyDataPriorToImport *bool
}

func NewImportAllTapesSpectraS3Request() *ImportAllTapesSpectraS3Request {
    return &ImportAllTapesSpectraS3Request{
    }
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.ConflictResolutionMode = conflictResolutionMode
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.DataPolicyId = &dataPolicyId
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithPriority(priority Priority) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.Priority = priority
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.StorageDomainId = &storageDomainId
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithUserId(userId string) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.UserId = &userId
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.VerifyDataAfterImport = verifyDataAfterImport
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.VerifyDataPriorToImport = &verifyDataPriorToImport
    return importAllTapesSpectraS3Request
}

