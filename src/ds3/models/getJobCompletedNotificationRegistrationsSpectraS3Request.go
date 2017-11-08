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

type GetJobCompletedNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetJobCompletedNotificationRegistrationsSpectraS3Request() *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    return &GetJobCompletedNotificationRegistrationsSpectraS3Request{
    }
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.LastPage = true
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

