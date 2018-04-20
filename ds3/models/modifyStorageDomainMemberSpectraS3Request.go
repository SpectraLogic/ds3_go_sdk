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

type ModifyStorageDomainMemberSpectraS3Request struct {
    AutoCompactionThreshold *int
    State StorageDomainMemberState
    StorageDomainMember string
    WritePreference WritePreferenceLevel
}

func NewModifyStorageDomainMemberSpectraS3Request(storageDomainMember string) *ModifyStorageDomainMemberSpectraS3Request {
    return &ModifyStorageDomainMemberSpectraS3Request{
        StorageDomainMember: storageDomainMember,
    }
}

func (modifyStorageDomainMemberSpectraS3Request *ModifyStorageDomainMemberSpectraS3Request) WithAutoCompactionThreshold(autoCompactionThreshold int) *ModifyStorageDomainMemberSpectraS3Request {
    modifyStorageDomainMemberSpectraS3Request.AutoCompactionThreshold = &autoCompactionThreshold
    return modifyStorageDomainMemberSpectraS3Request
}

func (modifyStorageDomainMemberSpectraS3Request *ModifyStorageDomainMemberSpectraS3Request) WithState(state StorageDomainMemberState) *ModifyStorageDomainMemberSpectraS3Request {
    modifyStorageDomainMemberSpectraS3Request.State = state
    return modifyStorageDomainMemberSpectraS3Request
}

func (modifyStorageDomainMemberSpectraS3Request *ModifyStorageDomainMemberSpectraS3Request) WithWritePreference(writePreference WritePreferenceLevel) *ModifyStorageDomainMemberSpectraS3Request {
    modifyStorageDomainMemberSpectraS3Request.WritePreference = writePreference
    return modifyStorageDomainMemberSpectraS3Request
}

