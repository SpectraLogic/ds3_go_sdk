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

type GetJobsSpectraS3Request struct {
    bucketId string
    queryParams *url.Values
}

func NewGetJobsSpectraS3Request() *GetJobsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetJobsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getJobsSpectraS3Request *GetJobsSpectraS3Request) WithBucketId(bucketId string) *GetJobsSpectraS3Request {
    getJobsSpectraS3Request.bucketId = bucketId
    getJobsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getJobsSpectraS3Request
}


func (getJobsSpectraS3Request *GetJobsSpectraS3Request) WithFullDetails() *GetJobsSpectraS3Request {
    getJobsSpectraS3Request.queryParams.Set("full_details", "")
    return getJobsSpectraS3Request
}

func (GetJobsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getJobsSpectraS3Request *GetJobsSpectraS3Request) Path() string {
    return "/_rest_/job"
}

func (getJobsSpectraS3Request *GetJobsSpectraS3Request) QueryParams() *url.Values {
    return getJobsSpectraS3Request.queryParams
}

func (GetJobsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetJobsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetJobsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
