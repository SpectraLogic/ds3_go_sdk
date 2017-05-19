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

type GetPoolFailuresSpectraS3Request struct {
    errorMessage *string
    pageLength int
    pageOffset int
    pageStartMarker string
    poolFailureType PoolFailureType
    poolId string
    queryParams *url.Values
}

func NewGetPoolFailuresSpectraS3Request() *GetPoolFailuresSpectraS3Request {
    queryParams := &url.Values{}

    return &GetPoolFailuresSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPageLength(pageLength int) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.pageLength = pageLength
    getPoolFailuresSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getPoolFailuresSpectraS3Request
}
func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.pageOffset = pageOffset
    getPoolFailuresSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getPoolFailuresSpectraS3Request
}
func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.pageStartMarker = pageStartMarker
    getPoolFailuresSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getPoolFailuresSpectraS3Request
}
func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPoolId(poolId string) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.poolId = poolId
    getPoolFailuresSpectraS3Request.queryParams.Set("pool_id", poolId)
    return getPoolFailuresSpectraS3Request
}
func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPoolFailureType(poolFailureType PoolFailureType) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.poolFailureType = poolFailureType
    getPoolFailuresSpectraS3Request.queryParams.Set("type", poolFailureType.String())
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithErrorMessage(errorMessage *string) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.errorMessage = errorMessage
    if errorMessage != nil {
        getPoolFailuresSpectraS3Request.queryParams.Set("error_message", *errorMessage)
    } else {
        getPoolFailuresSpectraS3Request.queryParams.Set("error_message", "")
    }
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithLastPage() *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.queryParams.Set("last_page", "")
    return getPoolFailuresSpectraS3Request
}

func (GetPoolFailuresSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) Path() string {
    return "/_rest_/pool_failure"
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) QueryParams() *url.Values {
    return getPoolFailuresSpectraS3Request.queryParams
}

func (GetPoolFailuresSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetPoolFailuresSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetPoolFailuresSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
