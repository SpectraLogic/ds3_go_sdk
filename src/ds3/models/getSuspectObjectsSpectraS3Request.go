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

type GetSuspectObjectsSpectraS3Request struct {
    BucketId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetSuspectObjectsSpectraS3Request() *GetSuspectObjectsSpectraS3Request {
    return &GetSuspectObjectsSpectraS3Request{
    }
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithBucketId(bucketId string) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.BucketId = &bucketId
    return getSuspectObjectsSpectraS3Request
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithLastPage() *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.LastPage = true
    return getSuspectObjectsSpectraS3Request
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.PageLength = &pageLength
    return getSuspectObjectsSpectraS3Request
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.PageOffset = &pageOffset
    return getSuspectObjectsSpectraS3Request
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectObjectsSpectraS3Request
}

