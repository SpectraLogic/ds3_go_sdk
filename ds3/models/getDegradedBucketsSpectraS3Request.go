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

type GetDegradedBucketsSpectraS3Request struct {
    DataPolicyId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetDegradedBucketsSpectraS3Request() *GetDegradedBucketsSpectraS3Request {
    return &GetDegradedBucketsSpectraS3Request{
    }
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithLastPage() *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.LastPage = true
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithName(name string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.Name = &name
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithPageLength(pageLength int) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.PageLength = &pageLength
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.PageOffset = &pageOffset
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithUserId(userId string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.UserId = &userId
    return getDegradedBucketsSpectraS3Request
}

