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

type PutStorageDomainFailureNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutStorageDomainFailureNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutStorageDomainFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutStorageDomainFailureNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putStorageDomainFailureNotificationRegistrationSpectraS3Request *PutStorageDomainFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutStorageDomainFailureNotificationRegistrationSpectraS3Request {
    putStorageDomainFailureNotificationRegistrationSpectraS3Request.format = format
    putStorageDomainFailureNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putStorageDomainFailureNotificationRegistrationSpectraS3Request
}
func (putStorageDomainFailureNotificationRegistrationSpectraS3Request *PutStorageDomainFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutStorageDomainFailureNotificationRegistrationSpectraS3Request {
    putStorageDomainFailureNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putStorageDomainFailureNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putStorageDomainFailureNotificationRegistrationSpectraS3Request
}
func (putStorageDomainFailureNotificationRegistrationSpectraS3Request *PutStorageDomainFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutStorageDomainFailureNotificationRegistrationSpectraS3Request {
    putStorageDomainFailureNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putStorageDomainFailureNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putStorageDomainFailureNotificationRegistrationSpectraS3Request
}



func (PutStorageDomainFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putStorageDomainFailureNotificationRegistrationSpectraS3Request *PutStorageDomainFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/storage_domain_failure_notification_registration"
}

func (putStorageDomainFailureNotificationRegistrationSpectraS3Request *PutStorageDomainFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putStorageDomainFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (PutStorageDomainFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutStorageDomainFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutStorageDomainFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
