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

type ImportTapeSpectraS3Request struct {
    conflictResolutionMode ImportConflictResolutionMode
    dataPolicyId string
    priority Priority
    storageDomainId string
    tapeId string
    userId string
    verifyDataAfterImport Priority
    verifyDataPriorToImport bool
    queryParams *url.Values
}

func NewImportTapeSpectraS3Request(tapeId string) *ImportTapeSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "import")

    return &ImportTapeSpectraS3Request{
        tapeId: tapeId,
        queryParams: queryParams,
    }
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.conflictResolutionMode = conflictResolutionMode
    importTapeSpectraS3Request.queryParams.Set("conflict_resolution_mode", conflictResolutionMode.String())
    return importTapeSpectraS3Request
}
func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.dataPolicyId = dataPolicyId
    importTapeSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return importTapeSpectraS3Request
}
func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithPriority(priority Priority) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.priority = priority
    importTapeSpectraS3Request.queryParams.Set("priority", priority.String())
    return importTapeSpectraS3Request
}
func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.storageDomainId = storageDomainId
    importTapeSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return importTapeSpectraS3Request
}
func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithUserId(userId string) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.userId = userId
    importTapeSpectraS3Request.queryParams.Set("user_id", userId)
    return importTapeSpectraS3Request
}
func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.verifyDataAfterImport = verifyDataAfterImport
    importTapeSpectraS3Request.queryParams.Set("verify_data_after_import", verifyDataAfterImport.String())
    return importTapeSpectraS3Request
}
func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.verifyDataPriorToImport = verifyDataPriorToImport
    importTapeSpectraS3Request.queryParams.Set("verify_data_prior_to_import", strconv.FormatBool(verifyDataPriorToImport))
    return importTapeSpectraS3Request
}



func (ImportTapeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) Path() string {
    return "/_rest_/tape/" + importTapeSpectraS3Request.tapeId
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) QueryParams() *url.Values {
    return importTapeSpectraS3Request.queryParams
}

func (ImportTapeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ImportTapeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ImportTapeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
