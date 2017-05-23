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

type GetSuspectBucketsSpectraS3Request struct {
    dataPolicyId string
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetSuspectBucketsSpectraS3Request() *GetSuspectBucketsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetSuspectBucketsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.dataPolicyId = dataPolicyId
    getSuspectBucketsSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getSuspectBucketsSpectraS3Request
}
func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.pageLength = pageLength
    getSuspectBucketsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getSuspectBucketsSpectraS3Request
}
func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.pageOffset = pageOffset
    getSuspectBucketsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getSuspectBucketsSpectraS3Request
}
func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.pageStartMarker = pageStartMarker
    getSuspectBucketsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getSuspectBucketsSpectraS3Request
}
func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithUserId(userId string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.userId = userId
    getSuspectBucketsSpectraS3Request.queryParams.Set("user_id", userId)
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithName(name *string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.name = name
    if name != nil {
        getSuspectBucketsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getSuspectBucketsSpectraS3Request.queryParams.Set("name", "")
    }
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithLastPage() *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.queryParams.Set("last_page", "")
    return getSuspectBucketsSpectraS3Request
}

func (GetSuspectBucketsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) Path() string {
    return "/_rest_/suspect_bucket"
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) QueryParams() *url.Values {
    return getSuspectBucketsSpectraS3Request.queryParams
}

func (GetSuspectBucketsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetSuspectBucketsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetSuspectBucketsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
