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

type GetSuspectBlobTapesSpectraS3Request struct {
    blobId string
    pageLength int
    pageOffset int
    pageStartMarker string
    tapeId string
    queryParams *url.Values
}

func NewGetSuspectBlobTapesSpectraS3Request() *GetSuspectBlobTapesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetSuspectBlobTapesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.blobId = blobId
    getSuspectBlobTapesSpectraS3Request.queryParams.Set("blob_id", blobId)
    return getSuspectBlobTapesSpectraS3Request
}
func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.pageLength = pageLength
    getSuspectBlobTapesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getSuspectBlobTapesSpectraS3Request
}
func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.pageOffset = pageOffset
    getSuspectBlobTapesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getSuspectBlobTapesSpectraS3Request
}
func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.pageStartMarker = pageStartMarker
    getSuspectBlobTapesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getSuspectBlobTapesSpectraS3Request
}
func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithTapeId(tapeId string) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.tapeId = tapeId
    getSuspectBlobTapesSpectraS3Request.queryParams.Set("tape_id", tapeId)
    return getSuspectBlobTapesSpectraS3Request
}


func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithLastPage() *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.queryParams.Set("last_page", "")
    return getSuspectBlobTapesSpectraS3Request
}

func (GetSuspectBlobTapesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_tape"
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) QueryParams() *url.Values {
    return getSuspectBlobTapesSpectraS3Request.queryParams
}

func (GetSuspectBlobTapesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetSuspectBlobTapesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetSuspectBlobTapesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
