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

type DeleteStorageDomainFailureSpectraS3Request struct {
    storageDomainFailure string
    queryParams *url.Values
}

func NewDeleteStorageDomainFailureSpectraS3Request(storageDomainFailure string) *DeleteStorageDomainFailureSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteStorageDomainFailureSpectraS3Request{
        storageDomainFailure: storageDomainFailure,
        queryParams: queryParams,
    }
}




func (DeleteStorageDomainFailureSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteStorageDomainFailureSpectraS3Request *DeleteStorageDomainFailureSpectraS3Request) Path() string {
    return "/_rest_/storage_domain_failure/" + deleteStorageDomainFailureSpectraS3Request.storageDomainFailure
}

func (deleteStorageDomainFailureSpectraS3Request *DeleteStorageDomainFailureSpectraS3Request) QueryParams() *url.Values {
    return deleteStorageDomainFailureSpectraS3Request.queryParams
}

func (DeleteStorageDomainFailureSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteStorageDomainFailureSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteStorageDomainFailureSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
