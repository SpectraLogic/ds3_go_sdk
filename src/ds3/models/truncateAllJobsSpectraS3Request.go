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

type TruncateAllJobsSpectraS3Request struct {
    bucketId string
    requestType JobRequestType
    queryParams *url.Values
}

func NewTruncateAllJobsSpectraS3Request() *TruncateAllJobsSpectraS3Request {
    queryParams := &url.Values{}

    return &TruncateAllJobsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (truncateAllJobsSpectraS3Request *TruncateAllJobsSpectraS3Request) WithBucketId(bucketId string) *TruncateAllJobsSpectraS3Request {
    truncateAllJobsSpectraS3Request.bucketId = bucketId
    truncateAllJobsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return truncateAllJobsSpectraS3Request
}
func (truncateAllJobsSpectraS3Request *TruncateAllJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *TruncateAllJobsSpectraS3Request {
    truncateAllJobsSpectraS3Request.requestType = requestType
    truncateAllJobsSpectraS3Request.queryParams.Set("request_type", requestType.String())
    return truncateAllJobsSpectraS3Request
}



func (TruncateAllJobsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (truncateAllJobsSpectraS3Request *TruncateAllJobsSpectraS3Request) Path() string {
    return "/_rest_/job"
}

func (truncateAllJobsSpectraS3Request *TruncateAllJobsSpectraS3Request) QueryParams() *url.Values {
    return truncateAllJobsSpectraS3Request.queryParams
}

func (TruncateAllJobsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (TruncateAllJobsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (TruncateAllJobsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
