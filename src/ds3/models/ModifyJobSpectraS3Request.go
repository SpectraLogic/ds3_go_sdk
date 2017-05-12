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
)

type ModifyJobSpectraS3Request struct {
    createdAt string
    jobId string
    name *string
    priority Priority
    queryParams *url.Values
}

func NewModifyJobSpectraS3Request(jobId string) *ModifyJobSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyJobSpectraS3Request{
        jobId: jobId,
        queryParams: queryParams,
    }
}

func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) WithCreatedAt(createdAt string) *ModifyJobSpectraS3Request {
    modifyJobSpectraS3Request.createdAt = createdAt
    modifyJobSpectraS3Request.queryParams.Set("created_at", createdAt)
    return modifyJobSpectraS3Request
}
func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) WithPriority(priority Priority) *ModifyJobSpectraS3Request {
    modifyJobSpectraS3Request.priority = priority
    modifyJobSpectraS3Request.queryParams.Set("priority", priority.String())
    return modifyJobSpectraS3Request
}

func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) WithName(name *string) *ModifyJobSpectraS3Request {
    modifyJobSpectraS3Request.name = name
    if name != nil {
        modifyJobSpectraS3Request.queryParams.Set("name", *name)
    } else {
        modifyJobSpectraS3Request.queryParams.Set("name", "")
    }
    return modifyJobSpectraS3Request
}


func (ModifyJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) Path() string {
    return "/_rest_/job/" + modifyJobSpectraS3Request.jobId
}

func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) QueryParams() *url.Values {
    return modifyJobSpectraS3Request.queryParams
}

func (ModifyJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
