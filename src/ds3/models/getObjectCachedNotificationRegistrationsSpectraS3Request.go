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

type GetObjectCachedNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetObjectCachedNotificationRegistrationsSpectraS3Request() *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    return &GetObjectCachedNotificationRegistrationsSpectraS3Request{
    }
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.LastPage = true
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

