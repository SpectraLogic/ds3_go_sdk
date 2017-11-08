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

type ModifyDataPolicySpectraS3Request struct {
    AlwaysForcePutJobCreation *bool
    AlwaysMinimizeSpanningAcrossMedia *bool
    BlobbingEnabled *bool
    ChecksumType ChecksumType
    DataPolicyId string
    DefaultBlobSize *int64
    DefaultGetJobPriority Priority
    DefaultPutJobPriority Priority
    DefaultVerifyAfterWrite *bool
    DefaultVerifyJobPriority Priority
    EndToEndCrcRequired *bool
    Name *string
    RebuildPriority Priority
    Versioning VersioningLevel
}

func NewModifyDataPolicySpectraS3Request(dataPolicyId string) *ModifyDataPolicySpectraS3Request {
    return &ModifyDataPolicySpectraS3Request{
        DataPolicyId: dataPolicyId,
    }
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithAlwaysForcePutJobCreation(alwaysForcePutJobCreation bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.AlwaysForcePutJobCreation = &alwaysForcePutJobCreation
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithAlwaysMinimizeSpanningAcrossMedia(alwaysMinimizeSpanningAcrossMedia bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.AlwaysMinimizeSpanningAcrossMedia = &alwaysMinimizeSpanningAcrossMedia
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithBlobbingEnabled(blobbingEnabled bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.BlobbingEnabled = &blobbingEnabled
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithChecksumType(checksumType ChecksumType) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.ChecksumType = checksumType
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultBlobSize(defaultBlobSize int64) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.DefaultBlobSize = &defaultBlobSize
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultGetJobPriority(defaultGetJobPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.DefaultGetJobPriority = defaultGetJobPriority
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultPutJobPriority(defaultPutJobPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.DefaultPutJobPriority = defaultPutJobPriority
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultVerifyAfterWrite(defaultVerifyAfterWrite bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.DefaultVerifyAfterWrite = &defaultVerifyAfterWrite
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultVerifyJobPriority(defaultVerifyJobPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.DefaultVerifyJobPriority = defaultVerifyJobPriority
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithEndToEndCrcRequired(endToEndCrcRequired bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.EndToEndCrcRequired = &endToEndCrcRequired
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithName(name string) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.Name = &name
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithRebuildPriority(rebuildPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.RebuildPriority = rebuildPriority
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithVersioning(versioning VersioningLevel) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.Versioning = versioning
    return modifyDataPolicySpectraS3Request
}

