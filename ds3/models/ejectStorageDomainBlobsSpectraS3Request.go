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

type EjectStorageDomainBlobsSpectraS3Request struct {
    BucketId string
    EjectLabel *string
    EjectLocation *string
    Objects []Ds3GetObject
    StorageDomain string
}

func NewEjectStorageDomainBlobsSpectraS3Request(bucketId string, storageDomain string, objectNames []string) *EjectStorageDomainBlobsSpectraS3Request {

    return &EjectStorageDomainBlobsSpectraS3Request{
        BucketId: bucketId,
        StorageDomain: storageDomain,
        Objects: buildDs3GetObjectSliceFromNames(objectNames),
    }
}

func NewEjectStorageDomainBlobsSpectraS3RequestWithPartialObjects(bucketId string, storageDomain string, objects []Ds3GetObject) *EjectStorageDomainBlobsSpectraS3Request {

    return &EjectStorageDomainBlobsSpectraS3Request{
        BucketId: bucketId,
        StorageDomain: storageDomain,
        Objects: objects,
    }
}

func (ejectStorageDomainBlobsSpectraS3Request *EjectStorageDomainBlobsSpectraS3Request) WithEjectLabel(ejectLabel string) *EjectStorageDomainBlobsSpectraS3Request {
    ejectStorageDomainBlobsSpectraS3Request.EjectLabel = &ejectLabel
    return ejectStorageDomainBlobsSpectraS3Request
}

func (ejectStorageDomainBlobsSpectraS3Request *EjectStorageDomainBlobsSpectraS3Request) WithEjectLocation(ejectLocation string) *EjectStorageDomainBlobsSpectraS3Request {
    ejectStorageDomainBlobsSpectraS3Request.EjectLocation = &ejectLocation
    return ejectStorageDomainBlobsSpectraS3Request
}

