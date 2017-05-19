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

type ImportAllPoolsSpectraS3Request struct {
    conflictResolutionMode ImportConflictResolutionMode
    dataPolicyId string
    priority Priority
    storageDomainId string
    userId string
    verifyDataAfterImport Priority
    verifyDataPriorToImport bool
    queryParams *url.Values
}

func NewImportAllPoolsSpectraS3Request() *ImportAllPoolsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "import")

    return &ImportAllPoolsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.conflictResolutionMode = conflictResolutionMode
    importAllPoolsSpectraS3Request.queryParams.Set("conflict_resolution_mode", conflictResolutionMode.String())
    return importAllPoolsSpectraS3Request
}
func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.dataPolicyId = dataPolicyId
    importAllPoolsSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return importAllPoolsSpectraS3Request
}
func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithPriority(priority Priority) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.priority = priority
    importAllPoolsSpectraS3Request.queryParams.Set("priority", priority.String())
    return importAllPoolsSpectraS3Request
}
func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.storageDomainId = storageDomainId
    importAllPoolsSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return importAllPoolsSpectraS3Request
}
func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithUserId(userId string) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.userId = userId
    importAllPoolsSpectraS3Request.queryParams.Set("user_id", userId)
    return importAllPoolsSpectraS3Request
}
func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.verifyDataAfterImport = verifyDataAfterImport
    importAllPoolsSpectraS3Request.queryParams.Set("verify_data_after_import", verifyDataAfterImport.String())
    return importAllPoolsSpectraS3Request
}
func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.verifyDataPriorToImport = verifyDataPriorToImport
    importAllPoolsSpectraS3Request.queryParams.Set("verify_data_prior_to_import", strconv.FormatBool(verifyDataPriorToImport))
    return importAllPoolsSpectraS3Request
}



func (ImportAllPoolsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) Path() string {
    return "/_rest_/pool"
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) QueryParams() *url.Values {
    return importAllPoolsSpectraS3Request.queryParams
}

func (ImportAllPoolsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ImportAllPoolsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ImportAllPoolsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
