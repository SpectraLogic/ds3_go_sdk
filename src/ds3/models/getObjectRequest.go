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
    "ds3/networking"
    "fmt"
)

type rangeHeader struct {
    start, end int
}

type GetObjectRequest struct {
    BucketName string
    ObjectName string
    Checksum networking.Checksum
    Job *string
    Metadata map[string]string
    Offset *int64
}

func NewGetObjectRequest(bucketName string, objectName string) *GetObjectRequest {
    return &GetObjectRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        Checksum: networking.NewNoneChecksum(),
    }
}

func (getObjectRequest *GetObjectRequest) WithJob(job string) *GetObjectRequest {
    getObjectRequest.Job = &job
    return getObjectRequest
}

func (getObjectRequest *GetObjectRequest) WithOffset(offset int64) *GetObjectRequest {
    getObjectRequest.Offset = &offset
    return getObjectRequest
}


func (getObjectRequest *GetObjectRequest) WithChecksum(contentHash string, checksumType networking.ChecksumType) *GetObjectRequest {
    getObjectRequest.Checksum.ContentHash = contentHash
    getObjectRequest.Checksum.Type = checksumType
    return getObjectRequest
}

func (getObjectRequest *GetObjectRequest) WithRange(start, end int) *GetObjectRequest {
    getObjectRequest.Metadata["Range"] = fmt.Sprintf("bytes=%d-%d", start, end)
    return getObjectRequest
}
