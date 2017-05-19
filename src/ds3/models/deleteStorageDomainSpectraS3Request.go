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

type DeleteStorageDomainSpectraS3Request struct {
    storageDomain string
    queryParams *url.Values
}

func NewDeleteStorageDomainSpectraS3Request(storageDomain string) *DeleteStorageDomainSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteStorageDomainSpectraS3Request{
        storageDomain: storageDomain,
        queryParams: queryParams,
    }
}




func (DeleteStorageDomainSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteStorageDomainSpectraS3Request *DeleteStorageDomainSpectraS3Request) Path() string {
    return "/_rest_/storage_domain/" + deleteStorageDomainSpectraS3Request.storageDomain
}

func (deleteStorageDomainSpectraS3Request *DeleteStorageDomainSpectraS3Request) QueryParams() *url.Values {
    return deleteStorageDomainSpectraS3Request.queryParams
}

func (DeleteStorageDomainSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteStorageDomainSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteStorageDomainSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
