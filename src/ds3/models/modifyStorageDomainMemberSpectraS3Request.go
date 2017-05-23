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

type ModifyStorageDomainMemberSpectraS3Request struct {
    storageDomainMember string
    writePreference WritePreferenceLevel
    queryParams *url.Values
}

func NewModifyStorageDomainMemberSpectraS3Request(storageDomainMember string) *ModifyStorageDomainMemberSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyStorageDomainMemberSpectraS3Request{
        storageDomainMember: storageDomainMember,
        queryParams: queryParams,
    }
}

func (modifyStorageDomainMemberSpectraS3Request *ModifyStorageDomainMemberSpectraS3Request) WithWritePreference(writePreference WritePreferenceLevel) *ModifyStorageDomainMemberSpectraS3Request {
    modifyStorageDomainMemberSpectraS3Request.writePreference = writePreference
    modifyStorageDomainMemberSpectraS3Request.queryParams.Set("write_preference", writePreference.String())
    return modifyStorageDomainMemberSpectraS3Request
}



func (ModifyStorageDomainMemberSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyStorageDomainMemberSpectraS3Request *ModifyStorageDomainMemberSpectraS3Request) Path() string {
    return "/_rest_/storage_domain_member/" + modifyStorageDomainMemberSpectraS3Request.storageDomainMember
}

func (modifyStorageDomainMemberSpectraS3Request *ModifyStorageDomainMemberSpectraS3Request) QueryParams() *url.Values {
    return modifyStorageDomainMemberSpectraS3Request.queryParams
}

func (ModifyStorageDomainMemberSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyStorageDomainMemberSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyStorageDomainMemberSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
