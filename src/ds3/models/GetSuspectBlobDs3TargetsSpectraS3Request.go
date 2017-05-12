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

type GetSuspectBlobDs3TargetsSpectraS3Request struct {
    blobId string
    pageLength int
    pageOffset int
    pageStartMarker string
    targetId string
    queryParams *url.Values
}

func NewGetSuspectBlobDs3TargetsSpectraS3Request() *GetSuspectBlobDs3TargetsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetSuspectBlobDs3TargetsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.blobId = blobId
    getSuspectBlobDs3TargetsSpectraS3Request.queryParams.Set("blob_id", blobId)
    return getSuspectBlobDs3TargetsSpectraS3Request
}
func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.pageLength = pageLength
    getSuspectBlobDs3TargetsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getSuspectBlobDs3TargetsSpectraS3Request
}
func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.pageOffset = pageOffset
    getSuspectBlobDs3TargetsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getSuspectBlobDs3TargetsSpectraS3Request
}
func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.pageStartMarker = pageStartMarker
    getSuspectBlobDs3TargetsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getSuspectBlobDs3TargetsSpectraS3Request
}
func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithTargetId(targetId string) *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.targetId = targetId
    getSuspectBlobDs3TargetsSpectraS3Request.queryParams.Set("target_id", targetId)
    return getSuspectBlobDs3TargetsSpectraS3Request
}


func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithLastPage() *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.queryParams.Set("last_page", "")
    return getSuspectBlobDs3TargetsSpectraS3Request
}

func (GetSuspectBlobDs3TargetsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_ds3_target"
}

func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) QueryParams() *url.Values {
    return getSuspectBlobDs3TargetsSpectraS3Request.queryParams
}

func (GetSuspectBlobDs3TargetsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetSuspectBlobDs3TargetsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetSuspectBlobDs3TargetsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
