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

type PutSystemFailureNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutSystemFailureNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutSystemFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutSystemFailureNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putSystemFailureNotificationRegistrationSpectraS3Request *PutSystemFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutSystemFailureNotificationRegistrationSpectraS3Request {
    putSystemFailureNotificationRegistrationSpectraS3Request.format = format
    putSystemFailureNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putSystemFailureNotificationRegistrationSpectraS3Request
}
func (putSystemFailureNotificationRegistrationSpectraS3Request *PutSystemFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutSystemFailureNotificationRegistrationSpectraS3Request {
    putSystemFailureNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putSystemFailureNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putSystemFailureNotificationRegistrationSpectraS3Request
}
func (putSystemFailureNotificationRegistrationSpectraS3Request *PutSystemFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutSystemFailureNotificationRegistrationSpectraS3Request {
    putSystemFailureNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putSystemFailureNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putSystemFailureNotificationRegistrationSpectraS3Request
}



func (PutSystemFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putSystemFailureNotificationRegistrationSpectraS3Request *PutSystemFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/system_failure_notification_registration"
}

func (putSystemFailureNotificationRegistrationSpectraS3Request *PutSystemFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putSystemFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (PutSystemFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutSystemFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutSystemFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
