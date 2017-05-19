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

type GetStorageDomainSpectraS3Request struct {
    storageDomain string
    queryParams *url.Values
}

func NewGetStorageDomainSpectraS3Request(storageDomain string) *GetStorageDomainSpectraS3Request {
    queryParams := &url.Values{}

    return &GetStorageDomainSpectraS3Request{
        storageDomain: storageDomain,
        queryParams: queryParams,
    }
}




func (GetStorageDomainSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getStorageDomainSpectraS3Request *GetStorageDomainSpectraS3Request) Path() string {
    return "/_rest_/storage_domain/" + getStorageDomainSpectraS3Request.storageDomain
}

func (getStorageDomainSpectraS3Request *GetStorageDomainSpectraS3Request) QueryParams() *url.Values {
    return getStorageDomainSpectraS3Request.queryParams
}

func (GetStorageDomainSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetStorageDomainSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetStorageDomainSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
