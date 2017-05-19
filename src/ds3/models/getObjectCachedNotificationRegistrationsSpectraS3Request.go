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

type GetObjectCachedNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetObjectCachedNotificationRegistrationsSpectraS3Request() *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetObjectCachedNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getObjectCachedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}
func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getObjectCachedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}
func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getObjectCachedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}
func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.userId = userId
    getObjectCachedNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}


func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

func (GetObjectCachedNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/object_cached_notification_registration"
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getObjectCachedNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetObjectCachedNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetObjectCachedNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetObjectCachedNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
