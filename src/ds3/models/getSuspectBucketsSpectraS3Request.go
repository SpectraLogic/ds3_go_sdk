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

type GetSuspectBucketsSpectraS3Request struct {
    DataPolicyId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetSuspectBucketsSpectraS3Request() *GetSuspectBucketsSpectraS3Request {
    return &GetSuspectBucketsSpectraS3Request{
    }
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.DataPolicyId = &dataPolicyId
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithLastPage() *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.LastPage = true
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithName(name string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.Name = &name
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.PageLength = &pageLength
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.PageOffset = &pageOffset
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithUserId(userId string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.UserId = &userId
    return getSuspectBucketsSpectraS3Request
}

