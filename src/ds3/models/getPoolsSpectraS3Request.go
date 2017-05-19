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

type GetPoolsSpectraS3Request struct {
    assignedToStorageDomain bool
    bucketId string
    health PoolHealth
    lastVerified string
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    partitionId string
    poolType PoolType
    poweredOn bool
    state PoolState
    storageDomainId string
    queryParams *url.Values
}

func NewGetPoolsSpectraS3Request() *GetPoolsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetPoolsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithAssignedToStorageDomain(assignedToStorageDomain bool) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.assignedToStorageDomain = assignedToStorageDomain
    getPoolsSpectraS3Request.queryParams.Set("assigned_to_storage_domain", strconv.FormatBool(assignedToStorageDomain))
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithBucketId(bucketId string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.bucketId = bucketId
    getPoolsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithHealth(health PoolHealth) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.health = health
    getPoolsSpectraS3Request.queryParams.Set("health", health.String())
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithLastVerified(lastVerified string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.lastVerified = lastVerified
    getPoolsSpectraS3Request.queryParams.Set("last_verified", lastVerified)
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPageLength(pageLength int) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.pageLength = pageLength
    getPoolsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.pageOffset = pageOffset
    getPoolsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.pageStartMarker = pageStartMarker
    getPoolsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPartitionId(partitionId string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.partitionId = partitionId
    getPoolsSpectraS3Request.queryParams.Set("partition_id", partitionId)
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPoweredOn(poweredOn bool) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.poweredOn = poweredOn
    getPoolsSpectraS3Request.queryParams.Set("powered_on", strconv.FormatBool(poweredOn))
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithState(state PoolState) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.state = state
    getPoolsSpectraS3Request.queryParams.Set("state", state.String())
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.storageDomainId = storageDomainId
    getPoolsSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return getPoolsSpectraS3Request
}
func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPoolType(poolType PoolType) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.poolType = poolType
    getPoolsSpectraS3Request.queryParams.Set("type", poolType.String())
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithName(name *string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.name = name
    if name != nil {
        getPoolsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getPoolsSpectraS3Request.queryParams.Set("name", "")
    }
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithLastPage() *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.queryParams.Set("last_page", "")
    return getPoolsSpectraS3Request
}

func (GetPoolsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) Path() string {
    return "/_rest_/pool"
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) QueryParams() *url.Values {
    return getPoolsSpectraS3Request.queryParams
}

func (GetPoolsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetPoolsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetPoolsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
