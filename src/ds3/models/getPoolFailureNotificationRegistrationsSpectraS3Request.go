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

type GetPoolFailureNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetPoolFailureNotificationRegistrationsSpectraS3Request() *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetPoolFailureNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getPoolFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}
func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getPoolFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}
func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getPoolFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}
func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.userId = userId
    getPoolFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}


func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

func (GetPoolFailureNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/pool_failure_notification_registration"
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getPoolFailureNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetPoolFailureNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetPoolFailureNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetPoolFailureNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
