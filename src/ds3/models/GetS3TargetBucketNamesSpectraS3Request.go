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

type GetS3TargetBucketNamesSpectraS3Request struct {
    bucketId string
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    targetId string
    queryParams *url.Values
}

func NewGetS3TargetBucketNamesSpectraS3Request() *GetS3TargetBucketNamesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetS3TargetBucketNamesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithBucketId(bucketId string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.bucketId = bucketId
    getS3TargetBucketNamesSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getS3TargetBucketNamesSpectraS3Request
}
func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.pageLength = pageLength
    getS3TargetBucketNamesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getS3TargetBucketNamesSpectraS3Request
}
func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.pageOffset = pageOffset
    getS3TargetBucketNamesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getS3TargetBucketNamesSpectraS3Request
}
func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.pageStartMarker = pageStartMarker
    getS3TargetBucketNamesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getS3TargetBucketNamesSpectraS3Request
}
func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithTargetId(targetId string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.targetId = targetId
    getS3TargetBucketNamesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithName(name *string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.name = name
    if name != nil {
        getS3TargetBucketNamesSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getS3TargetBucketNamesSpectraS3Request.queryParams.Set("name", "")
    }
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithLastPage() *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.queryParams.Set("last_page", "")
    return getS3TargetBucketNamesSpectraS3Request
}

func (GetS3TargetBucketNamesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) Path() string {
    return "/_rest_/s3_target_bucket_name"
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) QueryParams() *url.Values {
    return getS3TargetBucketNamesSpectraS3Request.queryParams
}

func (GetS3TargetBucketNamesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetS3TargetBucketNamesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetS3TargetBucketNamesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
