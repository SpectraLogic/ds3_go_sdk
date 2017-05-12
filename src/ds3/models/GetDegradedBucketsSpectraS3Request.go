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

type GetDegradedBucketsSpectraS3Request struct {
    dataPolicyId string
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetDegradedBucketsSpectraS3Request() *GetDegradedBucketsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDegradedBucketsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.dataPolicyId = dataPolicyId
    getDegradedBucketsSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getDegradedBucketsSpectraS3Request
}
func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithPageLength(pageLength int) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.pageLength = pageLength
    getDegradedBucketsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDegradedBucketsSpectraS3Request
}
func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.pageOffset = pageOffset
    getDegradedBucketsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDegradedBucketsSpectraS3Request
}
func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.pageStartMarker = pageStartMarker
    getDegradedBucketsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDegradedBucketsSpectraS3Request
}
func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithUserId(userId string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.userId = userId
    getDegradedBucketsSpectraS3Request.queryParams.Set("user_id", userId)
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithName(name *string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.name = name
    if name != nil {
        getDegradedBucketsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getDegradedBucketsSpectraS3Request.queryParams.Set("name", "")
    }
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithLastPage() *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.queryParams.Set("last_page", "")
    return getDegradedBucketsSpectraS3Request
}

func (GetDegradedBucketsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) Path() string {
    return "/_rest_/degraded_bucket"
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) QueryParams() *url.Values {
    return getDegradedBucketsSpectraS3Request.queryParams
}

func (GetDegradedBucketsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDegradedBucketsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDegradedBucketsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
