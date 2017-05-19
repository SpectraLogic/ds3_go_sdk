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

type VerifyAllTapesSpectraS3Request struct {
    taskPriority Priority
    queryParams *url.Values
}

func NewVerifyAllTapesSpectraS3Request() *VerifyAllTapesSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "verify")

    return &VerifyAllTapesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (verifyAllTapesSpectraS3Request *VerifyAllTapesSpectraS3Request) WithTaskPriority(taskPriority Priority) *VerifyAllTapesSpectraS3Request {
    verifyAllTapesSpectraS3Request.taskPriority = taskPriority
    verifyAllTapesSpectraS3Request.queryParams.Set("task_priority", taskPriority.String())
    return verifyAllTapesSpectraS3Request
}



func (VerifyAllTapesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (verifyAllTapesSpectraS3Request *VerifyAllTapesSpectraS3Request) Path() string {
    return "/_rest_/tape"
}

func (verifyAllTapesSpectraS3Request *VerifyAllTapesSpectraS3Request) QueryParams() *url.Values {
    return verifyAllTapesSpectraS3Request.queryParams
}

func (VerifyAllTapesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifyAllTapesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (VerifyAllTapesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
