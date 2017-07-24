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

type GetObjectsWithFullDetailsSpectraS3Request struct {
    bucketId string
    latest bool
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    s3ObjectType S3ObjectType
    version int64
    queryParams *url.Values
}

func NewGetObjectsWithFullDetailsSpectraS3Request() *GetObjectsWithFullDetailsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("full_details", "")

    return &GetObjectsWithFullDetailsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithBucketId(bucketId string) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.bucketId = bucketId
    getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getObjectsWithFullDetailsSpectraS3Request
}
func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithLatest(latest bool) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.latest = latest
    getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("latest", strconv.FormatBool(latest))
    return getObjectsWithFullDetailsSpectraS3Request
}
func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithPageLength(pageLength int) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.pageLength = pageLength
    getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getObjectsWithFullDetailsSpectraS3Request
}
func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.pageOffset = pageOffset
    getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getObjectsWithFullDetailsSpectraS3Request
}
func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.pageStartMarker = pageStartMarker
    getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getObjectsWithFullDetailsSpectraS3Request
}
func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithS3ObjectType(s3ObjectType S3ObjectType) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.s3ObjectType = s3ObjectType
    getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("type", s3ObjectType.String())
    return getObjectsWithFullDetailsSpectraS3Request
}
func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithVersion(version int64) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.version = version
    getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("version", strconv.FormatInt(version, 10))
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithName(name *string) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.name = name
    if name != nil {
        getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("name", "")
    }
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithIncludePhysicalPlacement() *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("include_physical_placement", "")
    return getObjectsWithFullDetailsSpectraS3Request
}
func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithLastPage() *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.queryParams.Set("last_page", "")
    return getObjectsWithFullDetailsSpectraS3Request
}

func (GetObjectsWithFullDetailsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) Path() string {
    return "/_rest_/object"
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) QueryParams() *url.Values {
    return getObjectsWithFullDetailsSpectraS3Request.queryParams
}

func (GetObjectsWithFullDetailsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetObjectsWithFullDetailsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetObjectsWithFullDetailsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
