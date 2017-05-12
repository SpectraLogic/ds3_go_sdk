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

type ModifyDataPathBackendSpectraS3Request struct {
    activated bool
    autoActivateTimeoutInMins *int
    autoInspect AutoInspectMode
    defaultImportConflictResolutionMode ImportConflictResolutionMode
    defaultVerifyDataAfterImport Priority
    defaultVerifyDataPriorToImport bool
    partiallyVerifyLastPercentOfTapes *int
    unavailableMediaPolicy UnavailableMediaUsagePolicy
    unavailablePoolMaxJobRetryInMins int
    unavailableTapePartitionMaxJobRetryInMins int
    queryParams *url.Values
}

func NewModifyDataPathBackendSpectraS3Request() *ModifyDataPathBackendSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyDataPathBackendSpectraS3Request{
        queryParams: queryParams,
    }
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithActivated(activated bool) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.activated = activated
    modifyDataPathBackendSpectraS3Request.queryParams.Set("activated", strconv.FormatBool(activated))
    return modifyDataPathBackendSpectraS3Request
}
func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithAutoInspect(autoInspect AutoInspectMode) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.autoInspect = autoInspect
    modifyDataPathBackendSpectraS3Request.queryParams.Set("auto_inspect", autoInspect.String())
    return modifyDataPathBackendSpectraS3Request
}
func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithDefaultImportConflictResolutionMode(defaultImportConflictResolutionMode ImportConflictResolutionMode) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.defaultImportConflictResolutionMode = defaultImportConflictResolutionMode
    modifyDataPathBackendSpectraS3Request.queryParams.Set("default_import_conflict_resolution_mode", defaultImportConflictResolutionMode.String())
    return modifyDataPathBackendSpectraS3Request
}
func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithDefaultVerifyDataAfterImport(defaultVerifyDataAfterImport Priority) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.defaultVerifyDataAfterImport = defaultVerifyDataAfterImport
    modifyDataPathBackendSpectraS3Request.queryParams.Set("default_verify_data_after_import", defaultVerifyDataAfterImport.String())
    return modifyDataPathBackendSpectraS3Request
}
func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithDefaultVerifyDataPriorToImport(defaultVerifyDataPriorToImport bool) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.defaultVerifyDataPriorToImport = defaultVerifyDataPriorToImport
    modifyDataPathBackendSpectraS3Request.queryParams.Set("default_verify_data_prior_to_import", strconv.FormatBool(defaultVerifyDataPriorToImport))
    return modifyDataPathBackendSpectraS3Request
}
func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithUnavailableMediaPolicy(unavailableMediaPolicy UnavailableMediaUsagePolicy) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.unavailableMediaPolicy = unavailableMediaPolicy
    modifyDataPathBackendSpectraS3Request.queryParams.Set("unavailable_media_policy", unavailableMediaPolicy.String())
    return modifyDataPathBackendSpectraS3Request
}
func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithUnavailablePoolMaxJobRetryInMins(unavailablePoolMaxJobRetryInMins int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.unavailablePoolMaxJobRetryInMins = unavailablePoolMaxJobRetryInMins
    modifyDataPathBackendSpectraS3Request.queryParams.Set("unavailable_pool_max_job_retry_in_mins", strconv.Itoa(unavailablePoolMaxJobRetryInMins))
    return modifyDataPathBackendSpectraS3Request
}
func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithUnavailableTapePartitionMaxJobRetryInMins(unavailableTapePartitionMaxJobRetryInMins int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.unavailableTapePartitionMaxJobRetryInMins = unavailableTapePartitionMaxJobRetryInMins
    modifyDataPathBackendSpectraS3Request.queryParams.Set("unavailable_tape_partition_max_job_retry_in_mins", strconv.Itoa(unavailableTapePartitionMaxJobRetryInMins))
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithAutoActivateTimeoutInMins(autoActivateTimeoutInMins *int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.autoActivateTimeoutInMins = autoActivateTimeoutInMins
    if autoActivateTimeoutInMins != nil {
        modifyDataPathBackendSpectraS3Request.queryParams.Set("auto_activate_timeout_in_mins", strconv.Itoa(*autoActivateTimeoutInMins))
    } else {
        modifyDataPathBackendSpectraS3Request.queryParams.Set("auto_activate_timeout_in_mins", "")
    }
    return modifyDataPathBackendSpectraS3Request
}
func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithPartiallyVerifyLastPercentOfTapes(partiallyVerifyLastPercentOfTapes *int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.partiallyVerifyLastPercentOfTapes = partiallyVerifyLastPercentOfTapes
    if partiallyVerifyLastPercentOfTapes != nil {
        modifyDataPathBackendSpectraS3Request.queryParams.Set("partially_verify_last_percent_of_tapes", strconv.Itoa(*partiallyVerifyLastPercentOfTapes))
    } else {
        modifyDataPathBackendSpectraS3Request.queryParams.Set("partially_verify_last_percent_of_tapes", "")
    }
    return modifyDataPathBackendSpectraS3Request
}


func (ModifyDataPathBackendSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) Path() string {
    return "/_rest_/data_path_backend"
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) QueryParams() *url.Values {
    return modifyDataPathBackendSpectraS3Request.queryParams
}

func (ModifyDataPathBackendSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyDataPathBackendSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyDataPathBackendSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
