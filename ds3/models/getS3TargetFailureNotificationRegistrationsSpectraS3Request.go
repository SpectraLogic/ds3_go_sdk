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

type GetS3TargetFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetS3TargetFailureNotificationRegistrationsSpectraS3Request() *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    return &GetS3TargetFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

