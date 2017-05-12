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

type CancelAllActiveJobsSpectraS3Request struct {
    bucketId string
    requestType JobRequestType
    queryParams *url.Values
}

func NewCancelAllActiveJobsSpectraS3Request() *CancelAllActiveJobsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("force", "")

    return &CancelAllActiveJobsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (cancelAllActiveJobsSpectraS3Request *CancelAllActiveJobsSpectraS3Request) WithBucketId(bucketId string) *CancelAllActiveJobsSpectraS3Request {
    cancelAllActiveJobsSpectraS3Request.bucketId = bucketId
    cancelAllActiveJobsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return cancelAllActiveJobsSpectraS3Request
}
func (cancelAllActiveJobsSpectraS3Request *CancelAllActiveJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *CancelAllActiveJobsSpectraS3Request {
    cancelAllActiveJobsSpectraS3Request.requestType = requestType
    cancelAllActiveJobsSpectraS3Request.queryParams.Set("request_type", requestType.String())
    return cancelAllActiveJobsSpectraS3Request
}



func (CancelAllActiveJobsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (cancelAllActiveJobsSpectraS3Request *CancelAllActiveJobsSpectraS3Request) Path() string {
    return "/_rest_/active_job"
}

func (cancelAllActiveJobsSpectraS3Request *CancelAllActiveJobsSpectraS3Request) QueryParams() *url.Values {
    return cancelAllActiveJobsSpectraS3Request.queryParams
}

func (CancelAllActiveJobsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (CancelAllActiveJobsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (CancelAllActiveJobsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
