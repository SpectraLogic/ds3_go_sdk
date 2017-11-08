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

type GetAzureTargetFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetAzureTargetFailureNotificationRegistrationsSpectraS3Request() *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    return &GetAzureTargetFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

