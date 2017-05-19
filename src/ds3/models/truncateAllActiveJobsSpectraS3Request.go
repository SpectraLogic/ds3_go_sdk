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

type TruncateAllActiveJobsSpectraS3Request struct {
    bucketId string
    requestType JobRequestType
    queryParams *url.Values
}

func NewTruncateAllActiveJobsSpectraS3Request() *TruncateAllActiveJobsSpectraS3Request {
    queryParams := &url.Values{}

    return &TruncateAllActiveJobsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (truncateAllActiveJobsSpectraS3Request *TruncateAllActiveJobsSpectraS3Request) WithBucketId(bucketId string) *TruncateAllActiveJobsSpectraS3Request {
    truncateAllActiveJobsSpectraS3Request.bucketId = bucketId
    truncateAllActiveJobsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return truncateAllActiveJobsSpectraS3Request
}
func (truncateAllActiveJobsSpectraS3Request *TruncateAllActiveJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *TruncateAllActiveJobsSpectraS3Request {
    truncateAllActiveJobsSpectraS3Request.requestType = requestType
    truncateAllActiveJobsSpectraS3Request.queryParams.Set("request_type", requestType.String())
    return truncateAllActiveJobsSpectraS3Request
}



func (TruncateAllActiveJobsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (truncateAllActiveJobsSpectraS3Request *TruncateAllActiveJobsSpectraS3Request) Path() string {
    return "/_rest_/active_job"
}

func (truncateAllActiveJobsSpectraS3Request *TruncateAllActiveJobsSpectraS3Request) QueryParams() *url.Values {
    return truncateAllActiveJobsSpectraS3Request.queryParams
}

func (TruncateAllActiveJobsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (TruncateAllActiveJobsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (TruncateAllActiveJobsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
