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

type GetBucketRequest struct {
    BucketName string
    Delimiter *string
    Marker *string
    MaxKeys *int
    Prefix *string
}

func NewGetBucketRequest(bucketName string) *GetBucketRequest {
    return &GetBucketRequest{
        BucketName: bucketName,
    }
}

func (getBucketRequest *GetBucketRequest) WithDelimiter(delimiter string) *GetBucketRequest {
    getBucketRequest.Delimiter = &delimiter
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithMarker(marker string) *GetBucketRequest {
    getBucketRequest.Marker = &marker
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithMaxKeys(maxKeys int) *GetBucketRequest {
    getBucketRequest.MaxKeys = &maxKeys
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithPrefix(prefix string) *GetBucketRequest {
    getBucketRequest.Prefix = &prefix
    return getBucketRequest
}

