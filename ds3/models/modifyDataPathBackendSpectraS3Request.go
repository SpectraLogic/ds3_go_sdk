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

type ModifyDataPathBackendSpectraS3Request struct {
    Activated *bool
    AutoActivateTimeoutInMins *int
    AutoInspect AutoInspectMode
    DefaultImportConflictResolutionMode ImportConflictResolutionMode
    DefaultVerifyDataAfterImport Priority
    DefaultVerifyDataPriorToImport *bool
    PartiallyVerifyLastPercentOfTapes *int
    UnavailableMediaPolicy UnavailableMediaUsagePolicy
    UnavailablePoolMaxJobRetryInMins *int
    UnavailableTapePartitionMaxJobRetryInMins *int
}

func NewModifyDataPathBackendSpectraS3Request() *ModifyDataPathBackendSpectraS3Request {
    return &ModifyDataPathBackendSpectraS3Request{
    }
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithActivated(activated bool) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.Activated = &activated
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithAutoActivateTimeoutInMins(autoActivateTimeoutInMins int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.AutoActivateTimeoutInMins = &autoActivateTimeoutInMins
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithAutoInspect(autoInspect AutoInspectMode) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.AutoInspect = autoInspect
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithDefaultImportConflictResolutionMode(defaultImportConflictResolutionMode ImportConflictResolutionMode) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.DefaultImportConflictResolutionMode = defaultImportConflictResolutionMode
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithDefaultVerifyDataAfterImport(defaultVerifyDataAfterImport Priority) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.DefaultVerifyDataAfterImport = defaultVerifyDataAfterImport
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithDefaultVerifyDataPriorToImport(defaultVerifyDataPriorToImport bool) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.DefaultVerifyDataPriorToImport = &defaultVerifyDataPriorToImport
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithPartiallyVerifyLastPercentOfTapes(partiallyVerifyLastPercentOfTapes int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.PartiallyVerifyLastPercentOfTapes = &partiallyVerifyLastPercentOfTapes
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithUnavailableMediaPolicy(unavailableMediaPolicy UnavailableMediaUsagePolicy) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.UnavailableMediaPolicy = unavailableMediaPolicy
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithUnavailablePoolMaxJobRetryInMins(unavailablePoolMaxJobRetryInMins int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.UnavailablePoolMaxJobRetryInMins = &unavailablePoolMaxJobRetryInMins
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithUnavailableTapePartitionMaxJobRetryInMins(unavailableTapePartitionMaxJobRetryInMins int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.UnavailableTapePartitionMaxJobRetryInMins = &unavailableTapePartitionMaxJobRetryInMins
    return modifyDataPathBackendSpectraS3Request
}

