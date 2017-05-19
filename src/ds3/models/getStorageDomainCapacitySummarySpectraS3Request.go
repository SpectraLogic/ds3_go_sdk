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

type GetStorageDomainCapacitySummarySpectraS3Request struct {
    poolHealth PoolHealth
    poolState PoolState
    poolType PoolType
    storageDomainId string
    tapeState TapeState
    tapeType TapeType
    queryParams *url.Values
}

func NewGetStorageDomainCapacitySummarySpectraS3Request(storageDomainId string) *GetStorageDomainCapacitySummarySpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("storage_domain_id", storageDomainId)

    return &GetStorageDomainCapacitySummarySpectraS3Request{
        storageDomainId: storageDomainId,
        queryParams: queryParams,
    }
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithPoolHealth(poolHealth PoolHealth) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.poolHealth = poolHealth
    getStorageDomainCapacitySummarySpectraS3Request.queryParams.Set("pool_health", poolHealth.String())
    return getStorageDomainCapacitySummarySpectraS3Request
}
func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithPoolState(poolState PoolState) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.poolState = poolState
    getStorageDomainCapacitySummarySpectraS3Request.queryParams.Set("pool_state", poolState.String())
    return getStorageDomainCapacitySummarySpectraS3Request
}
func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithPoolType(poolType PoolType) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.poolType = poolType
    getStorageDomainCapacitySummarySpectraS3Request.queryParams.Set("pool_type", poolType.String())
    return getStorageDomainCapacitySummarySpectraS3Request
}
func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithTapeState(tapeState TapeState) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.tapeState = tapeState
    getStorageDomainCapacitySummarySpectraS3Request.queryParams.Set("tape_state", tapeState.String())
    return getStorageDomainCapacitySummarySpectraS3Request
}
func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithTapeType(tapeType TapeType) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.tapeType = tapeType
    getStorageDomainCapacitySummarySpectraS3Request.queryParams.Set("tape_type", tapeType.String())
    return getStorageDomainCapacitySummarySpectraS3Request
}



func (GetStorageDomainCapacitySummarySpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) Path() string {
    return "/_rest_/capacity_summary"
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) QueryParams() *url.Values {
    return getStorageDomainCapacitySummarySpectraS3Request.queryParams
}

func (GetStorageDomainCapacitySummarySpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetStorageDomainCapacitySummarySpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetStorageDomainCapacitySummarySpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
