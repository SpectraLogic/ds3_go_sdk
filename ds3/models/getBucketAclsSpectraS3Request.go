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

type GetBucketAclsSpectraS3Request struct {
    BucketId *string
    GroupId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    Permission BucketAclPermission
    UserId *string
}

func NewGetBucketAclsSpectraS3Request() *GetBucketAclsSpectraS3Request {
    return &GetBucketAclsSpectraS3Request{
    }
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithBucketId(bucketId string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.BucketId = &bucketId
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithGroupId(groupId string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.GroupId = &groupId
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithLastPage() *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.LastPage = true
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPageLength(pageLength int) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.PageLength = &pageLength
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPageOffset(pageOffset int) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.PageOffset = &pageOffset
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPermission(permission BucketAclPermission) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.Permission = permission
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithUserId(userId string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.UserId = &userId
    return getBucketAclsSpectraS3Request
}

