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

type CancelJobSpectraS3Request struct {
    jobId string
    queryParams *url.Values
}

func NewCancelJobSpectraS3Request(jobId string) *CancelJobSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("force", "")

    return &CancelJobSpectraS3Request{
        jobId: jobId,
        queryParams: queryParams,
    }
}




func (CancelJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (cancelJobSpectraS3Request *CancelJobSpectraS3Request) Path() string {
    return "/_rest_/job/" + cancelJobSpectraS3Request.jobId
}

func (cancelJobSpectraS3Request *CancelJobSpectraS3Request) QueryParams() *url.Values {
    return cancelJobSpectraS3Request.queryParams
}

func (CancelJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (CancelJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (CancelJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
