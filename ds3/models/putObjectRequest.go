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
    "strings"
)

type PutObjectRequest struct {
    BucketName string
    ObjectName string
    Checksum Checksum
    Content ReaderWithSizeDecorator
    Job *string
    Metadata map[string]string
    Offset *int64
}

func NewPutObjectRequest(bucketName string, objectName string, content ReaderWithSizeDecorator) *PutObjectRequest {
    return &PutObjectRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        Content: content,
        Checksum: NewNoneChecksum(),
        Metadata: make(map[string]string),
    }
}

func (putObjectRequest *PutObjectRequest) WithJob(job string) *PutObjectRequest {
    putObjectRequest.Job = &job
    return putObjectRequest
}

func (putObjectRequest *PutObjectRequest) WithOffset(offset int64) *PutObjectRequest {
    putObjectRequest.Offset = &offset
    return putObjectRequest
}


func (putObjectRequest *PutObjectRequest) WithChecksum(contentHash string, checksumType ChecksumType) *PutObjectRequest {
    putObjectRequest.Checksum.ContentHash = contentHash
    putObjectRequest.Checksum.Type = checksumType
    return putObjectRequest
}

func (putObjectRequest *PutObjectRequest) WithMetaData(key string, values ...string) *PutObjectRequest {
    if strings.HasPrefix(strings.ToLower(key), AMZ_META_HEADER) {
        putObjectRequest.Metadata[key] = strings.Join(values, ",")
    } else {
        putObjectRequest.Metadata[strings.ToLower(AMZ_META_HEADER + key)] = strings.Join(values, ",")
    }
    return putObjectRequest
}
