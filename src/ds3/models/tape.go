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
    AssignedToStorageDomain bool `xml:"AssignedToStorageDomain"`
    AvailableRawCapacity *int64 `xml:"AvailableRawCapacity"`
    BarCode *string `xml:"BarCode"`
    BucketId *string `xml:"BucketId"`
    DescriptionForIdentification *string `xml:"DescriptionForIdentification"`
    EjectDate *string `xml:"EjectDate"`
    EjectLabel *string `xml:"EjectLabel"`
    EjectLocation *string `xml:"EjectLocation"`
    EjectPending *string `xml:"EjectPending"`
    FullOfData bool `xml:"FullOfData"`
    Id string `xml:"Id"`
    LastAccessed *string `xml:"LastAccessed"`
    LastCheckpoint *string `xml:"LastCheckpoint"`
    LastModified *string `xml:"LastModified"`
    LastVerified *string `xml:"LastVerified"`
    PartiallyVerifiedEndOfTape *string `xml:"PartiallyVerifiedEndOfTape"`
    PartitionId *string `xml:"PartitionId"`
    PreviousState *TapeState `xml:"PreviousState"`
    SerialNumber *string `xml:"SerialNumber"`
    State TapeState `xml:"State"`
    StorageDomainId *string `xml:"StorageDomainId"`
    TakeOwnershipPending bool `xml:"TakeOwnershipPending"`
    TotalRawCapacity *int64 `xml:"TotalRawCapacity"`
    Type TapeType `xml:"Type"`
    VerifyPending *Priority `xml:"VerifyPending"`
    WriteProtected bool `xml:"WriteProtected"`
}