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

type GetTapePartitionFailuresSpectraS3Request struct {
    errorMessage *string
    pageLength int
    pageOffset int
    pageStartMarker string
    partitionId string
    tapePartitionFailureType TapePartitionFailureType
    queryParams *url.Values
}

func NewGetTapePartitionFailuresSpectraS3Request() *GetTapePartitionFailuresSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapePartitionFailuresSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.pageLength = pageLength
    getTapePartitionFailuresSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getTapePartitionFailuresSpectraS3Request
}
func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.pageOffset = pageOffset
    getTapePartitionFailuresSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getTapePartitionFailuresSpectraS3Request
}
func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.pageStartMarker = pageStartMarker
    getTapePartitionFailuresSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getTapePartitionFailuresSpectraS3Request
}
func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPartitionId(partitionId string) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.partitionId = partitionId
    getTapePartitionFailuresSpectraS3Request.queryParams.Set("partition_id", partitionId)
    return getTapePartitionFailuresSpectraS3Request
}
func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithTapePartitionFailureType(tapePartitionFailureType TapePartitionFailureType) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.tapePartitionFailureType = tapePartitionFailureType
    getTapePartitionFailuresSpectraS3Request.queryParams.Set("type", tapePartitionFailureType.String())
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithErrorMessage(errorMessage *string) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.errorMessage = errorMessage
    if errorMessage != nil {
        getTapePartitionFailuresSpectraS3Request.queryParams.Set("error_message", *errorMessage)
    } else {
        getTapePartitionFailuresSpectraS3Request.queryParams.Set("error_message", "")
    }
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithLastPage() *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.queryParams.Set("last_page", "")
    return getTapePartitionFailuresSpectraS3Request
}

func (GetTapePartitionFailuresSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) Path() string {
    return "/_rest_/tape_partition_failure"
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) QueryParams() *url.Values {
    return getTapePartitionFailuresSpectraS3Request.queryParams
}

func (GetTapePartitionFailuresSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapePartitionFailuresSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapePartitionFailuresSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
