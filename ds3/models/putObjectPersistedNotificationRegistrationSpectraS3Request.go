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

type PutObjectPersistedNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    JobId *string
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutObjectPersistedNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    return &PutObjectPersistedNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.Format = format
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithJobId(jobId string) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.JobId = &jobId
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}

