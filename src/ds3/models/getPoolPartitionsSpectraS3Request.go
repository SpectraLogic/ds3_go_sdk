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

type GetPoolPartitionsSpectraS3Request struct {
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    poolType PoolType
    queryParams *url.Values
}

func NewGetPoolPartitionsSpectraS3Request() *GetPoolPartitionsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetPoolPartitionsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPageLength(pageLength int) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.pageLength = pageLength
    getPoolPartitionsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getPoolPartitionsSpectraS3Request
}
func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.pageOffset = pageOffset
    getPoolPartitionsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getPoolPartitionsSpectraS3Request
}
func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.pageStartMarker = pageStartMarker
    getPoolPartitionsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getPoolPartitionsSpectraS3Request
}
func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPoolType(poolType PoolType) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.poolType = poolType
    getPoolPartitionsSpectraS3Request.queryParams.Set("type", poolType.String())
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithName(name *string) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.name = name
    if name != nil {
        getPoolPartitionsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getPoolPartitionsSpectraS3Request.queryParams.Set("name", "")
    }
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithLastPage() *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.queryParams.Set("last_page", "")
    return getPoolPartitionsSpectraS3Request
}

func (GetPoolPartitionsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) Path() string {
    return "/_rest_/pool_partition"
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) QueryParams() *url.Values {
    return getPoolPartitionsSpectraS3Request.queryParams
}

func (GetPoolPartitionsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetPoolPartitionsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetPoolPartitionsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
