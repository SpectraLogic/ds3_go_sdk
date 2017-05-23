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
    "strconv"
)

type GetObjectsDetailsSpectraS3Request struct {
    bucketId string
    folder *string
    latest bool
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    s3ObjectType S3ObjectType
    version int64
    queryParams *url.Values
}

func NewGetObjectsDetailsSpectraS3Request() *GetObjectsDetailsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetObjectsDetailsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithBucketId(bucketId string) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.bucketId = bucketId
    getObjectsDetailsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getObjectsDetailsSpectraS3Request
}
func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithLatest(latest bool) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.latest = latest
    getObjectsDetailsSpectraS3Request.queryParams.Set("latest", strconv.FormatBool(latest))
    return getObjectsDetailsSpectraS3Request
}
func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithPageLength(pageLength int) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.pageLength = pageLength
    getObjectsDetailsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getObjectsDetailsSpectraS3Request
}
func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.pageOffset = pageOffset
    getObjectsDetailsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getObjectsDetailsSpectraS3Request
}
func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.pageStartMarker = pageStartMarker
    getObjectsDetailsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getObjectsDetailsSpectraS3Request
}
func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithS3ObjectType(s3ObjectType S3ObjectType) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.s3ObjectType = s3ObjectType
    getObjectsDetailsSpectraS3Request.queryParams.Set("type", s3ObjectType.String())
    return getObjectsDetailsSpectraS3Request
}
func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithVersion(version int64) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.version = version
    getObjectsDetailsSpectraS3Request.queryParams.Set("version", strconv.FormatInt(version, 10))
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithFolder(folder *string) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.folder = folder
    if folder != nil {
        getObjectsDetailsSpectraS3Request.queryParams.Set("folder", *folder)
    } else {
        getObjectsDetailsSpectraS3Request.queryParams.Set("folder", "")
    }
    return getObjectsDetailsSpectraS3Request
}
func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithName(name *string) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.name = name
    if name != nil {
        getObjectsDetailsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getObjectsDetailsSpectraS3Request.queryParams.Set("name", "")
    }
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithLastPage() *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.queryParams.Set("last_page", "")
    return getObjectsDetailsSpectraS3Request
}

func (GetObjectsDetailsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) Path() string {
    return "/_rest_/object"
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) QueryParams() *url.Values {
    return getObjectsDetailsSpectraS3Request.queryParams
}

func (GetObjectsDetailsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetObjectsDetailsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetObjectsDetailsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
