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

type GetBucketCapacitySummarySpectraS3Request struct {
    BucketId string
    PoolHealth PoolHealth
    PoolState PoolState
    PoolType PoolType
    StorageDomainId string
    TapeState TapeState
    TapeType *string
}

func NewGetBucketCapacitySummarySpectraS3Request(bucketId string, storageDomainId string) *GetBucketCapacitySummarySpectraS3Request {
    return &GetBucketCapacitySummarySpectraS3Request{
        BucketId: bucketId,
        StorageDomainId: storageDomainId,
    }
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithPoolHealth(poolHealth PoolHealth) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.PoolHealth = poolHealth
    return getBucketCapacitySummarySpectraS3Request
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithPoolState(poolState PoolState) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.PoolState = poolState
    return getBucketCapacitySummarySpectraS3Request
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithPoolType(poolType PoolType) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.PoolType = poolType
    return getBucketCapacitySummarySpectraS3Request
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithTapeState(tapeState TapeState) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.TapeState = tapeState
    return getBucketCapacitySummarySpectraS3Request
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithTapeType(tapeType string) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.TapeType = &tapeType
    return getBucketCapacitySummarySpectraS3Request
}

