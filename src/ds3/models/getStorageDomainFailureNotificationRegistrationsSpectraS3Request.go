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

type GetStorageDomainFailureNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetStorageDomainFailureNotificationRegistrationsSpectraS3Request() *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetStorageDomainFailureNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}
func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}
func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}
func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.userId = userId
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}


func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

func (GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/storage_domain_failure_notification_registration"
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
