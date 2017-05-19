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

type InspectTapeSpectraS3Request struct {
    tapeId string
    taskPriority Priority
    queryParams *url.Values
}

func NewInspectTapeSpectraS3Request(tapeId string) *InspectTapeSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "inspect")

    return &InspectTapeSpectraS3Request{
        tapeId: tapeId,
        queryParams: queryParams,
    }
}

func (inspectTapeSpectraS3Request *InspectTapeSpectraS3Request) WithTaskPriority(taskPriority Priority) *InspectTapeSpectraS3Request {
    inspectTapeSpectraS3Request.taskPriority = taskPriority
    inspectTapeSpectraS3Request.queryParams.Set("task_priority", taskPriority.String())
    return inspectTapeSpectraS3Request
}



func (InspectTapeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (inspectTapeSpectraS3Request *InspectTapeSpectraS3Request) Path() string {
    return "/_rest_/tape/" + inspectTapeSpectraS3Request.tapeId
}

func (inspectTapeSpectraS3Request *InspectTapeSpectraS3Request) QueryParams() *url.Values {
    return inspectTapeSpectraS3Request.queryParams
}

func (InspectTapeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (InspectTapeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (InspectTapeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
