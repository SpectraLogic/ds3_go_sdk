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

type ModifyDataPolicySpectraS3Request struct {
    alwaysForcePutJobCreation bool
    alwaysMinimizeSpanningAcrossMedia bool
    blobbingEnabled bool
    checksumType ChecksumType
    dataPolicyId string
    defaultBlobSize *int64
    defaultGetJobPriority Priority
    defaultPutJobPriority Priority
    defaultVerifyAfterWrite bool
    defaultVerifyJobPriority Priority
    endToEndCrcRequired bool
    name *string
    rebuildPriority Priority
    versioning VersioningLevel
    queryParams *url.Values
}

func NewModifyDataPolicySpectraS3Request(dataPolicyId string) *ModifyDataPolicySpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyDataPolicySpectraS3Request{
        dataPolicyId: dataPolicyId,
        queryParams: queryParams,
    }
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithAlwaysForcePutJobCreation(alwaysForcePutJobCreation bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.alwaysForcePutJobCreation = alwaysForcePutJobCreation
    modifyDataPolicySpectraS3Request.queryParams.Set("always_force_put_job_creation", strconv.FormatBool(alwaysForcePutJobCreation))
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithAlwaysMinimizeSpanningAcrossMedia(alwaysMinimizeSpanningAcrossMedia bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.alwaysMinimizeSpanningAcrossMedia = alwaysMinimizeSpanningAcrossMedia
    modifyDataPolicySpectraS3Request.queryParams.Set("always_minimize_spanning_across_media", strconv.FormatBool(alwaysMinimizeSpanningAcrossMedia))
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithBlobbingEnabled(blobbingEnabled bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.blobbingEnabled = blobbingEnabled
    modifyDataPolicySpectraS3Request.queryParams.Set("blobbing_enabled", strconv.FormatBool(blobbingEnabled))
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithChecksumType(checksumType ChecksumType) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.checksumType = checksumType
    modifyDataPolicySpectraS3Request.queryParams.Set("checksum_type", checksumType.String())
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultGetJobPriority(defaultGetJobPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.defaultGetJobPriority = defaultGetJobPriority
    modifyDataPolicySpectraS3Request.queryParams.Set("default_get_job_priority", defaultGetJobPriority.String())
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultPutJobPriority(defaultPutJobPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.defaultPutJobPriority = defaultPutJobPriority
    modifyDataPolicySpectraS3Request.queryParams.Set("default_put_job_priority", defaultPutJobPriority.String())
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultVerifyAfterWrite(defaultVerifyAfterWrite bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.defaultVerifyAfterWrite = defaultVerifyAfterWrite
    modifyDataPolicySpectraS3Request.queryParams.Set("default_verify_after_write", strconv.FormatBool(defaultVerifyAfterWrite))
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultVerifyJobPriority(defaultVerifyJobPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.defaultVerifyJobPriority = defaultVerifyJobPriority
    modifyDataPolicySpectraS3Request.queryParams.Set("default_verify_job_priority", defaultVerifyJobPriority.String())
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithEndToEndCrcRequired(endToEndCrcRequired bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.endToEndCrcRequired = endToEndCrcRequired
    modifyDataPolicySpectraS3Request.queryParams.Set("end_to_end_crc_required", strconv.FormatBool(endToEndCrcRequired))
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithRebuildPriority(rebuildPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.rebuildPriority = rebuildPriority
    modifyDataPolicySpectraS3Request.queryParams.Set("rebuild_priority", rebuildPriority.String())
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithVersioning(versioning VersioningLevel) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.versioning = versioning
    modifyDataPolicySpectraS3Request.queryParams.Set("versioning", versioning.String())
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultBlobSize(defaultBlobSize *int64) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.defaultBlobSize = defaultBlobSize
    if defaultBlobSize != nil {
        modifyDataPolicySpectraS3Request.queryParams.Set("default_blob_size", strconv.FormatInt(*defaultBlobSize, 10))
    } else {
        modifyDataPolicySpectraS3Request.queryParams.Set("default_blob_size", "")
    }
    return modifyDataPolicySpectraS3Request
}
func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithName(name *string) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.name = name
    if name != nil {
        modifyDataPolicySpectraS3Request.queryParams.Set("name", *name)
    } else {
        modifyDataPolicySpectraS3Request.queryParams.Set("name", "")
    }
    return modifyDataPolicySpectraS3Request
}


func (ModifyDataPolicySpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) Path() string {
    return "/_rest_/data_policy/" + modifyDataPolicySpectraS3Request.dataPolicyId
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) QueryParams() *url.Values {
    return modifyDataPolicySpectraS3Request.queryParams
}

func (ModifyDataPolicySpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyDataPolicySpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyDataPolicySpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
