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

type GetGroupsSpectraS3Request struct {
    builtIn bool
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    queryParams *url.Values
}

func NewGetGroupsSpectraS3Request() *GetGroupsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetGroupsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithBuiltIn(builtIn bool) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.builtIn = builtIn
    getGroupsSpectraS3Request.queryParams.Set("built_in", strconv.FormatBool(builtIn))
    return getGroupsSpectraS3Request
}
func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithPageLength(pageLength int) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.pageLength = pageLength
    getGroupsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getGroupsSpectraS3Request
}
func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithPageOffset(pageOffset int) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.pageOffset = pageOffset
    getGroupsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getGroupsSpectraS3Request
}
func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.pageStartMarker = pageStartMarker
    getGroupsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithName(name *string) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.name = name
    if name != nil {
        getGroupsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getGroupsSpectraS3Request.queryParams.Set("name", "")
    }
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithLastPage() *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.queryParams.Set("last_page", "")
    return getGroupsSpectraS3Request
}

func (GetGroupsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) Path() string {
    return "/_rest_/group"
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) QueryParams() *url.Values {
    return getGroupsSpectraS3Request.queryParams
}

func (GetGroupsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetGroupsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetGroupsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
