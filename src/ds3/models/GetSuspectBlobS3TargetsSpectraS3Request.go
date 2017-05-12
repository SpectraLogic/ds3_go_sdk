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

type GetSuspectBlobS3TargetsSpectraS3Request struct {
    blobId string
    pageLength int
    pageOffset int
    pageStartMarker string
    targetId string
    queryParams *url.Values
}

func NewGetSuspectBlobS3TargetsSpectraS3Request() *GetSuspectBlobS3TargetsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetSuspectBlobS3TargetsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.blobId = blobId
    getSuspectBlobS3TargetsSpectraS3Request.queryParams.Set("blob_id", blobId)
    return getSuspectBlobS3TargetsSpectraS3Request
}
func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.pageLength = pageLength
    getSuspectBlobS3TargetsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getSuspectBlobS3TargetsSpectraS3Request
}
func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.pageOffset = pageOffset
    getSuspectBlobS3TargetsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getSuspectBlobS3TargetsSpectraS3Request
}
func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.pageStartMarker = pageStartMarker
    getSuspectBlobS3TargetsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getSuspectBlobS3TargetsSpectraS3Request
}
func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithTargetId(targetId string) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.targetId = targetId
    getSuspectBlobS3TargetsSpectraS3Request.queryParams.Set("target_id", targetId)
    return getSuspectBlobS3TargetsSpectraS3Request
}


func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithLastPage() *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.queryParams.Set("last_page", "")
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (GetSuspectBlobS3TargetsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_s3_target"
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) QueryParams() *url.Values {
    return getSuspectBlobS3TargetsSpectraS3Request.queryParams
}

func (GetSuspectBlobS3TargetsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetSuspectBlobS3TargetsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetSuspectBlobS3TargetsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
