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

type TruncateJobSpectraS3Request struct {
    jobId string
    queryParams *url.Values
}

func NewTruncateJobSpectraS3Request(jobId string) *TruncateJobSpectraS3Request {
    queryParams := &url.Values{}

    return &TruncateJobSpectraS3Request{
        jobId: jobId,
        queryParams: queryParams,
    }
}




func (TruncateJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (truncateJobSpectraS3Request *TruncateJobSpectraS3Request) Path() string {
    return "/_rest_/job/" + truncateJobSpectraS3Request.jobId
}

func (truncateJobSpectraS3Request *TruncateJobSpectraS3Request) QueryParams() *url.Values {
    return truncateJobSpectraS3Request.queryParams
}

func (TruncateJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (TruncateJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (TruncateJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
