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
)

type PutTapeStorageDomainMemberSpectraS3Request struct {
    storageDomainId string
    tapePartitionId string
    tapeType TapeType
    writePreference WritePreferenceLevel
    queryParams *url.Values
}

func NewPutTapeStorageDomainMemberSpectraS3Request(storageDomainId string, tapePartitionId string, tapeType TapeType) *PutTapeStorageDomainMemberSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("storage_domain_id", storageDomainId)
    queryParams.Set("tape_partition_id", tapePartitionId)
    queryParams.Set("tape_type", tapeType.String())

    return &PutTapeStorageDomainMemberSpectraS3Request{
        storageDomainId: storageDomainId,
        tapePartitionId: tapePartitionId,
        tapeType: tapeType,
        queryParams: queryParams,
    }
}

func (putTapeStorageDomainMemberSpectraS3Request *PutTapeStorageDomainMemberSpectraS3Request) WithWritePreference(writePreference WritePreferenceLevel) *PutTapeStorageDomainMemberSpectraS3Request {
    putTapeStorageDomainMemberSpectraS3Request.writePreference = writePreference
    putTapeStorageDomainMemberSpectraS3Request.queryParams.Set("write_preference", writePreference.String())
    return putTapeStorageDomainMemberSpectraS3Request
}



func (PutTapeStorageDomainMemberSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putTapeStorageDomainMemberSpectraS3Request *PutTapeStorageDomainMemberSpectraS3Request) Path() string {
    return "/_rest_/storage_domain_member"
}

func (putTapeStorageDomainMemberSpectraS3Request *PutTapeStorageDomainMemberSpectraS3Request) QueryParams() *url.Values {
    return putTapeStorageDomainMemberSpectraS3Request.queryParams
}

func (PutTapeStorageDomainMemberSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutTapeStorageDomainMemberSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutTapeStorageDomainMemberSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
