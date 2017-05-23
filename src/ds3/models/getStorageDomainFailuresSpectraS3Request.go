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
    "strconv"
)

type GetStorageDomainFailuresSpectraS3Request struct {
    errorMessage *string
    pageLength int
    pageOffset int
    pageStartMarker string
    storageDomainFailureType StorageDomainFailureType
    storageDomainId string
    queryParams *url.Values
}

func NewGetStorageDomainFailuresSpectraS3Request() *GetStorageDomainFailuresSpectraS3Request {
    queryParams := &url.Values{}

    return &GetStorageDomainFailuresSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.pageLength = pageLength
    getStorageDomainFailuresSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getStorageDomainFailuresSpectraS3Request
}
func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.pageOffset = pageOffset
    getStorageDomainFailuresSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getStorageDomainFailuresSpectraS3Request
}
func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.pageStartMarker = pageStartMarker
    getStorageDomainFailuresSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getStorageDomainFailuresSpectraS3Request
}
func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.storageDomainId = storageDomainId
    getStorageDomainFailuresSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return getStorageDomainFailuresSpectraS3Request
}
func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithStorageDomainFailureType(storageDomainFailureType StorageDomainFailureType) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.storageDomainFailureType = storageDomainFailureType
    getStorageDomainFailuresSpectraS3Request.queryParams.Set("type", storageDomainFailureType.String())
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithErrorMessage(errorMessage *string) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.errorMessage = errorMessage
    if errorMessage != nil {
        getStorageDomainFailuresSpectraS3Request.queryParams.Set("error_message", *errorMessage)
    } else {
        getStorageDomainFailuresSpectraS3Request.queryParams.Set("error_message", "")
    }
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithLastPage() *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.queryParams.Set("last_page", "")
    return getStorageDomainFailuresSpectraS3Request
}

func (GetStorageDomainFailuresSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) Path() string {
    return "/_rest_/storage_domain_failure"
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) QueryParams() *url.Values {
    return getStorageDomainFailuresSpectraS3Request.queryParams
}

func (GetStorageDomainFailuresSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetStorageDomainFailuresSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetStorageDomainFailuresSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
