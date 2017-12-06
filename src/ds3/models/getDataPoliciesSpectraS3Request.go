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

type GetDataPoliciesSpectraS3Request struct {
    AlwaysForcePutJobCreation *bool
    AlwaysMinimizeSpanningAcrossMedia *bool
    ChecksumType ChecksumType
    EndToEndCrcRequired *bool
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetDataPoliciesSpectraS3Request() *GetDataPoliciesSpectraS3Request {
    return &GetDataPoliciesSpectraS3Request{
    }
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithAlwaysForcePutJobCreation(alwaysForcePutJobCreation bool) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.AlwaysForcePutJobCreation = &alwaysForcePutJobCreation
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithAlwaysMinimizeSpanningAcrossMedia(alwaysMinimizeSpanningAcrossMedia bool) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.AlwaysMinimizeSpanningAcrossMedia = &alwaysMinimizeSpanningAcrossMedia
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithChecksumType(checksumType ChecksumType) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.ChecksumType = checksumType
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithEndToEndCrcRequired(endToEndCrcRequired bool) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.EndToEndCrcRequired = &endToEndCrcRequired
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithLastPage() *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.LastPage = true
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithName(name string) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.Name = &name
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithPageLength(pageLength int) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.PageLength = &pageLength
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithPageOffset(pageOffset int) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.PageOffset = &pageOffset
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDataPoliciesSpectraS3Request
}

