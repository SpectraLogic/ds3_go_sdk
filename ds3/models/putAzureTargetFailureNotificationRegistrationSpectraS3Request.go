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

type PutAzureTargetFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutAzureTargetFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    return &PutAzureTargetFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.Format = format
    return putAzureTargetFailureNotificationRegistrationSpectraS3Request
}

func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putAzureTargetFailureNotificationRegistrationSpectraS3Request
}

func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putAzureTargetFailureNotificationRegistrationSpectraS3Request
}

