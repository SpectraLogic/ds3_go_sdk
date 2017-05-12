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

type GetTapePartitionFailureNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetTapePartitionFailureNotificationRegistrationsSpectraS3Request() *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapePartitionFailureNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}
func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}
func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}
func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.userId = userId
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}


func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

func (GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/tape_partition_failure_notification_registration"
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
