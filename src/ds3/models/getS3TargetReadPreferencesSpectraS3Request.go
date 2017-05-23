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

type GetS3TargetReadPreferencesSpectraS3Request struct {
    bucketId string
    pageLength int
    pageOffset int
    pageStartMarker string
    readPreference TargetReadPreferenceType
    targetId string
    queryParams *url.Values
}

func NewGetS3TargetReadPreferencesSpectraS3Request() *GetS3TargetReadPreferencesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetS3TargetReadPreferencesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithBucketId(bucketId string) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.bucketId = bucketId
    getS3TargetReadPreferencesSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getS3TargetReadPreferencesSpectraS3Request
}
func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.pageLength = pageLength
    getS3TargetReadPreferencesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getS3TargetReadPreferencesSpectraS3Request
}
func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.pageOffset = pageOffset
    getS3TargetReadPreferencesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getS3TargetReadPreferencesSpectraS3Request
}
func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.pageStartMarker = pageStartMarker
    getS3TargetReadPreferencesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getS3TargetReadPreferencesSpectraS3Request
}
func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithReadPreference(readPreference TargetReadPreferenceType) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.readPreference = readPreference
    getS3TargetReadPreferencesSpectraS3Request.queryParams.Set("read_preference", readPreference.String())
    return getS3TargetReadPreferencesSpectraS3Request
}
func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithTargetId(targetId string) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.targetId = targetId
    getS3TargetReadPreferencesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getS3TargetReadPreferencesSpectraS3Request
}


func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithLastPage() *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.queryParams.Set("last_page", "")
    return getS3TargetReadPreferencesSpectraS3Request
}

func (GetS3TargetReadPreferencesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) Path() string {
    return "/_rest_/s3_target_read_preference"
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) QueryParams() *url.Values {
    return getS3TargetReadPreferencesSpectraS3Request.queryParams
}

func (GetS3TargetReadPreferencesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetS3TargetReadPreferencesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetS3TargetReadPreferencesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
