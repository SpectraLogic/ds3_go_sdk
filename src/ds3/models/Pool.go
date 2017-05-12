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

type Pool struct {
    AssignedToStorageDomain bool `xml:"AssignedToStorageDomain"`
    AvailableCapacity int64 `xml:"AvailableCapacity"`
    BucketId *string `xml:"BucketId"`
    Guid *string `xml:"Guid"`
    Health PoolHealth `xml:"Health"`
    Id string `xml:"Id"`
    LastAccessed *string `xml:"LastAccessed"`
    LastModified *string `xml:"LastModified"`
    LastVerified *string `xml:"LastVerified"`
    Mountpoint *string `xml:"Mountpoint"`
    Name *string `xml:"Name"`
    PartitionId *string `xml:"PartitionId"`
    PoweredOn bool `xml:"PoweredOn"`
    Quiesced Quiesced `xml:"Quiesced"`
    ReservedCapacity int64 `xml:"ReservedCapacity"`
    State PoolState `xml:"State"`
    StorageDomainId *string `xml:"StorageDomainId"`
    TotalCapacity int64 `xml:"TotalCapacity"`
    Type PoolType `xml:"Type"`
    UsedCapacity int64 `xml:"UsedCapacity"`
}