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

type PutObjectCachedNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    jobId string
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutObjectCachedNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutObjectCachedNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutObjectCachedNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putObjectCachedNotificationRegistrationSpectraS3Request *PutObjectCachedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutObjectCachedNotificationRegistrationSpectraS3Request {
    putObjectCachedNotificationRegistrationSpectraS3Request.format = format
    putObjectCachedNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putObjectCachedNotificationRegistrationSpectraS3Request
}
func (putObjectCachedNotificationRegistrationSpectraS3Request *PutObjectCachedNotificationRegistrationSpectraS3Request) WithJobId(jobId string) *PutObjectCachedNotificationRegistrationSpectraS3Request {
    putObjectCachedNotificationRegistrationSpectraS3Request.jobId = jobId
    putObjectCachedNotificationRegistrationSpectraS3Request.queryParams.Set("job_id", jobId)
    return putObjectCachedNotificationRegistrationSpectraS3Request
}
func (putObjectCachedNotificationRegistrationSpectraS3Request *PutObjectCachedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutObjectCachedNotificationRegistrationSpectraS3Request {
    putObjectCachedNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putObjectCachedNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putObjectCachedNotificationRegistrationSpectraS3Request
}
func (putObjectCachedNotificationRegistrationSpectraS3Request *PutObjectCachedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutObjectCachedNotificationRegistrationSpectraS3Request {
    putObjectCachedNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putObjectCachedNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putObjectCachedNotificationRegistrationSpectraS3Request
}



func (PutObjectCachedNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putObjectCachedNotificationRegistrationSpectraS3Request *PutObjectCachedNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/object_cached_notification_registration"
}

func (putObjectCachedNotificationRegistrationSpectraS3Request *PutObjectCachedNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putObjectCachedNotificationRegistrationSpectraS3Request.queryParams
}

func (PutObjectCachedNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutObjectCachedNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutObjectCachedNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
