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

type PutObjectLostNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutObjectLostNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutObjectLostNotificationRegistrationSpectraS3Request {
    return &PutObjectLostNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutObjectLostNotificationRegistrationSpectraS3Request {
    putObjectLostNotificationRegistrationSpectraS3Request.Format = format
    return putObjectLostNotificationRegistrationSpectraS3Request
}

func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutObjectLostNotificationRegistrationSpectraS3Request {
    putObjectLostNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putObjectLostNotificationRegistrationSpectraS3Request
}

func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutObjectLostNotificationRegistrationSpectraS3Request {
    putObjectLostNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putObjectLostNotificationRegistrationSpectraS3Request
}

