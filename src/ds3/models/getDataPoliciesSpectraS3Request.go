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

import (
    "net/url"
    "net/http"
    "ds3/networking"
    "strconv"
)

type GetDataPoliciesSpectraS3Request struct {
    alwaysForcePutJobCreation bool
    alwaysMinimizeSpanningAcrossMedia bool
    checksumType ChecksumType
    endToEndCrcRequired bool
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    queryParams *url.Values
}

func NewGetDataPoliciesSpectraS3Request() *GetDataPoliciesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDataPoliciesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithAlwaysForcePutJobCreation(alwaysForcePutJobCreation bool) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.alwaysForcePutJobCreation = alwaysForcePutJobCreation
    getDataPoliciesSpectraS3Request.queryParams.Set("always_force_put_job_creation", strconv.FormatBool(alwaysForcePutJobCreation))
    return getDataPoliciesSpectraS3Request
}
func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithAlwaysMinimizeSpanningAcrossMedia(alwaysMinimizeSpanningAcrossMedia bool) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.alwaysMinimizeSpanningAcrossMedia = alwaysMinimizeSpanningAcrossMedia
    getDataPoliciesSpectraS3Request.queryParams.Set("always_minimize_spanning_across_media", strconv.FormatBool(alwaysMinimizeSpanningAcrossMedia))
    return getDataPoliciesSpectraS3Request
}
func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithChecksumType(checksumType ChecksumType) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.checksumType = checksumType
    getDataPoliciesSpectraS3Request.queryParams.Set("checksum_type", checksumType.String())
    return getDataPoliciesSpectraS3Request
}
func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithEndToEndCrcRequired(endToEndCrcRequired bool) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.endToEndCrcRequired = endToEndCrcRequired
    getDataPoliciesSpectraS3Request.queryParams.Set("end_to_end_crc_required", strconv.FormatBool(endToEndCrcRequired))
    return getDataPoliciesSpectraS3Request
}
func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithPageLength(pageLength int) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.pageLength = pageLength
    getDataPoliciesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDataPoliciesSpectraS3Request
}
func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithPageOffset(pageOffset int) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.pageOffset = pageOffset
    getDataPoliciesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDataPoliciesSpectraS3Request
}
func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.pageStartMarker = pageStartMarker
    getDataPoliciesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithName(name *string) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.name = name
    if name != nil {
        getDataPoliciesSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getDataPoliciesSpectraS3Request.queryParams.Set("name", "")
    }
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithLastPage() *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.queryParams.Set("last_page", "")
    return getDataPoliciesSpectraS3Request
}

func (GetDataPoliciesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) Path() string {
    return "/_rest_/data_policy"
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) QueryParams() *url.Values {
    return getDataPoliciesSpectraS3Request.queryParams
}

func (GetDataPoliciesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDataPoliciesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDataPoliciesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
