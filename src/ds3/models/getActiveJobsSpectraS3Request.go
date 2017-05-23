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

type GetActiveJobsSpectraS3Request struct {
    aggregating bool
    bucketId string
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

func NewGetActiveJobsSpectraS3Request() *GetActiveJobsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetActiveJobsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithAggregating(aggregating bool) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.aggregating = aggregating
    getActiveJobsSpectraS3Request.queryParams.Set("aggregating", strconv.FormatBool(aggregating))
    return getActiveJobsSpectraS3Request
}
func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithBucketId(bucketId string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.bucketId = bucketId
    getActiveJobsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getActiveJobsSpectraS3Request
}
func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.chunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    getActiveJobsSpectraS3Request.queryParams.Set("chunk_client_processing_order_guarantee", chunkClientProcessingOrderGuarantee.String())
    return getActiveJobsSpectraS3Request
}
func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPageLength(pageLength int) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.pageLength = pageLength
    getActiveJobsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getActiveJobsSpectraS3Request
}
func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPageOffset(pageOffset int) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.pageOffset = pageOffset
    getActiveJobsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getActiveJobsSpectraS3Request
}
func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.pageStartMarker = pageStartMarker
    getActiveJobsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getActiveJobsSpectraS3Request
}
func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPriority(priority Priority) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.priority = priority
    getActiveJobsSpectraS3Request.queryParams.Set("priority", priority.String())
    return getActiveJobsSpectraS3Request
}
func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithRechunked(rechunked string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.rechunked = rechunked
    getActiveJobsSpectraS3Request.queryParams.Set("rechunked", rechunked)
    return getActiveJobsSpectraS3Request
}
func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.requestType = requestType
    getActiveJobsSpectraS3Request.queryParams.Set("request_type", requestType.String())
    return getActiveJobsSpectraS3Request
}
func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithTruncated(truncated bool) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.truncated = truncated
    getActiveJobsSpectraS3Request.queryParams.Set("truncated", strconv.FormatBool(truncated))
    return getActiveJobsSpectraS3Request
}
func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithUserId(userId string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.userId = userId
    getActiveJobsSpectraS3Request.queryParams.Set("user_id", userId)
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithName(name *string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.name = name
    if name != nil {
        getActiveJobsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getActiveJobsSpectraS3Request.queryParams.Set("name", "")
    }
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithLastPage() *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.queryParams.Set("last_page", "")
    return getActiveJobsSpectraS3Request
}

func (GetActiveJobsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) Path() string {
    return "/_rest_/active_job"
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) QueryParams() *url.Values {
    return getActiveJobsSpectraS3Request.queryParams
}

func (GetActiveJobsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetActiveJobsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetActiveJobsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
