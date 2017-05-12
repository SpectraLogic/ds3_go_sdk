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

type PutDs3TargetFailureNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutDs3TargetFailureNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutDs3TargetFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutDs3TargetFailureNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putDs3TargetFailureNotificationRegistrationSpectraS3Request *PutDs3TargetFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutDs3TargetFailureNotificationRegistrationSpectraS3Request {
    putDs3TargetFailureNotificationRegistrationSpectraS3Request.format = format
    putDs3TargetFailureNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putDs3TargetFailureNotificationRegistrationSpectraS3Request
}
func (putDs3TargetFailureNotificationRegistrationSpectraS3Request *PutDs3TargetFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutDs3TargetFailureNotificationRegistrationSpectraS3Request {
    putDs3TargetFailureNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putDs3TargetFailureNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putDs3TargetFailureNotificationRegistrationSpectraS3Request
}
func (putDs3TargetFailureNotificationRegistrationSpectraS3Request *PutDs3TargetFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutDs3TargetFailureNotificationRegistrationSpectraS3Request {
    putDs3TargetFailureNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putDs3TargetFailureNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putDs3TargetFailureNotificationRegistrationSpectraS3Request
}



func (PutDs3TargetFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putDs3TargetFailureNotificationRegistrationSpectraS3Request *PutDs3TargetFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/ds3_target_failure_notification_registration"
}

func (putDs3TargetFailureNotificationRegistrationSpectraS3Request *PutDs3TargetFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putDs3TargetFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (PutDs3TargetFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutDs3TargetFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutDs3TargetFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
