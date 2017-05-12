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

type ConvertStorageDomainToDs3TargetSpectraS3Request struct {
    convertToDs3Target string
    storageDomain string
    queryParams *url.Values
}

func NewConvertStorageDomainToDs3TargetSpectraS3Request(convertToDs3Target string, storageDomain string) *ConvertStorageDomainToDs3TargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("convert_to_ds3_target", convertToDs3Target)

    return &ConvertStorageDomainToDs3TargetSpectraS3Request{
        storageDomain: storageDomain,
        convertToDs3Target: convertToDs3Target,
        queryParams: queryParams,
    }
}




func (ConvertStorageDomainToDs3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (convertStorageDomainToDs3TargetSpectraS3Request *ConvertStorageDomainToDs3TargetSpectraS3Request) Path() string {
    return "/_rest_/storage_domain/" + convertStorageDomainToDs3TargetSpectraS3Request.storageDomain
}

func (convertStorageDomainToDs3TargetSpectraS3Request *ConvertStorageDomainToDs3TargetSpectraS3Request) QueryParams() *url.Values {
    return convertStorageDomainToDs3TargetSpectraS3Request.queryParams
}

func (ConvertStorageDomainToDs3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ConvertStorageDomainToDs3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ConvertStorageDomainToDs3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
