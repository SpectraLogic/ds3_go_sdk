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

type InspectAllTapesSpectraS3Request struct {
    taskPriority Priority
    queryParams *url.Values
}

func NewInspectAllTapesSpectraS3Request() *InspectAllTapesSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "inspect")

    return &InspectAllTapesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (inspectAllTapesSpectraS3Request *InspectAllTapesSpectraS3Request) WithTaskPriority(taskPriority Priority) *InspectAllTapesSpectraS3Request {
    inspectAllTapesSpectraS3Request.taskPriority = taskPriority
    inspectAllTapesSpectraS3Request.queryParams.Set("task_priority", taskPriority.String())
    return inspectAllTapesSpectraS3Request
}



func (InspectAllTapesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (inspectAllTapesSpectraS3Request *InspectAllTapesSpectraS3Request) Path() string {
    return "/_rest_/tape"
}

func (inspectAllTapesSpectraS3Request *InspectAllTapesSpectraS3Request) QueryParams() *url.Values {
    return inspectAllTapesSpectraS3Request.queryParams
}

func (InspectAllTapesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (InspectAllTapesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (InspectAllTapesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
