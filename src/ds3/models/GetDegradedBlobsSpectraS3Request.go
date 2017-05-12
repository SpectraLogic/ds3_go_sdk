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

type GetDegradedBlobsSpectraS3Request struct {
    blobId string
    bucketId string
    ds3ReplicationRuleId string
    pageLength int
    pageOffset int
    pageStartMarker string
    persistenceRuleId string
    queryParams *url.Values
}

func NewGetDegradedBlobsSpectraS3Request() *GetDegradedBlobsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDegradedBlobsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithBlobId(blobId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.blobId = blobId
    getDegradedBlobsSpectraS3Request.queryParams.Set("blob_id", blobId)
    return getDegradedBlobsSpectraS3Request
}
func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithBucketId(bucketId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.bucketId = bucketId
    getDegradedBlobsSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getDegradedBlobsSpectraS3Request
}
func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithDs3ReplicationRuleId(ds3ReplicationRuleId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.ds3ReplicationRuleId = ds3ReplicationRuleId
    getDegradedBlobsSpectraS3Request.queryParams.Set("ds3_replication_rule_id", ds3ReplicationRuleId)
    return getDegradedBlobsSpectraS3Request
}
func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPageLength(pageLength int) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.pageLength = pageLength
    getDegradedBlobsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDegradedBlobsSpectraS3Request
}
func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.pageOffset = pageOffset
    getDegradedBlobsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDegradedBlobsSpectraS3Request
}
func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.pageStartMarker = pageStartMarker
    getDegradedBlobsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDegradedBlobsSpectraS3Request
}
func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPersistenceRuleId(persistenceRuleId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.persistenceRuleId = persistenceRuleId
    getDegradedBlobsSpectraS3Request.queryParams.Set("persistence_rule_id", persistenceRuleId)
    return getDegradedBlobsSpectraS3Request
}


func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithLastPage() *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.queryParams.Set("last_page", "")
    return getDegradedBlobsSpectraS3Request
}

func (GetDegradedBlobsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) Path() string {
    return "/_rest_/degraded_blob"
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) QueryParams() *url.Values {
    return getDegradedBlobsSpectraS3Request.queryParams
}

func (GetDegradedBlobsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDegradedBlobsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDegradedBlobsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
