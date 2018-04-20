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

type GetObjectsDetailsSpectraS3Request struct {
    BucketId *string
    EndDate *int64
    LastPage bool
    Latest *bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    S3ObjectType S3ObjectType
    StartDate *int64
}

func NewGetObjectsDetailsSpectraS3Request() *GetObjectsDetailsSpectraS3Request {
    return &GetObjectsDetailsSpectraS3Request{
    }
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithBucketId(bucketId string) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.BucketId = &bucketId
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithEndDate(endDate int64) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.EndDate = &endDate
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithLastPage() *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.LastPage = true
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithLatest(latest bool) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.Latest = &latest
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithName(name string) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.Name = &name
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithPageLength(pageLength int) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.PageLength = &pageLength
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.PageOffset = &pageOffset
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithStartDate(startDate int64) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.StartDate = &startDate
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithS3ObjectType(s3ObjectType S3ObjectType) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.S3ObjectType = s3ObjectType
    return getObjectsDetailsSpectraS3Request
}

