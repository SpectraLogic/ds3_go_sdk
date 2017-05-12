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

type GetStorageDomainMembersSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    poolPartitionId string
    state StorageDomainMemberState
    storageDomainId string
    tapePartitionId string
    tapeType TapeType
    writePreference WritePreferenceLevel
    queryParams *url.Values
}

func NewGetStorageDomainMembersSpectraS3Request() *GetStorageDomainMembersSpectraS3Request {
    queryParams := &url.Values{}

    return &GetStorageDomainMembersSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.pageLength = pageLength
    getStorageDomainMembersSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getStorageDomainMembersSpectraS3Request
}
func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.pageOffset = pageOffset
    getStorageDomainMembersSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getStorageDomainMembersSpectraS3Request
}
func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.pageStartMarker = pageStartMarker
    getStorageDomainMembersSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getStorageDomainMembersSpectraS3Request
}
func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPoolPartitionId(poolPartitionId string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.poolPartitionId = poolPartitionId
    getStorageDomainMembersSpectraS3Request.queryParams.Set("pool_partition_id", poolPartitionId)
    return getStorageDomainMembersSpectraS3Request
}
func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithState(state StorageDomainMemberState) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.state = state
    getStorageDomainMembersSpectraS3Request.queryParams.Set("state", state.String())
    return getStorageDomainMembersSpectraS3Request
}
func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.storageDomainId = storageDomainId
    getStorageDomainMembersSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return getStorageDomainMembersSpectraS3Request
}
func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithTapePartitionId(tapePartitionId string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.tapePartitionId = tapePartitionId
    getStorageDomainMembersSpectraS3Request.queryParams.Set("tape_partition_id", tapePartitionId)
    return getStorageDomainMembersSpectraS3Request
}
func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithTapeType(tapeType TapeType) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.tapeType = tapeType
    getStorageDomainMembersSpectraS3Request.queryParams.Set("tape_type", tapeType.String())
    return getStorageDomainMembersSpectraS3Request
}
func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithWritePreference(writePreference WritePreferenceLevel) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.writePreference = writePreference
    getStorageDomainMembersSpectraS3Request.queryParams.Set("write_preference", writePreference.String())
    return getStorageDomainMembersSpectraS3Request
}


func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithLastPage() *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.queryParams.Set("last_page", "")
    return getStorageDomainMembersSpectraS3Request
}

func (GetStorageDomainMembersSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) Path() string {
    return "/_rest_/storage_domain_member"
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) QueryParams() *url.Values {
    return getStorageDomainMembersSpectraS3Request.queryParams
}

func (GetStorageDomainMembersSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetStorageDomainMembersSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetStorageDomainMembersSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
