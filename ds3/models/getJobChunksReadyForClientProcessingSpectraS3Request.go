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

type GetJobChunksReadyForClientProcessingSpectraS3Request struct {
    Job string
    JobChunk *string
    PreferredNumberOfChunks *int
}

func NewGetJobChunksReadyForClientProcessingSpectraS3Request(job string) *GetJobChunksReadyForClientProcessingSpectraS3Request {
    return &GetJobChunksReadyForClientProcessingSpectraS3Request{
        Job: job,
    }
}

func (getJobChunksReadyForClientProcessingSpectraS3Request *GetJobChunksReadyForClientProcessingSpectraS3Request) WithJobChunk(jobChunk string) *GetJobChunksReadyForClientProcessingSpectraS3Request {
    getJobChunksReadyForClientProcessingSpectraS3Request.JobChunk = &jobChunk
    return getJobChunksReadyForClientProcessingSpectraS3Request
}

func (getJobChunksReadyForClientProcessingSpectraS3Request *GetJobChunksReadyForClientProcessingSpectraS3Request) WithPreferredNumberOfChunks(preferredNumberOfChunks int) *GetJobChunksReadyForClientProcessingSpectraS3Request {
    getJobChunksReadyForClientProcessingSpectraS3Request.PreferredNumberOfChunks = &preferredNumberOfChunks
    return getJobChunksReadyForClientProcessingSpectraS3Request
}

