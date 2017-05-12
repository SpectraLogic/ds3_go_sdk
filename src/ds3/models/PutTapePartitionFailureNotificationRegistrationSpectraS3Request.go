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

type PutTapePartitionFailureNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutTapePartitionFailureNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutTapePartitionFailureNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.format = format
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putTapePartitionFailureNotificationRegistrationSpectraS3Request
}
func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putTapePartitionFailureNotificationRegistrationSpectraS3Request
}
func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putTapePartitionFailureNotificationRegistrationSpectraS3Request
}



func (PutTapePartitionFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/tape_partition_failure_notification_registration"
}

func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putTapePartitionFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (PutTapePartitionFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutTapePartitionFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutTapePartitionFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
