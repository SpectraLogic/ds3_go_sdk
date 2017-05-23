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

type GetJobCompletedNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetJobCompletedNotificationRegistrationsSpectraS3Request() *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetJobCompletedNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getJobCompletedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}
func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getJobCompletedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}
func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getJobCompletedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}
func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.userId = userId
    getJobCompletedNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}


func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

func (GetJobCompletedNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/job_completed_notification_registration"
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getJobCompletedNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetJobCompletedNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetJobCompletedNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetJobCompletedNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
