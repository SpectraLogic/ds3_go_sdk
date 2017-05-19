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
    "fmt"
    "strconv"
)

type rangeHeader struct {
    start, end int
}

type GetObjectRequest struct {
    bucketName string
    objectName string
    checksum networking.Checksum
    job string
    offset int64
    rangeHeader *rangeHeader
    queryParams *url.Values
}

func NewGetObjectRequest(bucketName string, objectName string) *GetObjectRequest {
    queryParams := &url.Values{}

    return &GetObjectRequest{
        bucketName: bucketName,
        objectName: objectName,
        checksum: networking.NewNoneChecksum(),
        queryParams: queryParams,
    }
}

func (getObjectRequest *GetObjectRequest) WithJob(job string) *GetObjectRequest {
    getObjectRequest.job = job
    getObjectRequest.queryParams.Set("job", job)
    return getObjectRequest
}
func (getObjectRequest *GetObjectRequest) WithOffset(offset int64) *GetObjectRequest {
    getObjectRequest.offset = offset
    getObjectRequest.queryParams.Set("offset", strconv.FormatInt(offset, 10))
    return getObjectRequest
}



func (GetObjectRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (getObjectRequest *GetObjectRequest) Path() string {
    return "/" + getObjectRequest.bucketName + "/" + getObjectRequest.objectName
}

func (getObjectRequest *GetObjectRequest) QueryParams() *url.Values {
    return getObjectRequest.queryParams
}

func (getObjectRequest *GetObjectRequest) WithChecksum(contentHash string, checksumType networking.ChecksumType) *GetObjectRequest {
    getObjectRequest.checksum.ContentHash = contentHash
    getObjectRequest.checksum.Type = checksumType
    return getObjectRequest
}

func (getObjectRequest *GetObjectRequest) GetChecksum() networking.Checksum {
    return getObjectRequest.checksum
}

func (GetObjectRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}

func (getObjectRequest *GetObjectRequest) WithRange(start, end int) *GetObjectRequest {
    getObjectRequest.rangeHeader = &rangeHeader{start, end}
    return getObjectRequest
}

func (getObjectRequest *GetObjectRequest) Header() *http.Header {
    if getObjectRequest.rangeHeader == nil {
        return &http.Header{}
    } else {
        rng := fmt.Sprintf("bytes=%d-%d", getObjectRequest.rangeHeader.start, getObjectRequest.rangeHeader.end)
        return &http.Header{ "Range": []string{ rng } }
    }
}
