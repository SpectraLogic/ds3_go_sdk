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

type GetPoolFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetPoolFailureNotificationRegistrationsSpectraS3Request() *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    return &GetPoolFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

