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

type RawImportAllTapesSpectraS3Request struct {
    bucketId string
    taskPriority Priority
    queryParams *url.Values
}

func NewRawImportAllTapesSpectraS3Request(bucketId string) *RawImportAllTapesSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "import")
    queryParams.Set("bucket_id", bucketId)

    return &RawImportAllTapesSpectraS3Request{
        bucketId: bucketId,
        queryParams: queryParams,
    }
}

func (rawImportAllTapesSpectraS3Request *RawImportAllTapesSpectraS3Request) WithTaskPriority(taskPriority Priority) *RawImportAllTapesSpectraS3Request {
    rawImportAllTapesSpectraS3Request.taskPriority = taskPriority
    rawImportAllTapesSpectraS3Request.queryParams.Set("task_priority", taskPriority.String())
    return rawImportAllTapesSpectraS3Request
}



func (RawImportAllTapesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (rawImportAllTapesSpectraS3Request *RawImportAllTapesSpectraS3Request) Path() string {
    return "/_rest_/tape"
}

func (rawImportAllTapesSpectraS3Request *RawImportAllTapesSpectraS3Request) QueryParams() *url.Values {
    return rawImportAllTapesSpectraS3Request.queryParams
}

func (RawImportAllTapesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (RawImportAllTapesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (RawImportAllTapesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
