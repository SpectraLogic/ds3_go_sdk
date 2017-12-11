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

type PutTapePartitionFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutTapePartitionFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    return &PutTapePartitionFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.Format = format
    return putTapePartitionFailureNotificationRegistrationSpectraS3Request
}

func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putTapePartitionFailureNotificationRegistrationSpectraS3Request
}

func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putTapePartitionFailureNotificationRegistrationSpectraS3Request
}

