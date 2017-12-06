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

type ModifyJobSpectraS3Request struct {
    CreatedAt *string
    JobId string
    Name *string
    Priority Priority
}

func NewModifyJobSpectraS3Request(jobId string) *ModifyJobSpectraS3Request {
    return &ModifyJobSpectraS3Request{
        JobId: jobId,
    }
}

func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) WithCreatedAt(createdAt string) *ModifyJobSpectraS3Request {
    modifyJobSpectraS3Request.CreatedAt = &createdAt
    return modifyJobSpectraS3Request
}

func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) WithName(name string) *ModifyJobSpectraS3Request {
    modifyJobSpectraS3Request.Name = &name
    return modifyJobSpectraS3Request
}

func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) WithPriority(priority Priority) *ModifyJobSpectraS3Request {
    modifyJobSpectraS3Request.Priority = priority
    return modifyJobSpectraS3Request
}

