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

type PutObjectLostNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutObjectLostNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutObjectLostNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutObjectLostNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutObjectLostNotificationRegistrationSpectraS3Request {
    putObjectLostNotificationRegistrationSpectraS3Request.format = format
    putObjectLostNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putObjectLostNotificationRegistrationSpectraS3Request
}
func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutObjectLostNotificationRegistrationSpectraS3Request {
    putObjectLostNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putObjectLostNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putObjectLostNotificationRegistrationSpectraS3Request
}
func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutObjectLostNotificationRegistrationSpectraS3Request {
    putObjectLostNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putObjectLostNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putObjectLostNotificationRegistrationSpectraS3Request
}



func (PutObjectLostNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/object_lost_notification_registration"
}

func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putObjectLostNotificationRegistrationSpectraS3Request.queryParams
}

func (PutObjectLostNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutObjectLostNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutObjectLostNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
