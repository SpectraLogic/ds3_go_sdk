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

type GetObjectsWithFullDetailsSpectraS3Request struct {
    BucketId *string
    EndDate *int64
    IncludePhysicalPlacement bool
    LastPage bool
    Latest *bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    S3ObjectType S3ObjectType
    StartDate *int64
}

func NewGetObjectsWithFullDetailsSpectraS3Request() *GetObjectsWithFullDetailsSpectraS3Request {
    return &GetObjectsWithFullDetailsSpectraS3Request{
    }
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithBucketId(bucketId string) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.BucketId = &bucketId
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithEndDate(endDate int64) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.EndDate = &endDate
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithIncludePhysicalPlacement() *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.IncludePhysicalPlacement = true
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithLastPage() *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.LastPage = true
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithLatest(latest bool) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.Latest = &latest
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithName(name string) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.Name = &name
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithPageLength(pageLength int) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.PageLength = &pageLength
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.PageOffset = &pageOffset
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithStartDate(startDate int64) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.StartDate = &startDate
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithS3ObjectType(s3ObjectType S3ObjectType) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.S3ObjectType = s3ObjectType
    return getObjectsWithFullDetailsSpectraS3Request
}

