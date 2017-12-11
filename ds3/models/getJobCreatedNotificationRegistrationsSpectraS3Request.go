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

type GetJobCreatedNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetJobCreatedNotificationRegistrationsSpectraS3Request() *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    return &GetJobCreatedNotificationRegistrationsSpectraS3Request{
    }
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.LastPage = true
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

