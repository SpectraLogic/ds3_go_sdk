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

type GetJobCreatedNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetJobCreatedNotificationRegistrationsSpectraS3Request() *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetJobCreatedNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getJobCreatedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}
func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getJobCreatedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}
func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getJobCreatedNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}
func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.userId = userId
    getJobCreatedNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}


func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

func (GetJobCreatedNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/job_created_notification_registration"
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getJobCreatedNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetJobCreatedNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetJobCreatedNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetJobCreatedNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
