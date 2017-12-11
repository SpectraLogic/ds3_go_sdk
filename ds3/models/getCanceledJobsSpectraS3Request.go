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

type GetCanceledJobsSpectraS3Request struct {
    BucketId *string
    CanceledDueToTimeout *bool
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

func NewGetCanceledJobsSpectraS3Request() *GetCanceledJobsSpectraS3Request {
    return &GetCanceledJobsSpectraS3Request{
    }
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithBucketId(bucketId string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.BucketId = &bucketId
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithCanceledDueToTimeout(canceledDueToTimeout bool) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.CanceledDueToTimeout = &canceledDueToTimeout
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.ChunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithLastPage() *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.LastPage = true
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithName(name string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.Name = &name
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPageLength(pageLength int) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.PageLength = &pageLength
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPageOffset(pageOffset int) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.PageOffset = &pageOffset
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPriority(priority Priority) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.Priority = priority
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithRechunked(rechunked string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.Rechunked = &rechunked
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.RequestType = requestType
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithTruncated(truncated bool) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.Truncated = &truncated
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithUserId(userId string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.UserId = &userId
    return getCanceledJobsSpectraS3Request
}

