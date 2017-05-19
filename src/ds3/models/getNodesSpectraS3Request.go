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

type GetNodesSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    queryParams *url.Values
}

func NewGetNodesSpectraS3Request() *GetNodesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetNodesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithPageLength(pageLength int) *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.pageLength = pageLength
    getNodesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getNodesSpectraS3Request
}
func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithPageOffset(pageOffset int) *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.pageOffset = pageOffset
    getNodesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getNodesSpectraS3Request
}
func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.pageStartMarker = pageStartMarker
    getNodesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getNodesSpectraS3Request
}


func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithLastPage() *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.queryParams.Set("last_page", "")
    return getNodesSpectraS3Request
}

func (GetNodesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) Path() string {
    return "/_rest_/node"
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) QueryParams() *url.Values {
    return getNodesSpectraS3Request.queryParams
}

func (GetNodesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetNodesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetNodesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
