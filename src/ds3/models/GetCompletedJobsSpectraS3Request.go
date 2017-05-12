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

type GetCompletedJobsSpectraS3Request struct {
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

func NewGetCompletedJobsSpectraS3Request() *GetCompletedJobsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetCompletedJobsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithBucketId(bucketId string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.bucketId = bucketId
    getCompletedJobsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getCompletedJobsSpectraS3Request
}
func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.chunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    getCompletedJobsSpectraS3Request.queryParams.Set("chunk_client_processing_order_guarantee", chunkClientProcessingOrderGuarantee.String())
    return getCompletedJobsSpectraS3Request
}
func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPageLength(pageLength int) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.pageLength = pageLength
    getCompletedJobsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getCompletedJobsSpectraS3Request
}
func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPageOffset(pageOffset int) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.pageOffset = pageOffset
    getCompletedJobsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getCompletedJobsSpectraS3Request
}
func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.pageStartMarker = pageStartMarker
    getCompletedJobsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getCompletedJobsSpectraS3Request
}
func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPriority(priority Priority) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.priority = priority
    getCompletedJobsSpectraS3Request.queryParams.Set("priority", priority.String())
    return getCompletedJobsSpectraS3Request
}
func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithRechunked(rechunked string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.rechunked = rechunked
    getCompletedJobsSpectraS3Request.queryParams.Set("rechunked", rechunked)
    return getCompletedJobsSpectraS3Request
}
func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.requestType = requestType
    getCompletedJobsSpectraS3Request.queryParams.Set("request_type", requestType.String())
    return getCompletedJobsSpectraS3Request
}
func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithTruncated(truncated bool) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.truncated = truncated
    getCompletedJobsSpectraS3Request.queryParams.Set("truncated", strconv.FormatBool(truncated))
    return getCompletedJobsSpectraS3Request
}
func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithUserId(userId string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.userId = userId
    getCompletedJobsSpectraS3Request.queryParams.Set("user_id", userId)
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithName(name *string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.name = name
    if name != nil {
        getCompletedJobsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getCompletedJobsSpectraS3Request.queryParams.Set("name", "")
    }
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithLastPage() *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.queryParams.Set("last_page", "")
    return getCompletedJobsSpectraS3Request
}

func (GetCompletedJobsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) Path() string {
    return "/_rest_/completed_job"
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) QueryParams() *url.Values {
    return getCompletedJobsSpectraS3Request.queryParams
}

func (GetCompletedJobsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetCompletedJobsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetCompletedJobsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
