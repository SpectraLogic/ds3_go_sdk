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

type PutJobCreationFailedNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutJobCreationFailedNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutJobCreationFailedNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    putJobCreationFailedNotificationRegistrationSpectraS3Request.format = format
    putJobCreationFailedNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putJobCreationFailedNotificationRegistrationSpectraS3Request
}
func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    putJobCreationFailedNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putJobCreationFailedNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putJobCreationFailedNotificationRegistrationSpectraS3Request
}
func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    putJobCreationFailedNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putJobCreationFailedNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putJobCreationFailedNotificationRegistrationSpectraS3Request
}



func (PutJobCreationFailedNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/job_creation_failed_notification_registration"
}

func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putJobCreationFailedNotificationRegistrationSpectraS3Request.queryParams
}

func (PutJobCreationFailedNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutJobCreationFailedNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutJobCreationFailedNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
