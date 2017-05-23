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

type GetAzureTargetBucketNamesSpectraS3Request struct {
    bucketId string
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    targetId string
    queryParams *url.Values
}

func NewGetAzureTargetBucketNamesSpectraS3Request() *GetAzureTargetBucketNamesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetAzureTargetBucketNamesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithBucketId(bucketId string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.bucketId = bucketId
    getAzureTargetBucketNamesSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getAzureTargetBucketNamesSpectraS3Request
}
func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.pageLength = pageLength
    getAzureTargetBucketNamesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getAzureTargetBucketNamesSpectraS3Request
}
func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.pageOffset = pageOffset
    getAzureTargetBucketNamesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getAzureTargetBucketNamesSpectraS3Request
}
func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.pageStartMarker = pageStartMarker
    getAzureTargetBucketNamesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getAzureTargetBucketNamesSpectraS3Request
}
func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithTargetId(targetId string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.targetId = targetId
    getAzureTargetBucketNamesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithName(name *string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.name = name
    if name != nil {
        getAzureTargetBucketNamesSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getAzureTargetBucketNamesSpectraS3Request.queryParams.Set("name", "")
    }
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithLastPage() *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.queryParams.Set("last_page", "")
    return getAzureTargetBucketNamesSpectraS3Request
}

func (GetAzureTargetBucketNamesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) Path() string {
    return "/_rest_/azure_target_bucket_name"
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) QueryParams() *url.Values {
    return getAzureTargetBucketNamesSpectraS3Request.queryParams
}

func (GetAzureTargetBucketNamesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetAzureTargetBucketNamesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetAzureTargetBucketNamesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
