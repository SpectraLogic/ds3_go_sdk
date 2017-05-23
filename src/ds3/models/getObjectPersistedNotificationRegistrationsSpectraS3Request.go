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

type GetObjectPersistedNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetObjectPersistedNotificationRegistrationsSpectraS3Request() *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetObjectPersistedNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    getObjectPersistedNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getObjectPersistedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getObjectPersistedNotificationRegistrationsSpectraS3Request
}
func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    getObjectPersistedNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getObjectPersistedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getObjectPersistedNotificationRegistrationsSpectraS3Request
}
func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    getObjectPersistedNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getObjectPersistedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getObjectPersistedNotificationRegistrationsSpectraS3Request
}
func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    getObjectPersistedNotificationRegistrationsSpectraS3Request.userId = userId
    getObjectPersistedNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getObjectPersistedNotificationRegistrationsSpectraS3Request
}


func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    getObjectPersistedNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getObjectPersistedNotificationRegistrationsSpectraS3Request
}

func (GetObjectPersistedNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/object_persisted_notification_registration"
}

func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getObjectPersistedNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetObjectPersistedNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetObjectPersistedNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetObjectPersistedNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
