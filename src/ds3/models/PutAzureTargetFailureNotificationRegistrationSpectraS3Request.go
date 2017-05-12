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

type PutAzureTargetFailureNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutAzureTargetFailureNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutAzureTargetFailureNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.format = format
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putAzureTargetFailureNotificationRegistrationSpectraS3Request
}
func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putAzureTargetFailureNotificationRegistrationSpectraS3Request
}
func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putAzureTargetFailureNotificationRegistrationSpectraS3Request
}



func (PutAzureTargetFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/azure_target_failure_notification_registration"
}

func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putAzureTargetFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (PutAzureTargetFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutAzureTargetFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutAzureTargetFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
