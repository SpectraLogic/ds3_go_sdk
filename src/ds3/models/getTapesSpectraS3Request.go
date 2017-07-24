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

type GetTapesSpectraS3Request struct {
    assignedToStorageDomain bool
    availableRawCapacity *int64
    barCode *string
    bucketId string
    ejectLabel *string
    ejectLocation *string
    fullOfData bool
    lastVerified string
    pageLength int
    pageOffset int
    pageStartMarker string
    partiallyVerifiedEndOfTape string
    partitionId string
    previousState TapeState
    serialNumber *string
    sortBy *string
    state TapeState
    storageDomainId string
    tapeType TapeType
    verifyPending Priority
    writeProtected bool
    queryParams *url.Values
}

func NewGetTapesSpectraS3Request() *GetTapesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithAssignedToStorageDomain(assignedToStorageDomain bool) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.assignedToStorageDomain = assignedToStorageDomain
    getTapesSpectraS3Request.queryParams.Set("assigned_to_storage_domain", strconv.FormatBool(assignedToStorageDomain))
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithBucketId(bucketId string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.bucketId = bucketId
    getTapesSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithFullOfData(fullOfData bool) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.fullOfData = fullOfData
    getTapesSpectraS3Request.queryParams.Set("full_of_data", strconv.FormatBool(fullOfData))
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithLastVerified(lastVerified string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.lastVerified = lastVerified
    getTapesSpectraS3Request.queryParams.Set("last_verified", lastVerified)
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPageLength(pageLength int) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.pageLength = pageLength
    getTapesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.pageOffset = pageOffset
    getTapesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.pageStartMarker = pageStartMarker
    getTapesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPartiallyVerifiedEndOfTape(partiallyVerifiedEndOfTape string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.partiallyVerifiedEndOfTape = partiallyVerifiedEndOfTape
    getTapesSpectraS3Request.queryParams.Set("partially_verified_end_of_tape", partiallyVerifiedEndOfTape)
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPartitionId(partitionId string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.partitionId = partitionId
    getTapesSpectraS3Request.queryParams.Set("partition_id", partitionId)
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPreviousState(previousState TapeState) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.previousState = previousState
    getTapesSpectraS3Request.queryParams.Set("previous_state", previousState.String())
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithState(state TapeState) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.state = state
    getTapesSpectraS3Request.queryParams.Set("state", state.String())
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.storageDomainId = storageDomainId
    getTapesSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithTapeType(tapeType TapeType) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.tapeType = tapeType
    getTapesSpectraS3Request.queryParams.Set("type", tapeType.String())
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithVerifyPending(verifyPending Priority) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.verifyPending = verifyPending
    getTapesSpectraS3Request.queryParams.Set("verify_pending", verifyPending.String())
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithWriteProtected(writeProtected bool) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.writeProtected = writeProtected
    getTapesSpectraS3Request.queryParams.Set("write_protected", strconv.FormatBool(writeProtected))
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithAvailableRawCapacity(availableRawCapacity *int64) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.availableRawCapacity = availableRawCapacity
    if availableRawCapacity != nil {
        getTapesSpectraS3Request.queryParams.Set("available_raw_capacity", strconv.FormatInt(*availableRawCapacity, 10))
    } else {
        getTapesSpectraS3Request.queryParams.Set("available_raw_capacity", "")
    }
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithBarCode(barCode *string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.barCode = barCode
    if barCode != nil {
        getTapesSpectraS3Request.queryParams.Set("bar_code", *barCode)
    } else {
        getTapesSpectraS3Request.queryParams.Set("bar_code", "")
    }
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithEjectLabel(ejectLabel *string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.ejectLabel = ejectLabel
    if ejectLabel != nil {
        getTapesSpectraS3Request.queryParams.Set("eject_label", *ejectLabel)
    } else {
        getTapesSpectraS3Request.queryParams.Set("eject_label", "")
    }
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithEjectLocation(ejectLocation *string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.ejectLocation = ejectLocation
    if ejectLocation != nil {
        getTapesSpectraS3Request.queryParams.Set("eject_location", *ejectLocation)
    } else {
        getTapesSpectraS3Request.queryParams.Set("eject_location", "")
    }
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithSerialNumber(serialNumber *string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.serialNumber = serialNumber
    if serialNumber != nil {
        getTapesSpectraS3Request.queryParams.Set("serial_number", *serialNumber)
    } else {
        getTapesSpectraS3Request.queryParams.Set("serial_number", "")
    }
    return getTapesSpectraS3Request
}
func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithSortBy(sortBy *string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.sortBy = sortBy
    if sortBy != nil {
        getTapesSpectraS3Request.queryParams.Set("sort_by", *sortBy)
    } else {
        getTapesSpectraS3Request.queryParams.Set("sort_by", "")
    }
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithLastPage() *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.queryParams.Set("last_page", "")
    return getTapesSpectraS3Request
}

func (GetTapesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) Path() string {
    return "/_rest_/tape"
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) QueryParams() *url.Values {
    return getTapesSpectraS3Request.queryParams
}

func (GetTapesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
