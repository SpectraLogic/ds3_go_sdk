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

type PutObjectPersistedNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    jobId string
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutObjectPersistedNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutObjectPersistedNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.format = format
    putObjectPersistedNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}
func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithJobId(jobId string) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.jobId = jobId
    putObjectPersistedNotificationRegistrationSpectraS3Request.queryParams.Set("job_id", jobId)
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}
func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putObjectPersistedNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}
func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putObjectPersistedNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}



func (PutObjectPersistedNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/object_persisted_notification_registration"
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putObjectPersistedNotificationRegistrationSpectraS3Request.queryParams
}

func (PutObjectPersistedNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutObjectPersistedNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutObjectPersistedNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
