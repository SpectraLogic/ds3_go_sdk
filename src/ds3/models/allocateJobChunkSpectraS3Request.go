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

type AllocateJobChunkSpectraS3Request struct {
    jobChunkId string
    queryParams *url.Values
}

func NewAllocateJobChunkSpectraS3Request(jobChunkId string) *AllocateJobChunkSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "allocate")

    return &AllocateJobChunkSpectraS3Request{
        jobChunkId: jobChunkId,
        queryParams: queryParams,
    }
}




func (AllocateJobChunkSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (allocateJobChunkSpectraS3Request *AllocateJobChunkSpectraS3Request) Path() string {
    return "/_rest_/job_chunk/" + allocateJobChunkSpectraS3Request.jobChunkId
}

func (allocateJobChunkSpectraS3Request *AllocateJobChunkSpectraS3Request) QueryParams() *url.Values {
    return allocateJobChunkSpectraS3Request.queryParams
}

func (AllocateJobChunkSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (AllocateJobChunkSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (AllocateJobChunkSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
