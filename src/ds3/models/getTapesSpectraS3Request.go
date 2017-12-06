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

type GetTapesSpectraS3Request struct {
    AssignedToStorageDomain *bool
    AvailableRawCapacity *int64
    BarCode *string
    BucketId *string
    EjectLabel *string
    EjectLocation *string
    FullOfData *bool
    LastPage bool
    LastVerified *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PartiallyVerifiedEndOfTape *string
    PartitionId *string
    PreviousState TapeState
    SerialNumber *string
    SortBy *string
    State TapeState
    StorageDomainId *string
    TapeType TapeType
    VerifyPending Priority
    WriteProtected *bool
}

func NewGetTapesSpectraS3Request() *GetTapesSpectraS3Request {
    return &GetTapesSpectraS3Request{
    }
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithAssignedToStorageDomain(assignedToStorageDomain bool) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.AssignedToStorageDomain = &assignedToStorageDomain
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithAvailableRawCapacity(availableRawCapacity int64) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.AvailableRawCapacity = &availableRawCapacity
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithBarCode(barCode string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.BarCode = &barCode
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithBucketId(bucketId string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.BucketId = &bucketId
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithEjectLabel(ejectLabel string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.EjectLabel = &ejectLabel
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithEjectLocation(ejectLocation string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.EjectLocation = &ejectLocation
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithFullOfData(fullOfData bool) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.FullOfData = &fullOfData
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithLastPage() *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.LastPage = true
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithLastVerified(lastVerified string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.LastVerified = &lastVerified
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPageLength(pageLength int) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PageLength = &pageLength
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PageOffset = &pageOffset
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPartiallyVerifiedEndOfTape(partiallyVerifiedEndOfTape string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PartiallyVerifiedEndOfTape = &partiallyVerifiedEndOfTape
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPartitionId(partitionId string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PartitionId = &partitionId
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPreviousState(previousState TapeState) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PreviousState = previousState
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithSerialNumber(serialNumber string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.SerialNumber = &serialNumber
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithSortBy(sortBy string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.SortBy = &sortBy
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithState(state TapeState) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.State = state
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.StorageDomainId = &storageDomainId
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithTapeType(tapeType TapeType) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.TapeType = tapeType
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithVerifyPending(verifyPending Priority) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.VerifyPending = verifyPending
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithWriteProtected(writeProtected bool) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.WriteProtected = &writeProtected
    return getTapesSpectraS3Request
}

