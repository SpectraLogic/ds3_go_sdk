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

type GetTapePartitionFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetTapePartitionFailureNotificationRegistrationsSpectraS3Request() *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    return &GetTapePartitionFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

