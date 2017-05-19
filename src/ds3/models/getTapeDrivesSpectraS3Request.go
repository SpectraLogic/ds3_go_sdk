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

type GetTapeDrivesSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    partitionId string
    serialNumber *string
    state TapeDriveState
    tapeDriveType TapeDriveType
    queryParams *url.Values
}

func NewGetTapeDrivesSpectraS3Request() *GetTapeDrivesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapeDrivesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPageLength(pageLength int) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.pageLength = pageLength
    getTapeDrivesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getTapeDrivesSpectraS3Request
}
func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.pageOffset = pageOffset
    getTapeDrivesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getTapeDrivesSpectraS3Request
}
func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.pageStartMarker = pageStartMarker
    getTapeDrivesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getTapeDrivesSpectraS3Request
}
func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPartitionId(partitionId string) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.partitionId = partitionId
    getTapeDrivesSpectraS3Request.queryParams.Set("partition_id", partitionId)
    return getTapeDrivesSpectraS3Request
}
func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithState(state TapeDriveState) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.state = state
    getTapeDrivesSpectraS3Request.queryParams.Set("state", state.String())
    return getTapeDrivesSpectraS3Request
}
func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithTapeDriveType(tapeDriveType TapeDriveType) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.tapeDriveType = tapeDriveType
    getTapeDrivesSpectraS3Request.queryParams.Set("type", tapeDriveType.String())
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithSerialNumber(serialNumber *string) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.serialNumber = serialNumber
    if serialNumber != nil {
        getTapeDrivesSpectraS3Request.queryParams.Set("serial_number", *serialNumber)
    } else {
        getTapeDrivesSpectraS3Request.queryParams.Set("serial_number", "")
    }
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithLastPage() *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.queryParams.Set("last_page", "")
    return getTapeDrivesSpectraS3Request
}

func (GetTapeDrivesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) Path() string {
    return "/_rest_/tape_drive"
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) QueryParams() *url.Values {
    return getTapeDrivesSpectraS3Request.queryParams
}

func (GetTapeDrivesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapeDrivesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapeDrivesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
