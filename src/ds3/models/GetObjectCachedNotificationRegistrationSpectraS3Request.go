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

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type GetObjectCachedNotificationRegistrationSpectraS3Request struct {
    notificationId string
    queryParams *url.Values
}

func NewGetObjectCachedNotificationRegistrationSpectraS3Request(notificationId string) *GetObjectCachedNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}

    return &GetObjectCachedNotificationRegistrationSpectraS3Request{
        notificationId: notificationId,
        queryParams: queryParams,
    }
}




func (GetObjectCachedNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getObjectCachedNotificationRegistrationSpectraS3Request *GetObjectCachedNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/object_cached_notification_registration/" + getObjectCachedNotificationRegistrationSpectraS3Request.notificationId
}

func (getObjectCachedNotificationRegistrationSpectraS3Request *GetObjectCachedNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return getObjectCachedNotificationRegistrationSpectraS3Request.queryParams
}

func (GetObjectCachedNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetObjectCachedNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetObjectCachedNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
