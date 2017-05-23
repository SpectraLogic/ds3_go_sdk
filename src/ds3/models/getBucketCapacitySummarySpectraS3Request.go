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

type GetBucketCapacitySummarySpectraS3Request struct {
    bucketId string
    poolHealth PoolHealth
    poolState PoolState
    poolType PoolType
    storageDomainId string
    tapeState TapeState
    tapeType TapeType
    queryParams *url.Values
}

func NewGetBucketCapacitySummarySpectraS3Request(bucketId string, storageDomainId string) *GetBucketCapacitySummarySpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("bucket_id", bucketId)
    queryParams.Set("storage_domain_id", storageDomainId)

    return &GetBucketCapacitySummarySpectraS3Request{
        bucketId: bucketId,
        storageDomainId: storageDomainId,
        queryParams: queryParams,
    }
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithPoolHealth(poolHealth PoolHealth) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.poolHealth = poolHealth
    getBucketCapacitySummarySpectraS3Request.queryParams.Set("pool_health", poolHealth.String())
    return getBucketCapacitySummarySpectraS3Request
}
func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithPoolState(poolState PoolState) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.poolState = poolState
    getBucketCapacitySummarySpectraS3Request.queryParams.Set("pool_state", poolState.String())
    return getBucketCapacitySummarySpectraS3Request
}
func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithPoolType(poolType PoolType) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.poolType = poolType
    getBucketCapacitySummarySpectraS3Request.queryParams.Set("pool_type", poolType.String())
    return getBucketCapacitySummarySpectraS3Request
}
func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithTapeState(tapeState TapeState) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.tapeState = tapeState
    getBucketCapacitySummarySpectraS3Request.queryParams.Set("tape_state", tapeState.String())
    return getBucketCapacitySummarySpectraS3Request
}
func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithTapeType(tapeType TapeType) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.tapeType = tapeType
    getBucketCapacitySummarySpectraS3Request.queryParams.Set("tape_type", tapeType.String())
    return getBucketCapacitySummarySpectraS3Request
}



func (GetBucketCapacitySummarySpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) Path() string {
    return "/_rest_/capacity_summary"
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) QueryParams() *url.Values {
    return getBucketCapacitySummarySpectraS3Request.queryParams
}

func (GetBucketCapacitySummarySpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetBucketCapacitySummarySpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetBucketCapacitySummarySpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
