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

type GetAzureTargetBucketNamesSpectraS3Request struct {
    BucketId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetId *string
}

func NewGetAzureTargetBucketNamesSpectraS3Request() *GetAzureTargetBucketNamesSpectraS3Request {
    return &GetAzureTargetBucketNamesSpectraS3Request{
    }
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithBucketId(bucketId string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.BucketId = &bucketId
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithLastPage() *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.LastPage = true
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithName(name string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.Name = &name
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.PageLength = &pageLength
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.PageOffset = &pageOffset
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithTargetId(targetId string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.TargetId = &targetId
    return getAzureTargetBucketNamesSpectraS3Request
}

