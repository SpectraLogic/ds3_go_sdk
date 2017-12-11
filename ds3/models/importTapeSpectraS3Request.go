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

type ImportTapeSpectraS3Request struct {
    ConflictResolutionMode ImportConflictResolutionMode
    DataPolicyId *string
    Priority Priority
    StorageDomainId *string
    TapeId string
    UserId *string
    VerifyDataAfterImport Priority
    VerifyDataPriorToImport *bool
}

func NewImportTapeSpectraS3Request(tapeId string) *ImportTapeSpectraS3Request {
    return &ImportTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.ConflictResolutionMode = conflictResolutionMode
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.DataPolicyId = &dataPolicyId
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithPriority(priority Priority) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.Priority = priority
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.StorageDomainId = &storageDomainId
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithUserId(userId string) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.UserId = &userId
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.VerifyDataAfterImport = verifyDataAfterImport
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.VerifyDataPriorToImport = &verifyDataPriorToImport
    return importTapeSpectraS3Request
}

