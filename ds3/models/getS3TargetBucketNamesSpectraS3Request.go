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

type GetS3TargetBucketNamesSpectraS3Request struct {
    BucketId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetId *string
}

func NewGetS3TargetBucketNamesSpectraS3Request() *GetS3TargetBucketNamesSpectraS3Request {
    return &GetS3TargetBucketNamesSpectraS3Request{
    }
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithBucketId(bucketId string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.BucketId = &bucketId
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithLastPage() *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.LastPage = true
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithName(name string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.Name = &name
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.PageLength = &pageLength
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.PageOffset = &pageOffset
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithTargetId(targetId string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.TargetId = &targetId
    return getS3TargetBucketNamesSpectraS3Request
}

