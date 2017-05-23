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

type DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request struct {
    notificationId string
    queryParams *url.Values
}

func NewDeleteStorageDomainFailureNotificationRegistrationSpectraS3Request(notificationId string) *DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request{
        notificationId: notificationId,
        queryParams: queryParams,
    }
}




func (DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteStorageDomainFailureNotificationRegistrationSpectraS3Request *DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/storage_domain_failure_notification_registration/" + deleteStorageDomainFailureNotificationRegistrationSpectraS3Request.notificationId
}

func (deleteStorageDomainFailureNotificationRegistrationSpectraS3Request *DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return deleteStorageDomainFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
