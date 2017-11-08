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

type PutPoolFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutPoolFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    return &PutPoolFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    putPoolFailureNotificationRegistrationSpectraS3Request.Format = format
    return putPoolFailureNotificationRegistrationSpectraS3Request
}

func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    putPoolFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putPoolFailureNotificationRegistrationSpectraS3Request
}

func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    putPoolFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putPoolFailureNotificationRegistrationSpectraS3Request
}

