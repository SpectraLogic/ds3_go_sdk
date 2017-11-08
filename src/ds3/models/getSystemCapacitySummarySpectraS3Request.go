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

type GetSystemCapacitySummarySpectraS3Request struct {
    PoolHealth PoolHealth
    PoolState PoolState
    PoolType PoolType
    TapeState TapeState
    TapeType TapeType
}

func NewGetSystemCapacitySummarySpectraS3Request() *GetSystemCapacitySummarySpectraS3Request {
    return &GetSystemCapacitySummarySpectraS3Request{
    }
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithPoolHealth(poolHealth PoolHealth) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.PoolHealth = poolHealth
    return getSystemCapacitySummarySpectraS3Request
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithPoolState(poolState PoolState) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.PoolState = poolState
    return getSystemCapacitySummarySpectraS3Request
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithPoolType(poolType PoolType) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.PoolType = poolType
    return getSystemCapacitySummarySpectraS3Request
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithTapeState(tapeState TapeState) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.TapeState = tapeState
    return getSystemCapacitySummarySpectraS3Request
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithTapeType(tapeType TapeType) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.TapeType = tapeType
    return getSystemCapacitySummarySpectraS3Request
}

