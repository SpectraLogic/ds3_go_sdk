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
    "strings"
    "strconv"
)

const ( AMZ_META_HEADER = "x-amz-meta-" )

type PutObjectRequest struct {
    bucketName string
    objectName string
    checksum networking.Checksum
    content networking.ReaderWithSizeDecorator
    headers *http.Header
    job string
    offset int64
    queryParams *url.Values
}

func NewPutObjectRequest(bucketName string, objectName string, content networking.ReaderWithSizeDecorator) *PutObjectRequest {
    queryParams := &url.Values{}

    return &PutObjectRequest{
        bucketName: bucketName,
        objectName: objectName,
        content: content,
        checksum: networking.NewNoneChecksum(),
        headers: &http.Header{},
        queryParams: queryParams,
    }
}

func (putObjectRequest *PutObjectRequest) WithJob(job string) *PutObjectRequest {
    putObjectRequest.job = job
    putObjectRequest.queryParams.Set("job", job)
    return putObjectRequest
}
func (putObjectRequest *PutObjectRequest) WithOffset(offset int64) *PutObjectRequest {
    putObjectRequest.offset = offset
    putObjectRequest.queryParams.Set("offset", strconv.FormatInt(offset, 10))
    return putObjectRequest
}



func (PutObjectRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (putObjectRequest *PutObjectRequest) Path() string {
    return "/" + putObjectRequest.bucketName + "/" + putObjectRequest.objectName
}

func (putObjectRequest *PutObjectRequest) QueryParams() *url.Values {
    return putObjectRequest.queryParams
}

func (putObjectRequest *PutObjectRequest) WithChecksum(contentHash string, checksumType networking.ChecksumType) *PutObjectRequest {
    putObjectRequest.checksum.ContentHash = contentHash
    putObjectRequest.checksum.Type = checksumType
    return putObjectRequest
}

func (putObjectRequest *PutObjectRequest) GetChecksum() networking.Checksum {
    return putObjectRequest.checksum
}

func (putObjectRequest *PutObjectRequest) WithMetaData(key string, value string) *PutObjectRequest {
    if strings.HasPrefix(strings.ToLower(key), AMZ_META_HEADER) {
        putObjectRequest.headers.Add(strings.ToLower(key), value)
    } else {
        putObjectRequest.headers.Add(strings.ToLower(AMZ_META_HEADER + key), value)
    }
    return putObjectRequest
}

func (putObjectRequest *PutObjectRequest) Header() *http.Header {
    return putObjectRequest.headers
}

func (putObjectRequest *PutObjectRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return putObjectRequest.content
}
