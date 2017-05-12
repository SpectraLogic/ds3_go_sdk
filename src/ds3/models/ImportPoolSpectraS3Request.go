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

type ImportPoolSpectraS3Request struct {
    conflictResolutionMode ImportConflictResolutionMode
    dataPolicyId string
    pool string
    priority Priority
    storageDomainId string
    userId string
    verifyDataAfterImport Priority
    verifyDataPriorToImport bool
    queryParams *url.Values
}

func NewImportPoolSpectraS3Request(pool string) *ImportPoolSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "import")

    return &ImportPoolSpectraS3Request{
        pool: pool,
        queryParams: queryParams,
    }
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.conflictResolutionMode = conflictResolutionMode
    importPoolSpectraS3Request.queryParams.Set("conflict_resolution_mode", conflictResolutionMode.String())
    return importPoolSpectraS3Request
}
func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.dataPolicyId = dataPolicyId
    importPoolSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return importPoolSpectraS3Request
}
func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithPriority(priority Priority) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.priority = priority
    importPoolSpectraS3Request.queryParams.Set("priority", priority.String())
    return importPoolSpectraS3Request
}
func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.storageDomainId = storageDomainId
    importPoolSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return importPoolSpectraS3Request
}
func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithUserId(userId string) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.userId = userId
    importPoolSpectraS3Request.queryParams.Set("user_id", userId)
    return importPoolSpectraS3Request
}
func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.verifyDataAfterImport = verifyDataAfterImport
    importPoolSpectraS3Request.queryParams.Set("verify_data_after_import", verifyDataAfterImport.String())
    return importPoolSpectraS3Request
}
func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.verifyDataPriorToImport = verifyDataPriorToImport
    importPoolSpectraS3Request.queryParams.Set("verify_data_prior_to_import", strconv.FormatBool(verifyDataPriorToImport))
    return importPoolSpectraS3Request
}



func (ImportPoolSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) Path() string {
    return "/_rest_/pool/" + importPoolSpectraS3Request.pool
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) QueryParams() *url.Values {
    return importPoolSpectraS3Request.queryParams
}

func (ImportPoolSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ImportPoolSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ImportPoolSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
