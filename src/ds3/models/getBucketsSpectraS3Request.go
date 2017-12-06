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

type GetBucketsSpectraS3Request struct {
    DataPolicyId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetBucketsSpectraS3Request() *GetBucketsSpectraS3Request {
    return &GetBucketsSpectraS3Request{
    }
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.DataPolicyId = &dataPolicyId
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithLastPage() *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.LastPage = true
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithName(name string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.Name = &name
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithPageLength(pageLength int) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.PageLength = &pageLength
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithPageOffset(pageOffset int) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.PageOffset = &pageOffset
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithUserId(userId string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.UserId = &userId
    return getBucketsSpectraS3Request
}

