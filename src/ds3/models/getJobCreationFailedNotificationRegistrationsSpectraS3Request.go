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

type GetJobCreationFailedNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetJobCreationFailedNotificationRegistrationsSpectraS3Request() *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetJobCreationFailedNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}
func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}
func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}
func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.userId = userId
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}


func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

func (GetJobCreationFailedNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/job_creation_failed_notification_registration"
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetJobCreationFailedNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetJobCreationFailedNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetJobCreationFailedNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
