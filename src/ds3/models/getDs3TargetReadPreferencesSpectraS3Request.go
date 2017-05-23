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

type GetDs3TargetReadPreferencesSpectraS3Request struct {
    bucketId string
    pageLength int
    pageOffset int
    pageStartMarker string
    readPreference TargetReadPreferenceType
    targetId string
    queryParams *url.Values
}

func NewGetDs3TargetReadPreferencesSpectraS3Request() *GetDs3TargetReadPreferencesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDs3TargetReadPreferencesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithBucketId(bucketId string) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.bucketId = bucketId
    getDs3TargetReadPreferencesSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getDs3TargetReadPreferencesSpectraS3Request
}
func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.pageLength = pageLength
    getDs3TargetReadPreferencesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDs3TargetReadPreferencesSpectraS3Request
}
func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.pageOffset = pageOffset
    getDs3TargetReadPreferencesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDs3TargetReadPreferencesSpectraS3Request
}
func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.pageStartMarker = pageStartMarker
    getDs3TargetReadPreferencesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDs3TargetReadPreferencesSpectraS3Request
}
func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithReadPreference(readPreference TargetReadPreferenceType) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.readPreference = readPreference
    getDs3TargetReadPreferencesSpectraS3Request.queryParams.Set("read_preference", readPreference.String())
    return getDs3TargetReadPreferencesSpectraS3Request
}
func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithTargetId(targetId string) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.targetId = targetId
    getDs3TargetReadPreferencesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getDs3TargetReadPreferencesSpectraS3Request
}


func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithLastPage() *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.queryParams.Set("last_page", "")
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (GetDs3TargetReadPreferencesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) Path() string {
    return "/_rest_/ds3_target_read_preference"
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) QueryParams() *url.Values {
    return getDs3TargetReadPreferencesSpectraS3Request.queryParams
}

func (GetDs3TargetReadPreferencesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDs3TargetReadPreferencesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDs3TargetReadPreferencesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
