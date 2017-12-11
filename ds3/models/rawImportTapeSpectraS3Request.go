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

type RawImportTapeSpectraS3Request struct {
    BucketId string
    StorageDomainId *string
    TapeId string
    TaskPriority Priority
}

func NewRawImportTapeSpectraS3Request(bucketId string, tapeId string) *RawImportTapeSpectraS3Request {
    return &RawImportTapeSpectraS3Request{
        TapeId: tapeId,
        BucketId: bucketId,
    }
}

func (rawImportTapeSpectraS3Request *RawImportTapeSpectraS3Request) WithStorageDomainId(storageDomainId string) *RawImportTapeSpectraS3Request {
    rawImportTapeSpectraS3Request.StorageDomainId = &storageDomainId
    return rawImportTapeSpectraS3Request
}

func (rawImportTapeSpectraS3Request *RawImportTapeSpectraS3Request) WithTaskPriority(taskPriority Priority) *RawImportTapeSpectraS3Request {
    rawImportTapeSpectraS3Request.TaskPriority = taskPriority
    return rawImportTapeSpectraS3Request
}

