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

type ImportAllTapesSpectraS3Request struct {
    conflictResolutionMode ImportConflictResolutionMode
    dataPolicyId string
    priority Priority
    storageDomainId string
    userId string
    verifyDataAfterImport Priority
    verifyDataPriorToImport bool
    queryParams *url.Values
}

func NewImportAllTapesSpectraS3Request() *ImportAllTapesSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "import")

    return &ImportAllTapesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.conflictResolutionMode = conflictResolutionMode
    importAllTapesSpectraS3Request.queryParams.Set("conflict_resolution_mode", conflictResolutionMode.String())
    return importAllTapesSpectraS3Request
}
func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.dataPolicyId = dataPolicyId
    importAllTapesSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return importAllTapesSpectraS3Request
}
func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithPriority(priority Priority) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.priority = priority
    importAllTapesSpectraS3Request.queryParams.Set("priority", priority.String())
    return importAllTapesSpectraS3Request
}
func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.storageDomainId = storageDomainId
    importAllTapesSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return importAllTapesSpectraS3Request
}
func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithUserId(userId string) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.userId = userId
    importAllTapesSpectraS3Request.queryParams.Set("user_id", userId)
    return importAllTapesSpectraS3Request
}
func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.verifyDataAfterImport = verifyDataAfterImport
    importAllTapesSpectraS3Request.queryParams.Set("verify_data_after_import", verifyDataAfterImport.String())
    return importAllTapesSpectraS3Request
}
func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.verifyDataPriorToImport = verifyDataPriorToImport
    importAllTapesSpectraS3Request.queryParams.Set("verify_data_prior_to_import", strconv.FormatBool(verifyDataPriorToImport))
    return importAllTapesSpectraS3Request
}



func (ImportAllTapesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) Path() string {
    return "/_rest_/tape"
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) QueryParams() *url.Values {
    return importAllTapesSpectraS3Request.queryParams
}

func (ImportAllTapesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ImportAllTapesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ImportAllTapesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
