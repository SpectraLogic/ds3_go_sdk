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

type GetStorageDomainMembersSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PoolPartitionId *string
    State StorageDomainMemberState
    StorageDomainId *string
    TapePartitionId *string
    TapeType TapeType
    WritePreference WritePreferenceLevel
}

func NewGetStorageDomainMembersSpectraS3Request() *GetStorageDomainMembersSpectraS3Request {
    return &GetStorageDomainMembersSpectraS3Request{
    }
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithLastPage() *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.LastPage = true
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.PageLength = &pageLength
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.PageOffset = &pageOffset
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.PageStartMarker = &pageStartMarker
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPoolPartitionId(poolPartitionId string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.PoolPartitionId = &poolPartitionId
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithState(state StorageDomainMemberState) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.State = state
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.StorageDomainId = &storageDomainId
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithTapePartitionId(tapePartitionId string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.TapePartitionId = &tapePartitionId
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithTapeType(tapeType TapeType) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.TapeType = tapeType
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithWritePreference(writePreference WritePreferenceLevel) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.WritePreference = writePreference
    return getStorageDomainMembersSpectraS3Request
}

