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

type RawImportTapeSpectraS3Request struct {
    bucketId string
    storageDomainId string
    tapeId string
    taskPriority Priority
    queryParams *url.Values
}

func NewRawImportTapeSpectraS3Request(bucketId string, tapeId string) *RawImportTapeSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "import")
    queryParams.Set("bucket_id", bucketId)

    return &RawImportTapeSpectraS3Request{
        tapeId: tapeId,
        bucketId: bucketId,
        queryParams: queryParams,
    }
}

func (rawImportTapeSpectraS3Request *RawImportTapeSpectraS3Request) WithStorageDomainId(storageDomainId string) *RawImportTapeSpectraS3Request {
    rawImportTapeSpectraS3Request.storageDomainId = storageDomainId
    rawImportTapeSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return rawImportTapeSpectraS3Request
}
func (rawImportTapeSpectraS3Request *RawImportTapeSpectraS3Request) WithTaskPriority(taskPriority Priority) *RawImportTapeSpectraS3Request {
    rawImportTapeSpectraS3Request.taskPriority = taskPriority
    rawImportTapeSpectraS3Request.queryParams.Set("task_priority", taskPriority.String())
    return rawImportTapeSpectraS3Request
}



func (RawImportTapeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (rawImportTapeSpectraS3Request *RawImportTapeSpectraS3Request) Path() string {
    return "/_rest_/tape/" + rawImportTapeSpectraS3Request.tapeId
}

func (rawImportTapeSpectraS3Request *RawImportTapeSpectraS3Request) QueryParams() *url.Values {
    return rawImportTapeSpectraS3Request.queryParams
}

func (RawImportTapeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (RawImportTapeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (RawImportTapeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
