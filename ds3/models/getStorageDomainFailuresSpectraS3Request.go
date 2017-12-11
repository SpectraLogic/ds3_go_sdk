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

type GetStorageDomainFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    StorageDomainFailureType StorageDomainFailureType
    StorageDomainId *string
}

func NewGetStorageDomainFailuresSpectraS3Request() *GetStorageDomainFailuresSpectraS3Request {
    return &GetStorageDomainFailuresSpectraS3Request{
    }
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithLastPage() *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.LastPage = true
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.PageLength = &pageLength
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.PageOffset = &pageOffset
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.StorageDomainId = &storageDomainId
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithStorageDomainFailureType(storageDomainFailureType StorageDomainFailureType) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.StorageDomainFailureType = storageDomainFailureType
    return getStorageDomainFailuresSpectraS3Request
}

