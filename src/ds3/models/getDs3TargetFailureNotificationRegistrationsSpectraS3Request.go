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

type GetDs3TargetFailureNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetDs3TargetFailureNotificationRegistrationsSpectraS3Request() *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDs3TargetFailureNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}
func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}
func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}
func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.userId = userId
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}


func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/ds3_target_failure_notification_registration"
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
