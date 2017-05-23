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

type GetSuspectBlobPoolsSpectraS3Request struct {
    blobId string
    pageLength int
    pageOffset int
    pageStartMarker string
    poolId string
    queryParams *url.Values
}

func NewGetSuspectBlobPoolsSpectraS3Request() *GetSuspectBlobPoolsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetSuspectBlobPoolsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.blobId = blobId
    getSuspectBlobPoolsSpectraS3Request.queryParams.Set("blob_id", blobId)
    return getSuspectBlobPoolsSpectraS3Request
}
func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.pageLength = pageLength
    getSuspectBlobPoolsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getSuspectBlobPoolsSpectraS3Request
}
func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.pageOffset = pageOffset
    getSuspectBlobPoolsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getSuspectBlobPoolsSpectraS3Request
}
func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.pageStartMarker = pageStartMarker
    getSuspectBlobPoolsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getSuspectBlobPoolsSpectraS3Request
}
func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithPoolId(poolId string) *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.poolId = poolId
    getSuspectBlobPoolsSpectraS3Request.queryParams.Set("pool_id", poolId)
    return getSuspectBlobPoolsSpectraS3Request
}


func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithLastPage() *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.queryParams.Set("last_page", "")
    return getSuspectBlobPoolsSpectraS3Request
}

func (GetSuspectBlobPoolsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_pool"
}

func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) QueryParams() *url.Values {
    return getSuspectBlobPoolsSpectraS3Request.queryParams
}

func (GetSuspectBlobPoolsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetSuspectBlobPoolsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetSuspectBlobPoolsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
