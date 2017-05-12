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

type DeleteObjectPersistedNotificationRegistrationSpectraS3Request struct {
    notificationId string
    queryParams *url.Values
}

func NewDeleteObjectPersistedNotificationRegistrationSpectraS3Request(notificationId string) *DeleteObjectPersistedNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteObjectPersistedNotificationRegistrationSpectraS3Request{
        notificationId: notificationId,
        queryParams: queryParams,
    }
}




func (DeleteObjectPersistedNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteObjectPersistedNotificationRegistrationSpectraS3Request *DeleteObjectPersistedNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/object_persisted_notification_registration/" + deleteObjectPersistedNotificationRegistrationSpectraS3Request.notificationId
}

func (deleteObjectPersistedNotificationRegistrationSpectraS3Request *DeleteObjectPersistedNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return deleteObjectPersistedNotificationRegistrationSpectraS3Request.queryParams
}

func (DeleteObjectPersistedNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteObjectPersistedNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteObjectPersistedNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
