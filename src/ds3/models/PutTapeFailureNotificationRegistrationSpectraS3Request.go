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

type PutTapeFailureNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutTapeFailureNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutTapeFailureNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    putTapeFailureNotificationRegistrationSpectraS3Request.format = format
    putTapeFailureNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putTapeFailureNotificationRegistrationSpectraS3Request
}
func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    putTapeFailureNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putTapeFailureNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putTapeFailureNotificationRegistrationSpectraS3Request
}
func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    putTapeFailureNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putTapeFailureNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putTapeFailureNotificationRegistrationSpectraS3Request
}



func (PutTapeFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/tape_failure_notification_registration"
}

func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putTapeFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (PutTapeFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutTapeFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutTapeFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
