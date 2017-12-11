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

type GetCompletedJobsSpectraS3Request struct {
    BucketId *string
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    Priority Priority
    Rechunked *string
    RequestType JobRequestType
    Truncated *bool
    UserId *string
}

func NewGetCompletedJobsSpectraS3Request() *GetCompletedJobsSpectraS3Request {
    return &GetCompletedJobsSpectraS3Request{
    }
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithBucketId(bucketId string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.BucketId = &bucketId
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.ChunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithLastPage() *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.LastPage = true
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithName(name string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.Name = &name
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPageLength(pageLength int) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.PageLength = &pageLength
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPageOffset(pageOffset int) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.PageOffset = &pageOffset
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPriority(priority Priority) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.Priority = priority
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithRechunked(rechunked string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.Rechunked = &rechunked
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.RequestType = requestType
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithTruncated(truncated bool) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.Truncated = &truncated
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithUserId(userId string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.UserId = &userId
    return getCompletedJobsSpectraS3Request
}

