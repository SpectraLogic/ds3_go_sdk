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

type GetStorageDomainFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetStorageDomainFailureNotificationRegistrationsSpectraS3Request() *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    return &GetStorageDomainFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

