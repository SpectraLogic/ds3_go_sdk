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
    "strconv"
)

type GetJobChunksReadyForClientProcessingSpectraS3Request struct {
    job string
    jobChunk string
    preferredNumberOfChunks int
    queryParams *url.Values
}

func NewGetJobChunksReadyForClientProcessingSpectraS3Request(job string) *GetJobChunksReadyForClientProcessingSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("job", job)

    return &GetJobChunksReadyForClientProcessingSpectraS3Request{
        job: job,
        queryParams: queryParams,
    }
}

func (getJobChunksReadyForClientProcessingSpectraS3Request *GetJobChunksReadyForClientProcessingSpectraS3Request) WithJobChunk(jobChunk string) *GetJobChunksReadyForClientProcessingSpectraS3Request {
    getJobChunksReadyForClientProcessingSpectraS3Request.jobChunk = jobChunk
    getJobChunksReadyForClientProcessingSpectraS3Request.queryParams.Set("job_chunk", jobChunk)
    return getJobChunksReadyForClientProcessingSpectraS3Request
}
func (getJobChunksReadyForClientProcessingSpectraS3Request *GetJobChunksReadyForClientProcessingSpectraS3Request) WithPreferredNumberOfChunks(preferredNumberOfChunks int) *GetJobChunksReadyForClientProcessingSpectraS3Request {
    getJobChunksReadyForClientProcessingSpectraS3Request.preferredNumberOfChunks = preferredNumberOfChunks
    getJobChunksReadyForClientProcessingSpectraS3Request.queryParams.Set("preferred_number_of_chunks", strconv.Itoa(preferredNumberOfChunks))
    return getJobChunksReadyForClientProcessingSpectraS3Request
}



func (GetJobChunksReadyForClientProcessingSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getJobChunksReadyForClientProcessingSpectraS3Request *GetJobChunksReadyForClientProcessingSpectraS3Request) Path() string {
    return "/_rest_/job_chunk"
}

func (getJobChunksReadyForClientProcessingSpectraS3Request *GetJobChunksReadyForClientProcessingSpectraS3Request) QueryParams() *url.Values {
    return getJobChunksReadyForClientProcessingSpectraS3Request.queryParams
}

func (GetJobChunksReadyForClientProcessingSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetJobChunksReadyForClientProcessingSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetJobChunksReadyForClientProcessingSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
