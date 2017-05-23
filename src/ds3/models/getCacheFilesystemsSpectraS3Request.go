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

type GetCacheFilesystemsSpectraS3Request struct {
    nodeId string
    pageLength int
    pageOffset int
    pageStartMarker string
    queryParams *url.Values
}

func NewGetCacheFilesystemsSpectraS3Request() *GetCacheFilesystemsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetCacheFilesystemsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) WithNodeId(nodeId string) *GetCacheFilesystemsSpectraS3Request {
    getCacheFilesystemsSpectraS3Request.nodeId = nodeId
    getCacheFilesystemsSpectraS3Request.queryParams.Set("node_id", nodeId)
    return getCacheFilesystemsSpectraS3Request
}
func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) WithPageLength(pageLength int) *GetCacheFilesystemsSpectraS3Request {
    getCacheFilesystemsSpectraS3Request.pageLength = pageLength
    getCacheFilesystemsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getCacheFilesystemsSpectraS3Request
}
func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) WithPageOffset(pageOffset int) *GetCacheFilesystemsSpectraS3Request {
    getCacheFilesystemsSpectraS3Request.pageOffset = pageOffset
    getCacheFilesystemsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getCacheFilesystemsSpectraS3Request
}
func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetCacheFilesystemsSpectraS3Request {
    getCacheFilesystemsSpectraS3Request.pageStartMarker = pageStartMarker
    getCacheFilesystemsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getCacheFilesystemsSpectraS3Request
}


func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) WithLastPage() *GetCacheFilesystemsSpectraS3Request {
    getCacheFilesystemsSpectraS3Request.queryParams.Set("last_page", "")
    return getCacheFilesystemsSpectraS3Request
}

func (GetCacheFilesystemsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) Path() string {
    return "/_rest_/cache_filesystem"
}

func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) QueryParams() *url.Values {
    return getCacheFilesystemsSpectraS3Request.queryParams
}

func (GetCacheFilesystemsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetCacheFilesystemsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetCacheFilesystemsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
