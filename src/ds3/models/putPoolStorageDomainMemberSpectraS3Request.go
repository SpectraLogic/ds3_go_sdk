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

type PutPoolStorageDomainMemberSpectraS3Request struct {
    poolPartitionId string
    storageDomainId string
    writePreference WritePreferenceLevel
    queryParams *url.Values
}

func NewPutPoolStorageDomainMemberSpectraS3Request(poolPartitionId string, storageDomainId string) *PutPoolStorageDomainMemberSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("pool_partition_id", poolPartitionId)
    queryParams.Set("storage_domain_id", storageDomainId)

    return &PutPoolStorageDomainMemberSpectraS3Request{
        poolPartitionId: poolPartitionId,
        storageDomainId: storageDomainId,
        queryParams: queryParams,
    }
}

func (putPoolStorageDomainMemberSpectraS3Request *PutPoolStorageDomainMemberSpectraS3Request) WithWritePreference(writePreference WritePreferenceLevel) *PutPoolStorageDomainMemberSpectraS3Request {
    putPoolStorageDomainMemberSpectraS3Request.writePreference = writePreference
    putPoolStorageDomainMemberSpectraS3Request.queryParams.Set("write_preference", writePreference.String())
    return putPoolStorageDomainMemberSpectraS3Request
}



func (PutPoolStorageDomainMemberSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putPoolStorageDomainMemberSpectraS3Request *PutPoolStorageDomainMemberSpectraS3Request) Path() string {
    return "/_rest_/storage_domain_member"
}

func (putPoolStorageDomainMemberSpectraS3Request *PutPoolStorageDomainMemberSpectraS3Request) QueryParams() *url.Values {
    return putPoolStorageDomainMemberSpectraS3Request.queryParams
}

func (PutPoolStorageDomainMemberSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutPoolStorageDomainMemberSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutPoolStorageDomainMemberSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
