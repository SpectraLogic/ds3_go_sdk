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

type GetTapeLibrariesSpectraS3Request struct {
    managementUrl *string
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    serialNumber *string
    queryParams *url.Values
}

func NewGetTapeLibrariesSpectraS3Request() *GetTapeLibrariesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapeLibrariesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithPageLength(pageLength int) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.pageLength = pageLength
    getTapeLibrariesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getTapeLibrariesSpectraS3Request
}
func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.pageOffset = pageOffset
    getTapeLibrariesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getTapeLibrariesSpectraS3Request
}
func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.pageStartMarker = pageStartMarker
    getTapeLibrariesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithManagementUrl(managementUrl *string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.managementUrl = managementUrl
    if managementUrl != nil {
        getTapeLibrariesSpectraS3Request.queryParams.Set("management_url", *managementUrl)
    } else {
        getTapeLibrariesSpectraS3Request.queryParams.Set("management_url", "")
    }
    return getTapeLibrariesSpectraS3Request
}
func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithName(name *string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.name = name
    if name != nil {
        getTapeLibrariesSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getTapeLibrariesSpectraS3Request.queryParams.Set("name", "")
    }
    return getTapeLibrariesSpectraS3Request
}
func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithSerialNumber(serialNumber *string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.serialNumber = serialNumber
    if serialNumber != nil {
        getTapeLibrariesSpectraS3Request.queryParams.Set("serial_number", *serialNumber)
    } else {
        getTapeLibrariesSpectraS3Request.queryParams.Set("serial_number", "")
    }
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithLastPage() *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.queryParams.Set("last_page", "")
    return getTapeLibrariesSpectraS3Request
}

func (GetTapeLibrariesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) Path() string {
    return "/_rest_/tape_library"
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) QueryParams() *url.Values {
    return getTapeLibrariesSpectraS3Request.queryParams
}

func (GetTapeLibrariesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapeLibrariesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapeLibrariesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
