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

type PutDataPolicySpectraS3Request struct {
    AlwaysForcePutJobCreation *bool
    AlwaysMinimizeSpanningAcrossMedia *bool
    BlobbingEnabled *bool
    ChecksumType ChecksumType
    DefaultBlobSize *int64
    DefaultGetJobPriority Priority
    DefaultPutJobPriority Priority
    DefaultVerifyAfterWrite *bool
    DefaultVerifyJobPriority Priority
    EndToEndCrcRequired *bool
    Name string
    RebuildPriority Priority
    Versioning VersioningLevel
}

func NewPutDataPolicySpectraS3Request(name string) *PutDataPolicySpectraS3Request {
    return &PutDataPolicySpectraS3Request{
        Name: name,
    }
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithAlwaysForcePutJobCreation(alwaysForcePutJobCreation bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.AlwaysForcePutJobCreation = &alwaysForcePutJobCreation
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithAlwaysMinimizeSpanningAcrossMedia(alwaysMinimizeSpanningAcrossMedia bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.AlwaysMinimizeSpanningAcrossMedia = &alwaysMinimizeSpanningAcrossMedia
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithBlobbingEnabled(blobbingEnabled bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.BlobbingEnabled = &blobbingEnabled
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithChecksumType(checksumType ChecksumType) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.ChecksumType = checksumType
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultBlobSize(defaultBlobSize int64) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.DefaultBlobSize = &defaultBlobSize
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultGetJobPriority(defaultGetJobPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.DefaultGetJobPriority = defaultGetJobPriority
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultPutJobPriority(defaultPutJobPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.DefaultPutJobPriority = defaultPutJobPriority
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultVerifyAfterWrite(defaultVerifyAfterWrite bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.DefaultVerifyAfterWrite = &defaultVerifyAfterWrite
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultVerifyJobPriority(defaultVerifyJobPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.DefaultVerifyJobPriority = defaultVerifyJobPriority
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithEndToEndCrcRequired(endToEndCrcRequired bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.EndToEndCrcRequired = &endToEndCrcRequired
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithRebuildPriority(rebuildPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.RebuildPriority = rebuildPriority
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithVersioning(versioning VersioningLevel) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.Versioning = versioning
    return putDataPolicySpectraS3Request
}

