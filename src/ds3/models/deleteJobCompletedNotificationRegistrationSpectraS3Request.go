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

type DeleteJobCompletedNotificationRegistrationSpectraS3Request struct {
    notificationId string
    queryParams *url.Values
}

func NewDeleteJobCompletedNotificationRegistrationSpectraS3Request(notificationId string) *DeleteJobCompletedNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteJobCompletedNotificationRegistrationSpectraS3Request{
        notificationId: notificationId,
        queryParams: queryParams,
    }
}




func (DeleteJobCompletedNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteJobCompletedNotificationRegistrationSpectraS3Request *DeleteJobCompletedNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/job_completed_notification_registration/" + deleteJobCompletedNotificationRegistrationSpectraS3Request.notificationId
}

func (deleteJobCompletedNotificationRegistrationSpectraS3Request *DeleteJobCompletedNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return deleteJobCompletedNotificationRegistrationSpectraS3Request.queryParams
}

func (DeleteJobCompletedNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteJobCompletedNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteJobCompletedNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
