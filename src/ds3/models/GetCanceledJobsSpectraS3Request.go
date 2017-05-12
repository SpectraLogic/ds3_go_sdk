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

type GetCanceledJobsSpectraS3Request struct {
    bucketId string
    canceledDueToTimeout bool
    chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    priority Priority
    rechunked string
    requestType JobRequestType
    truncated bool
    userId string
    queryParams *url.Values
}

func NewGetCanceledJobsSpectraS3Request() *GetCanceledJobsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetCanceledJobsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithBucketId(bucketId string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.bucketId = bucketId
    getCanceledJobsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getCanceledJobsSpectraS3Request
}
func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithCanceledDueToTimeout(canceledDueToTimeout bool) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.canceledDueToTimeout = canceledDueToTimeout
    getCanceledJobsSpectraS3Request.queryParams.Set("canceled_due_to_timeout", strconv.FormatBool(canceledDueToTimeout))
    return getCanceledJobsSpectraS3Request
}
func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.chunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    getCanceledJobsSpectraS3Request.queryParams.Set("chunk_client_processing_order_guarantee", chunkClientProcessingOrderGuarantee.String())
    return getCanceledJobsSpectraS3Request
}
func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPageLength(pageLength int) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.pageLength = pageLength
    getCanceledJobsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getCanceledJobsSpectraS3Request
}
func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPageOffset(pageOffset int) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.pageOffset = pageOffset
    getCanceledJobsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getCanceledJobsSpectraS3Request
}
func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.pageStartMarker = pageStartMarker
    getCanceledJobsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getCanceledJobsSpectraS3Request
}
func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPriority(priority Priority) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.priority = priority
    getCanceledJobsSpectraS3Request.queryParams.Set("priority", priority.String())
    return getCanceledJobsSpectraS3Request
}
func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithRechunked(rechunked string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.rechunked = rechunked
    getCanceledJobsSpectraS3Request.queryParams.Set("rechunked", rechunked)
    return getCanceledJobsSpectraS3Request
}
func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.requestType = requestType
    getCanceledJobsSpectraS3Request.queryParams.Set("request_type", requestType.String())
    return getCanceledJobsSpectraS3Request
}
func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithTruncated(truncated bool) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.truncated = truncated
    getCanceledJobsSpectraS3Request.queryParams.Set("truncated", strconv.FormatBool(truncated))
    return getCanceledJobsSpectraS3Request
}
func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithUserId(userId string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.userId = userId
    getCanceledJobsSpectraS3Request.queryParams.Set("user_id", userId)
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithName(name *string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.name = name
    if name != nil {
        getCanceledJobsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getCanceledJobsSpectraS3Request.queryParams.Set("name", "")
    }
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithLastPage() *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.queryParams.Set("last_page", "")
    return getCanceledJobsSpectraS3Request
}

func (GetCanceledJobsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) Path() string {
    return "/_rest_/canceled_job"
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) QueryParams() *url.Values {
    return getCanceledJobsSpectraS3Request.queryParams
}

func (GetCanceledJobsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetCanceledJobsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetCanceledJobsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
