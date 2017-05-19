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

type CancelAllJobsSpectraS3Request struct {
    bucketId string
    requestType JobRequestType
    queryParams *url.Values
}

func NewCancelAllJobsSpectraS3Request() *CancelAllJobsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("force", "")

    return &CancelAllJobsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (cancelAllJobsSpectraS3Request *CancelAllJobsSpectraS3Request) WithBucketId(bucketId string) *CancelAllJobsSpectraS3Request {
    cancelAllJobsSpectraS3Request.bucketId = bucketId
    cancelAllJobsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return cancelAllJobsSpectraS3Request
}
func (cancelAllJobsSpectraS3Request *CancelAllJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *CancelAllJobsSpectraS3Request {
    cancelAllJobsSpectraS3Request.requestType = requestType
    cancelAllJobsSpectraS3Request.queryParams.Set("request_type", requestType.String())
    return cancelAllJobsSpectraS3Request
}



func (CancelAllJobsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (cancelAllJobsSpectraS3Request *CancelAllJobsSpectraS3Request) Path() string {
    return "/_rest_/job"
}

func (cancelAllJobsSpectraS3Request *CancelAllJobsSpectraS3Request) QueryParams() *url.Values {
    return cancelAllJobsSpectraS3Request.queryParams
}

func (CancelAllJobsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (CancelAllJobsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (CancelAllJobsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
