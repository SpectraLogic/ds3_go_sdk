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
)

type GetSystemCapacitySummarySpectraS3Request struct {
    poolHealth PoolHealth
    poolState PoolState
    poolType PoolType
    tapeState TapeState
    tapeType TapeType
    queryParams *url.Values
}

func NewGetSystemCapacitySummarySpectraS3Request() *GetSystemCapacitySummarySpectraS3Request {
    queryParams := &url.Values{}

    return &GetSystemCapacitySummarySpectraS3Request{
        queryParams: queryParams,
    }
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithPoolHealth(poolHealth PoolHealth) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.poolHealth = poolHealth
    getSystemCapacitySummarySpectraS3Request.queryParams.Set("pool_health", poolHealth.String())
    return getSystemCapacitySummarySpectraS3Request
}
func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithPoolState(poolState PoolState) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.poolState = poolState
    getSystemCapacitySummarySpectraS3Request.queryParams.Set("pool_state", poolState.String())
    return getSystemCapacitySummarySpectraS3Request
}
func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithPoolType(poolType PoolType) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.poolType = poolType
    getSystemCapacitySummarySpectraS3Request.queryParams.Set("pool_type", poolType.String())
    return getSystemCapacitySummarySpectraS3Request
}
func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithTapeState(tapeState TapeState) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.tapeState = tapeState
    getSystemCapacitySummarySpectraS3Request.queryParams.Set("tape_state", tapeState.String())
    return getSystemCapacitySummarySpectraS3Request
}
func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithTapeType(tapeType TapeType) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.tapeType = tapeType
    getSystemCapacitySummarySpectraS3Request.queryParams.Set("tape_type", tapeType.String())
    return getSystemCapacitySummarySpectraS3Request
}



func (GetSystemCapacitySummarySpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) Path() string {
    return "/_rest_/capacity_summary"
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) QueryParams() *url.Values {
    return getSystemCapacitySummarySpectraS3Request.queryParams
}

func (GetSystemCapacitySummarySpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetSystemCapacitySummarySpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetSystemCapacitySummarySpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
