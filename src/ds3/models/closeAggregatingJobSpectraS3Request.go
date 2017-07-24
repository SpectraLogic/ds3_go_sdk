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

type CloseAggregatingJobSpectraS3Request struct {
    jobId string
    queryParams *url.Values
}

func NewCloseAggregatingJobSpectraS3Request(jobId string) *CloseAggregatingJobSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("close_aggregating_job", "")

    return &CloseAggregatingJobSpectraS3Request{
        jobId: jobId,
        queryParams: queryParams,
    }
}




func (CloseAggregatingJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (closeAggregatingJobSpectraS3Request *CloseAggregatingJobSpectraS3Request) Path() string {
    return "/_rest_/job/" + closeAggregatingJobSpectraS3Request.jobId
}

func (closeAggregatingJobSpectraS3Request *CloseAggregatingJobSpectraS3Request) QueryParams() *url.Values {
    return closeAggregatingJobSpectraS3Request.queryParams
}

func (CloseAggregatingJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (CloseAggregatingJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (CloseAggregatingJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
