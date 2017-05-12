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

type EjectStorageDomainSpectraS3Request struct {
    bucketId string
    ejectLabel *string
    ejectLocation *string
    storageDomainId string
    queryParams *url.Values
}

func NewEjectStorageDomainSpectraS3Request(storageDomainId string) *EjectStorageDomainSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "eject")
    queryParams.Set("storage_domain_id", storageDomainId)

    return &EjectStorageDomainSpectraS3Request{
        storageDomainId: storageDomainId,
        queryParams: queryParams,
    }
}

func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) WithBucketId(bucketId string) *EjectStorageDomainSpectraS3Request {
    ejectStorageDomainSpectraS3Request.bucketId = bucketId
    ejectStorageDomainSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return ejectStorageDomainSpectraS3Request
}

func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) WithEjectLabel(ejectLabel *string) *EjectStorageDomainSpectraS3Request {
    ejectStorageDomainSpectraS3Request.ejectLabel = ejectLabel
    if ejectLabel != nil {
        ejectStorageDomainSpectraS3Request.queryParams.Set("eject_label", *ejectLabel)
    } else {
        ejectStorageDomainSpectraS3Request.queryParams.Set("eject_label", "")
    }
    return ejectStorageDomainSpectraS3Request
}
func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) WithEjectLocation(ejectLocation *string) *EjectStorageDomainSpectraS3Request {
    ejectStorageDomainSpectraS3Request.ejectLocation = ejectLocation
    if ejectLocation != nil {
        ejectStorageDomainSpectraS3Request.queryParams.Set("eject_location", *ejectLocation)
    } else {
        ejectStorageDomainSpectraS3Request.queryParams.Set("eject_location", "")
    }
    return ejectStorageDomainSpectraS3Request
}


func (EjectStorageDomainSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) Path() string {
    return "/_rest_/tape"
}

func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) QueryParams() *url.Values {
    return ejectStorageDomainSpectraS3Request.queryParams
}

func (EjectStorageDomainSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (EjectStorageDomainSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (EjectStorageDomainSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
