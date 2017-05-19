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

type GetObjectLostNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetObjectLostNotificationRegistrationsSpectraS3Request() *GetObjectLostNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetObjectLostNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getObjectLostNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getObjectLostNotificationRegistrationsSpectraS3Request
}
func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getObjectLostNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getObjectLostNotificationRegistrationsSpectraS3Request
}
func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getObjectLostNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getObjectLostNotificationRegistrationsSpectraS3Request
}
func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.userId = userId
    getObjectLostNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getObjectLostNotificationRegistrationsSpectraS3Request
}


func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithLastPage() *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

func (GetObjectLostNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/object_lost_notification_registration"
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getObjectLostNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetObjectLostNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetObjectLostNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetObjectLostNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
