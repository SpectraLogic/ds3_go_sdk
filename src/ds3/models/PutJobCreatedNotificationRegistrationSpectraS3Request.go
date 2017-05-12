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

type PutJobCreatedNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutJobCreatedNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutJobCreatedNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutJobCreatedNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putJobCreatedNotificationRegistrationSpectraS3Request *PutJobCreatedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutJobCreatedNotificationRegistrationSpectraS3Request {
    putJobCreatedNotificationRegistrationSpectraS3Request.format = format
    putJobCreatedNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putJobCreatedNotificationRegistrationSpectraS3Request
}
func (putJobCreatedNotificationRegistrationSpectraS3Request *PutJobCreatedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutJobCreatedNotificationRegistrationSpectraS3Request {
    putJobCreatedNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putJobCreatedNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putJobCreatedNotificationRegistrationSpectraS3Request
}
func (putJobCreatedNotificationRegistrationSpectraS3Request *PutJobCreatedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutJobCreatedNotificationRegistrationSpectraS3Request {
    putJobCreatedNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putJobCreatedNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putJobCreatedNotificationRegistrationSpectraS3Request
}



func (PutJobCreatedNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putJobCreatedNotificationRegistrationSpectraS3Request *PutJobCreatedNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/job_created_notification_registration"
}

func (putJobCreatedNotificationRegistrationSpectraS3Request *PutJobCreatedNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putJobCreatedNotificationRegistrationSpectraS3Request.queryParams
}

func (PutJobCreatedNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutJobCreatedNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutJobCreatedNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
