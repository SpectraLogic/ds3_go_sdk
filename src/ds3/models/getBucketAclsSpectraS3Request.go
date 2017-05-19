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

type GetBucketAclsSpectraS3Request struct {
    bucketId string
    groupId string
    pageLength int
    pageOffset int
    pageStartMarker string
    permission BucketAclPermission
    userId string
    queryParams *url.Values
}

func NewGetBucketAclsSpectraS3Request() *GetBucketAclsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetBucketAclsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithBucketId(bucketId string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.bucketId = bucketId
    getBucketAclsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getBucketAclsSpectraS3Request
}
func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithGroupId(groupId string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.groupId = groupId
    getBucketAclsSpectraS3Request.queryParams.Set("group_id", groupId)
    return getBucketAclsSpectraS3Request
}
func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPageLength(pageLength int) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.pageLength = pageLength
    getBucketAclsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getBucketAclsSpectraS3Request
}
func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPageOffset(pageOffset int) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.pageOffset = pageOffset
    getBucketAclsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getBucketAclsSpectraS3Request
}
func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.pageStartMarker = pageStartMarker
    getBucketAclsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getBucketAclsSpectraS3Request
}
func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPermission(permission BucketAclPermission) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.permission = permission
    getBucketAclsSpectraS3Request.queryParams.Set("permission", permission.String())
    return getBucketAclsSpectraS3Request
}
func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithUserId(userId string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.userId = userId
    getBucketAclsSpectraS3Request.queryParams.Set("user_id", userId)
    return getBucketAclsSpectraS3Request
}


func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithLastPage() *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.queryParams.Set("last_page", "")
    return getBucketAclsSpectraS3Request
}

func (GetBucketAclsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) Path() string {
    return "/_rest_/bucket_acl"
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) QueryParams() *url.Values {
    return getBucketAclsSpectraS3Request.queryParams
}

func (GetBucketAclsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetBucketAclsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetBucketAclsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
