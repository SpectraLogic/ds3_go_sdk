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

type PutJobCompletedNotificationRegistrationSpectraS3Request struct {
    format HttpResponseFormatType
    jobId string
    namingConvention NamingConventionType
    notificationEndPoint *string
    notificationHttpMethod RequestType
    queryParams *url.Values
}

func NewPutJobCompletedNotificationRegistrationSpectraS3Request(notificationEndPoint *string) *PutJobCompletedNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("notification_end_point", *notificationEndPoint)

    return &PutJobCompletedNotificationRegistrationSpectraS3Request{
        notificationEndPoint: notificationEndPoint,
        queryParams: queryParams,
    }
}

func (putJobCompletedNotificationRegistrationSpectraS3Request *PutJobCompletedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutJobCompletedNotificationRegistrationSpectraS3Request {
    putJobCompletedNotificationRegistrationSpectraS3Request.format = format
    putJobCompletedNotificationRegistrationSpectraS3Request.queryParams.Set("format", format.String())
    return putJobCompletedNotificationRegistrationSpectraS3Request
}
func (putJobCompletedNotificationRegistrationSpectraS3Request *PutJobCompletedNotificationRegistrationSpectraS3Request) WithJobId(jobId string) *PutJobCompletedNotificationRegistrationSpectraS3Request {
    putJobCompletedNotificationRegistrationSpectraS3Request.jobId = jobId
    putJobCompletedNotificationRegistrationSpectraS3Request.queryParams.Set("job_id", jobId)
    return putJobCompletedNotificationRegistrationSpectraS3Request
}
func (putJobCompletedNotificationRegistrationSpectraS3Request *PutJobCompletedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutJobCompletedNotificationRegistrationSpectraS3Request {
    putJobCompletedNotificationRegistrationSpectraS3Request.namingConvention = namingConvention
    putJobCompletedNotificationRegistrationSpectraS3Request.queryParams.Set("naming_convention", namingConvention.String())
    return putJobCompletedNotificationRegistrationSpectraS3Request
}
func (putJobCompletedNotificationRegistrationSpectraS3Request *PutJobCompletedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutJobCompletedNotificationRegistrationSpectraS3Request {
    putJobCompletedNotificationRegistrationSpectraS3Request.notificationHttpMethod = notificationHttpMethod
    putJobCompletedNotificationRegistrationSpectraS3Request.queryParams.Set("notification_http_method", notificationHttpMethod.String())
    return putJobCompletedNotificationRegistrationSpectraS3Request
}



func (PutJobCompletedNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putJobCompletedNotificationRegistrationSpectraS3Request *PutJobCompletedNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/job_completed_notification_registration"
}

func (putJobCompletedNotificationRegistrationSpectraS3Request *PutJobCompletedNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return putJobCompletedNotificationRegistrationSpectraS3Request.queryParams
}

func (PutJobCompletedNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutJobCompletedNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutJobCompletedNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
