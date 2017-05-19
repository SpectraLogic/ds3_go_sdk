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

type GetSystemFailureNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetSystemFailureNotificationRegistrationsSpectraS3Request() *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetSystemFailureNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getSystemFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}
func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getSystemFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}
func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getSystemFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}
func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.userId = userId
    getSystemFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}


func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

func (GetSystemFailureNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/system_failure_notification_registration"
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getSystemFailureNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetSystemFailureNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetSystemFailureNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetSystemFailureNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
