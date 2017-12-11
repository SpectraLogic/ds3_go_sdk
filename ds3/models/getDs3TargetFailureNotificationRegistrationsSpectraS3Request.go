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

type GetDs3TargetFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetDs3TargetFailureNotificationRegistrationsSpectraS3Request() *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    return &GetDs3TargetFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

