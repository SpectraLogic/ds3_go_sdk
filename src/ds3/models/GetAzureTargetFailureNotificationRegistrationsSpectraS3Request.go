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

type GetAzureTargetFailureNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetAzureTargetFailureNotificationRegistrationsSpectraS3Request() *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetAzureTargetFailureNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}
func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}
func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}
func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.userId = userId
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}


func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

func (GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/azure_target_failure_notification_registration"
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
