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

type VerifyTapeSpectraS3Request struct {
    tapeId string
    taskPriority Priority
    queryParams *url.Values
}

func NewVerifyTapeSpectraS3Request(tapeId string) *VerifyTapeSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "verify")

    return &VerifyTapeSpectraS3Request{
        tapeId: tapeId,
        queryParams: queryParams,
    }
}

func (verifyTapeSpectraS3Request *VerifyTapeSpectraS3Request) WithTaskPriority(taskPriority Priority) *VerifyTapeSpectraS3Request {
    verifyTapeSpectraS3Request.taskPriority = taskPriority
    verifyTapeSpectraS3Request.queryParams.Set("task_priority", taskPriority.String())
    return verifyTapeSpectraS3Request
}



func (VerifyTapeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (verifyTapeSpectraS3Request *VerifyTapeSpectraS3Request) Path() string {
    return "/_rest_/tape/" + verifyTapeSpectraS3Request.tapeId
}

func (verifyTapeSpectraS3Request *VerifyTapeSpectraS3Request) QueryParams() *url.Values {
    return verifyTapeSpectraS3Request.queryParams
}

func (VerifyTapeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifyTapeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (VerifyTapeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
