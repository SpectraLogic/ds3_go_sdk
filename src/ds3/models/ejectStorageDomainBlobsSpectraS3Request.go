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

type EjectStorageDomainBlobsSpectraS3Request struct {
    bucketId string
    content networking.ReaderWithSizeDecorator
    ejectLabel *string
    ejectLocation *string
    storageDomainId string
    queryParams *url.Values
}

//TODO update autogen and add unit test
func NewEjectStorageDomainBlobsSpectraS3Request(bucketId string, objects []string, storageDomainId string) *EjectStorageDomainBlobsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "eject")
    queryParams.Set("blobs", "")
    queryParams.Set("bucket_id", bucketId)
    queryParams.Set("storage_domain_id", storageDomainId)

    return &EjectStorageDomainBlobsSpectraS3Request{
        bucketId: bucketId,
        storageDomainId: storageDomainId,
        content: buildDs3ObjectStreamFromNames(objects),
        queryParams: queryParams,
    }
}


func (ejectStorageDomainBlobsSpectraS3Request *EjectStorageDomainBlobsSpectraS3Request) WithEjectLabel(ejectLabel *string) *EjectStorageDomainBlobsSpectraS3Request {
    ejectStorageDomainBlobsSpectraS3Request.ejectLabel = ejectLabel
    if ejectLabel != nil {
        ejectStorageDomainBlobsSpectraS3Request.queryParams.Set("eject_label", *ejectLabel)
    } else {
        ejectStorageDomainBlobsSpectraS3Request.queryParams.Set("eject_label", "")
    }
    return ejectStorageDomainBlobsSpectraS3Request
}
func (ejectStorageDomainBlobsSpectraS3Request *EjectStorageDomainBlobsSpectraS3Request) WithEjectLocation(ejectLocation *string) *EjectStorageDomainBlobsSpectraS3Request {
    ejectStorageDomainBlobsSpectraS3Request.ejectLocation = ejectLocation
    if ejectLocation != nil {
        ejectStorageDomainBlobsSpectraS3Request.queryParams.Set("eject_location", *ejectLocation)
    } else {
        ejectStorageDomainBlobsSpectraS3Request.queryParams.Set("eject_location", "")
    }
    return ejectStorageDomainBlobsSpectraS3Request
}


func (EjectStorageDomainBlobsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (ejectStorageDomainBlobsSpectraS3Request *EjectStorageDomainBlobsSpectraS3Request) Path() string {
    return "/_rest_/tape"
}

func (ejectStorageDomainBlobsSpectraS3Request *EjectStorageDomainBlobsSpectraS3Request) QueryParams() *url.Values {
    return ejectStorageDomainBlobsSpectraS3Request.queryParams
}

func (EjectStorageDomainBlobsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (EjectStorageDomainBlobsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ejectStorageDomainBlobsSpectraS3Request *EjectStorageDomainBlobsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return ejectStorageDomainBlobsSpectraS3Request.content
}
