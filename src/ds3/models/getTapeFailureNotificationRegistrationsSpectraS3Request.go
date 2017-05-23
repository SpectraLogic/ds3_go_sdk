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

type GetTapeFailureNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetTapeFailureNotificationRegistrationsSpectraS3Request() *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapeFailureNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    getTapeFailureNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getTapeFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getTapeFailureNotificationRegistrationsSpectraS3Request
}
func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    getTapeFailureNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getTapeFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getTapeFailureNotificationRegistrationsSpectraS3Request
}
func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    getTapeFailureNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getTapeFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getTapeFailureNotificationRegistrationsSpectraS3Request
}
func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    getTapeFailureNotificationRegistrationsSpectraS3Request.userId = userId
    getTapeFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getTapeFailureNotificationRegistrationsSpectraS3Request
}


func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    getTapeFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getTapeFailureNotificationRegistrationsSpectraS3Request
}

func (GetTapeFailureNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/tape_failure_notification_registration"
}

func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getTapeFailureNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetTapeFailureNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapeFailureNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapeFailureNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
