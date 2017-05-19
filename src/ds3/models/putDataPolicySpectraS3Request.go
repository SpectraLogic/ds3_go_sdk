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

type PutDataPolicySpectraS3Request struct {
    alwaysForcePutJobCreation bool
    alwaysMinimizeSpanningAcrossMedia bool
    blobbingEnabled bool
    checksumType ChecksumType
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

func NewPutDataPolicySpectraS3Request(name *string) *PutDataPolicySpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("name", *name)

    return &PutDataPolicySpectraS3Request{
        name: name,
        queryParams: queryParams,
    }
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithAlwaysForcePutJobCreation(alwaysForcePutJobCreation bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.alwaysForcePutJobCreation = alwaysForcePutJobCreation
    putDataPolicySpectraS3Request.queryParams.Set("always_force_put_job_creation", strconv.FormatBool(alwaysForcePutJobCreation))
    return putDataPolicySpectraS3Request
}
func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithAlwaysMinimizeSpanningAcrossMedia(alwaysMinimizeSpanningAcrossMedia bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.alwaysMinimizeSpanningAcrossMedia = alwaysMinimizeSpanningAcrossMedia
    putDataPolicySpectraS3Request.queryParams.Set("always_minimize_spanning_across_media", strconv.FormatBool(alwaysMinimizeSpanningAcrossMedia))
    return putDataPolicySpectraS3Request
}
func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithBlobbingEnabled(blobbingEnabled bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.blobbingEnabled = blobbingEnabled
    putDataPolicySpectraS3Request.queryParams.Set("blobbing_enabled", strconv.FormatBool(blobbingEnabled))
    return putDataPolicySpectraS3Request
}
func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithChecksumType(checksumType ChecksumType) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.checksumType = checksumType
    putDataPolicySpectraS3Request.queryParams.Set("checksum_type", checksumType.String())
    return putDataPolicySpectraS3Request
}
func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultGetJobPriority(defaultGetJobPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.defaultGetJobPriority = defaultGetJobPriority
    putDataPolicySpectraS3Request.queryParams.Set("default_get_job_priority", defaultGetJobPriority.String())
    return putDataPolicySpectraS3Request
}
func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultPutJobPriority(defaultPutJobPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.defaultPutJobPriority = defaultPutJobPriority
    putDataPolicySpectraS3Request.queryParams.Set("default_put_job_priority", defaultPutJobPriority.String())
    return putDataPolicySpectraS3Request
}
func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultVerifyAfterWrite(defaultVerifyAfterWrite bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.defaultVerifyAfterWrite = defaultVerifyAfterWrite
    putDataPolicySpectraS3Request.queryParams.Set("default_verify_after_write", strconv.FormatBool(defaultVerifyAfterWrite))
    return putDataPolicySpectraS3Request
}
func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultVerifyJobPriority(defaultVerifyJobPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.defaultVerifyJobPriority = defaultVerifyJobPriority
    putDataPolicySpectraS3Request.queryParams.Set("default_verify_job_priority", defaultVerifyJobPriority.String())
    return putDataPolicySpectraS3Request
}
func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithEndToEndCrcRequired(endToEndCrcRequired bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.endToEndCrcRequired = endToEndCrcRequired
    putDataPolicySpectraS3Request.queryParams.Set("end_to_end_crc_required", strconv.FormatBool(endToEndCrcRequired))
    return putDataPolicySpectraS3Request
}
func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithRebuildPriority(rebuildPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.rebuildPriority = rebuildPriority
    putDataPolicySpectraS3Request.queryParams.Set("rebuild_priority", rebuildPriority.String())
    return putDataPolicySpectraS3Request
}
func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithVersioning(versioning VersioningLevel) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.versioning = versioning
    putDataPolicySpectraS3Request.queryParams.Set("versioning", versioning.String())
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultBlobSize(defaultBlobSize *int64) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.defaultBlobSize = defaultBlobSize
    if defaultBlobSize != nil {
        putDataPolicySpectraS3Request.queryParams.Set("default_blob_size", strconv.FormatInt(*defaultBlobSize, 10))
    } else {
        putDataPolicySpectraS3Request.queryParams.Set("default_blob_size", "")
    }
    return putDataPolicySpectraS3Request
}


func (PutDataPolicySpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) Path() string {
    return "/_rest_/data_policy"
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) QueryParams() *url.Values {
    return putDataPolicySpectraS3Request.queryParams
}

func (PutDataPolicySpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutDataPolicySpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutDataPolicySpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
