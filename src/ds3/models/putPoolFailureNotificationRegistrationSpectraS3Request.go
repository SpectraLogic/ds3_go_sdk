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

type PutPoolFailureNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutPoolFailureNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutPoolFailureNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    putPoolFailureNotificationRegistrationSpectraS3Request.format = format
    putPoolFailureNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putPoolFailureNotificationRegistrationSpectraS3Request
}
func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    putPoolFailureNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putPoolFailureNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putPoolFailureNotificationRegistrationSpectraS3Request
}
func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    putPoolFailureNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putPoolFailureNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putPoolFailureNotificationRegistrationSpectraS3Request
}



func (PutPoolFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/pool_failure_notification_registration"
}

func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putPoolFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (PutPoolFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutPoolFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutPoolFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
