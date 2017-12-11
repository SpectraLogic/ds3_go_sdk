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

type EjectStorageDomainSpectraS3Request struct {
    BucketId *string
    EjectLabel *string
    EjectLocation *string
    StorageDomainId string
}

func NewEjectStorageDomainSpectraS3Request(storageDomainId string) *EjectStorageDomainSpectraS3Request {
    return &EjectStorageDomainSpectraS3Request{
        StorageDomainId: storageDomainId,
    }
}

func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) WithBucketId(bucketId string) *EjectStorageDomainSpectraS3Request {
    ejectStorageDomainSpectraS3Request.BucketId = &bucketId
    return ejectStorageDomainSpectraS3Request
}

func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) WithEjectLabel(ejectLabel string) *EjectStorageDomainSpectraS3Request {
    ejectStorageDomainSpectraS3Request.EjectLabel = &ejectLabel
    return ejectStorageDomainSpectraS3Request
}

func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) WithEjectLocation(ejectLocation string) *EjectStorageDomainSpectraS3Request {
    ejectStorageDomainSpectraS3Request.EjectLocation = &ejectLocation
    return ejectStorageDomainSpectraS3Request
}

