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

type GetJobCreationFailedNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetJobCreationFailedNotificationRegistrationsSpectraS3Request() *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    return &GetJobCreationFailedNotificationRegistrationsSpectraS3Request{
    }
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.LastPage = true
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

