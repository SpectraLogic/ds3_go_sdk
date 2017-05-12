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

type ModifyActiveJobSpectraS3Request struct {
    activeJobId string
    createdAt string
    name *string
    priority Priority
    queryParams *url.Values
}

func NewModifyActiveJobSpectraS3Request(activeJobId string) *ModifyActiveJobSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyActiveJobSpectraS3Request{
        activeJobId: activeJobId,
        queryParams: queryParams,
    }
}

func (modifyActiveJobSpectraS3Request *ModifyActiveJobSpectraS3Request) WithCreatedAt(createdAt string) *ModifyActiveJobSpectraS3Request {
    modifyActiveJobSpectraS3Request.createdAt = createdAt
    modifyActiveJobSpectraS3Request.queryParams.Set("created_at", createdAt)
    return modifyActiveJobSpectraS3Request
}
func (modifyActiveJobSpectraS3Request *ModifyActiveJobSpectraS3Request) WithPriority(priority Priority) *ModifyActiveJobSpectraS3Request {
    modifyActiveJobSpectraS3Request.priority = priority
    modifyActiveJobSpectraS3Request.queryParams.Set("priority", priority.String())
    return modifyActiveJobSpectraS3Request
}

func (modifyActiveJobSpectraS3Request *ModifyActiveJobSpectraS3Request) WithName(name *string) *ModifyActiveJobSpectraS3Request {
    modifyActiveJobSpectraS3Request.name = name
    if name != nil {
        modifyActiveJobSpectraS3Request.queryParams.Set("name", *name)
    } else {
        modifyActiveJobSpectraS3Request.queryParams.Set("name", "")
    }
    return modifyActiveJobSpectraS3Request
}


func (ModifyActiveJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyActiveJobSpectraS3Request *ModifyActiveJobSpectraS3Request) Path() string {
    return "/_rest_/active_job/" + modifyActiveJobSpectraS3Request.activeJobId
}

func (modifyActiveJobSpectraS3Request *ModifyActiveJobSpectraS3Request) QueryParams() *url.Values {
    return modifyActiveJobSpectraS3Request.queryParams
}

func (ModifyActiveJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyActiveJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyActiveJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
