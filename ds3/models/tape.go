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

type Tape struct {
    AssignedToStorageDomain bool
    AvailableRawCapacity *int64
    BarCode *string
    BucketId *string
    DescriptionForIdentification *string
    EjectDate *string
    EjectLabel *string
    EjectLocation *string
    EjectPending *string
    FullOfData bool
    Id string
    LastAccessed *string
    LastCheckpoint *string
    LastModified *string
    LastVerified *string
    PartiallyVerifiedEndOfTape *string
    PartitionId *string
    PreviousState *TapeState
    SerialNumber *string
    State TapeState
    StorageDomainId *string
    TakeOwnershipPending bool
    TotalRawCapacity *int64
    Type string
    VerifyPending *Priority
    WriteProtected bool
}