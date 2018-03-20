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
    "fmt"
    "strings"
)

type Range struct {
    Start int64
    End int64
}

type GetObjectRequest struct {
    BucketName string
    ObjectName string
    Checksum Checksum
    Job *string
    Metadata map[string]string
    Offset *int64
    VersionId *string
}

func NewGetObjectRequest(bucketName string, objectName string) *GetObjectRequest {
    return &GetObjectRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        Checksum: NewNoneChecksum(),
        Metadata: make(map[string]string),
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

func (getObjectRequest *GetObjectRequest) WithVersionId(versionId string) *GetObjectRequest {
    getObjectRequest.VersionId = &versionId
    return getObjectRequest
}


func (getObjectRequest *GetObjectRequest) WithChecksum(contentHash string, checksumType ChecksumType) *GetObjectRequest {
    getObjectRequest.Checksum.ContentHash = contentHash
    getObjectRequest.Checksum.Type = checksumType
    return getObjectRequest
}

func (getObjectRequest *GetObjectRequest) WithRanges(ranges ...Range) *GetObjectRequest {
    var rangeElements []string
    for _, cur := range ranges {
        rangeElements = append(rangeElements, fmt.Sprintf("%d-%d", cur.Start, cur.End))
    }
    getObjectRequest.Metadata["Range"] = fmt.Sprintf("bytes=%s", strings.Join(rangeElements[:], ","))
    return getObjectRequest
}
