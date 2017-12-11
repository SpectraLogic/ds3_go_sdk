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

type GetObjectLostNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetObjectLostNotificationRegistrationsSpectraS3Request() *GetObjectLostNotificationRegistrationsSpectraS3Request {
    return &GetObjectLostNotificationRegistrationsSpectraS3Request{
    }
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithLastPage() *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.LastPage = true
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

