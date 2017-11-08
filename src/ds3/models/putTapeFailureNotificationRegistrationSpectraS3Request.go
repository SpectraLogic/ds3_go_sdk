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

type PutTapeFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutTapeFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    return &PutTapeFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    putTapeFailureNotificationRegistrationSpectraS3Request.Format = format
    return putTapeFailureNotificationRegistrationSpectraS3Request
}

func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    putTapeFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putTapeFailureNotificationRegistrationSpectraS3Request
}

func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    putTapeFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putTapeFailureNotificationRegistrationSpectraS3Request
}

