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

type GetTapeDrivesSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PartitionId *string
    ReservedTaskType ReservedTaskType
    SerialNumber *string
    State TapeDriveState
    TapeDriveType TapeDriveType
}

func NewGetTapeDrivesSpectraS3Request() *GetTapeDrivesSpectraS3Request {
    return &GetTapeDrivesSpectraS3Request{
    }
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithLastPage() *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.LastPage = true
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPageLength(pageLength int) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.PageLength = &pageLength
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.PageOffset = &pageOffset
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPartitionId(partitionId string) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.PartitionId = &partitionId
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithReservedTaskType(reservedTaskType ReservedTaskType) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.ReservedTaskType = reservedTaskType
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithSerialNumber(serialNumber string) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.SerialNumber = &serialNumber
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithState(state TapeDriveState) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.State = state
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithTapeDriveType(tapeDriveType TapeDriveType) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.TapeDriveType = tapeDriveType
    return getTapeDrivesSpectraS3Request
}

