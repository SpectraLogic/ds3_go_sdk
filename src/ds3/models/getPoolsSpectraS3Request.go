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

type GetPoolsSpectraS3Request struct {
    AssignedToStorageDomain *bool
    BucketId *string
    Guid *string
    Health PoolHealth
    LastPage bool
    LastVerified *string
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PartitionId *string
    PoolType PoolType
    PoweredOn *bool
    State PoolState
    StorageDomainId *string
}

func NewGetPoolsSpectraS3Request() *GetPoolsSpectraS3Request {
    return &GetPoolsSpectraS3Request{
    }
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithAssignedToStorageDomain(assignedToStorageDomain bool) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.AssignedToStorageDomain = &assignedToStorageDomain
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithBucketId(bucketId string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.BucketId = &bucketId
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithGuid(guid string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.Guid = &guid
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithHealth(health PoolHealth) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.Health = health
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithLastPage() *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.LastPage = true
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithLastVerified(lastVerified string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.LastVerified = &lastVerified
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithName(name string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.Name = &name
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPageLength(pageLength int) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PageLength = &pageLength
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PageOffset = &pageOffset
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPartitionId(partitionId string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PartitionId = &partitionId
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPoweredOn(poweredOn bool) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PoweredOn = &poweredOn
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithState(state PoolState) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.State = state
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.StorageDomainId = &storageDomainId
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPoolType(poolType PoolType) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PoolType = poolType
    return getPoolsSpectraS3Request
}

