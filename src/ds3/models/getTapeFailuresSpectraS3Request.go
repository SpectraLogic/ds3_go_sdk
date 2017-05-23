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

type GetTapeFailuresSpectraS3Request struct {
    errorMessage *string
    pageLength int
    pageOffset int
    pageStartMarker string
    tapeDriveId string
    tapeFailureType TapeFailureType
    tapeId string
    queryParams *url.Values
}

func NewGetTapeFailuresSpectraS3Request() *GetTapeFailuresSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapeFailuresSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithPageLength(pageLength int) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.pageLength = pageLength
    getTapeFailuresSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getTapeFailuresSpectraS3Request
}
func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.pageOffset = pageOffset
    getTapeFailuresSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getTapeFailuresSpectraS3Request
}
func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.pageStartMarker = pageStartMarker
    getTapeFailuresSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getTapeFailuresSpectraS3Request
}
func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithTapeDriveId(tapeDriveId string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.tapeDriveId = tapeDriveId
    getTapeFailuresSpectraS3Request.queryParams.Set("tape_drive_id", tapeDriveId)
    return getTapeFailuresSpectraS3Request
}
func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithTapeId(tapeId string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.tapeId = tapeId
    getTapeFailuresSpectraS3Request.queryParams.Set("tape_id", tapeId)
    return getTapeFailuresSpectraS3Request
}
func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithTapeFailureType(tapeFailureType TapeFailureType) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.tapeFailureType = tapeFailureType
    getTapeFailuresSpectraS3Request.queryParams.Set("type", tapeFailureType.String())
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithErrorMessage(errorMessage *string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.errorMessage = errorMessage
    if errorMessage != nil {
        getTapeFailuresSpectraS3Request.queryParams.Set("error_message", *errorMessage)
    } else {
        getTapeFailuresSpectraS3Request.queryParams.Set("error_message", "")
    }
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithLastPage() *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.queryParams.Set("last_page", "")
    return getTapeFailuresSpectraS3Request
}

func (GetTapeFailuresSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) Path() string {
    return "/_rest_/tape_failure"
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) QueryParams() *url.Values {
    return getTapeFailuresSpectraS3Request.queryParams
}

func (GetTapeFailuresSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapeFailuresSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapeFailuresSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
