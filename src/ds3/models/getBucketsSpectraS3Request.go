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

type GetBucketsSpectraS3Request struct {
    dataPolicyId string
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetBucketsSpectraS3Request() *GetBucketsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetBucketsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.dataPolicyId = dataPolicyId
    getBucketsSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getBucketsSpectraS3Request
}
func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithPageLength(pageLength int) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.pageLength = pageLength
    getBucketsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getBucketsSpectraS3Request
}
func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithPageOffset(pageOffset int) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.pageOffset = pageOffset
    getBucketsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getBucketsSpectraS3Request
}
func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.pageStartMarker = pageStartMarker
    getBucketsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getBucketsSpectraS3Request
}
func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithUserId(userId string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.userId = userId
    getBucketsSpectraS3Request.queryParams.Set("user_id", userId)
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithName(name *string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.name = name
    if name != nil {
        getBucketsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getBucketsSpectraS3Request.queryParams.Set("name", "")
    }
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithLastPage() *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.queryParams.Set("last_page", "")
    return getBucketsSpectraS3Request
}

func (GetBucketsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) Path() string {
    return "/_rest_/bucket"
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) QueryParams() *url.Values {
    return getBucketsSpectraS3Request.queryParams
}

func (GetBucketsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetBucketsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetBucketsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
