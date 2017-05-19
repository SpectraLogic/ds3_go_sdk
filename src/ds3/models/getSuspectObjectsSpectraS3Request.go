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

type GetSuspectObjectsSpectraS3Request struct {
    bucketId string
    pageLength int
    pageOffset int
    pageStartMarker string
    queryParams *url.Values
}

func NewGetSuspectObjectsSpectraS3Request() *GetSuspectObjectsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetSuspectObjectsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithBucketId(bucketId string) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.bucketId = bucketId
    getSuspectObjectsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getSuspectObjectsSpectraS3Request
}
func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.pageLength = pageLength
    getSuspectObjectsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getSuspectObjectsSpectraS3Request
}
func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.pageOffset = pageOffset
    getSuspectObjectsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getSuspectObjectsSpectraS3Request
}
func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.pageStartMarker = pageStartMarker
    getSuspectObjectsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getSuspectObjectsSpectraS3Request
}


func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithLastPage() *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.queryParams.Set("last_page", "")
    return getSuspectObjectsSpectraS3Request
}

func (GetSuspectObjectsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) Path() string {
    return "/_rest_/suspect_object"
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) QueryParams() *url.Values {
    return getSuspectObjectsSpectraS3Request.queryParams
}

func (GetSuspectObjectsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetSuspectObjectsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetSuspectObjectsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
