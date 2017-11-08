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

type PutJobCreationFailedNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutJobCreationFailedNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    return &PutJobCreationFailedNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    putJobCreationFailedNotificationRegistrationSpectraS3Request.Format = format
    return putJobCreationFailedNotificationRegistrationSpectraS3Request
}

func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    putJobCreationFailedNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putJobCreationFailedNotificationRegistrationSpectraS3Request
}

func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    putJobCreationFailedNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putJobCreationFailedNotificationRegistrationSpectraS3Request
}

