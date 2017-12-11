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

type GetActiveJobsSpectraS3Request struct {
    Aggregating *bool
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

func NewGetActiveJobsSpectraS3Request() *GetActiveJobsSpectraS3Request {
    return &GetActiveJobsSpectraS3Request{
    }
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithAggregating(aggregating bool) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.Aggregating = &aggregating
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithBucketId(bucketId string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.BucketId = &bucketId
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.ChunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithLastPage() *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.LastPage = true
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithName(name string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.Name = &name
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPageLength(pageLength int) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.PageLength = &pageLength
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPageOffset(pageOffset int) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.PageOffset = &pageOffset
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPriority(priority Priority) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.Priority = priority
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithRechunked(rechunked string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.Rechunked = &rechunked
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.RequestType = requestType
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithTruncated(truncated bool) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.Truncated = &truncated
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithUserId(userId string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.UserId = &userId
    return getActiveJobsSpectraS3Request
}

