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

type GetStorageDomainCapacitySummarySpectraS3Request struct {
    PoolHealth PoolHealth
    PoolState PoolState
    PoolType PoolType
    StorageDomainId string
    TapeState TapeState
    TapeType TapeType
}

func NewGetStorageDomainCapacitySummarySpectraS3Request(storageDomainId string) *GetStorageDomainCapacitySummarySpectraS3Request {
    return &GetStorageDomainCapacitySummarySpectraS3Request{
        StorageDomainId: storageDomainId,
    }
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithPoolHealth(poolHealth PoolHealth) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.PoolHealth = poolHealth
    return getStorageDomainCapacitySummarySpectraS3Request
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithPoolState(poolState PoolState) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.PoolState = poolState
    return getStorageDomainCapacitySummarySpectraS3Request
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithPoolType(poolType PoolType) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.PoolType = poolType
    return getStorageDomainCapacitySummarySpectraS3Request
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithTapeState(tapeState TapeState) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.TapeState = tapeState
    return getStorageDomainCapacitySummarySpectraS3Request
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithTapeType(tapeType TapeType) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.TapeType = tapeType
    return getStorageDomainCapacitySummarySpectraS3Request
}

