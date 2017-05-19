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

type GetSuspectBlobAzureTargetsSpectraS3Request struct {
    blobId string
    pageLength int
    pageOffset int
    pageStartMarker string
    targetId string
    queryParams *url.Values
}

func NewGetSuspectBlobAzureTargetsSpectraS3Request() *GetSuspectBlobAzureTargetsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetSuspectBlobAzureTargetsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.blobId = blobId
    getSuspectBlobAzureTargetsSpectraS3Request.queryParams.Set("blob_id", blobId)
    return getSuspectBlobAzureTargetsSpectraS3Request
}
func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.pageLength = pageLength
    getSuspectBlobAzureTargetsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getSuspectBlobAzureTargetsSpectraS3Request
}
func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.pageOffset = pageOffset
    getSuspectBlobAzureTargetsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getSuspectBlobAzureTargetsSpectraS3Request
}
func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.pageStartMarker = pageStartMarker
    getSuspectBlobAzureTargetsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getSuspectBlobAzureTargetsSpectraS3Request
}
func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithTargetId(targetId string) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.targetId = targetId
    getSuspectBlobAzureTargetsSpectraS3Request.queryParams.Set("target_id", targetId)
    return getSuspectBlobAzureTargetsSpectraS3Request
}


func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithLastPage() *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.queryParams.Set("last_page", "")
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (GetSuspectBlobAzureTargetsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_azure_target"
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) QueryParams() *url.Values {
    return getSuspectBlobAzureTargetsSpectraS3Request.queryParams
}

func (GetSuspectBlobAzureTargetsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetSuspectBlobAzureTargetsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetSuspectBlobAzureTargetsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
