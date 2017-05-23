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

type GetAzureTargetReadPreferencesSpectraS3Request struct {
    bucketId string
    pageLength int
    pageOffset int
    pageStartMarker string
    readPreference TargetReadPreferenceType
    targetId string
    queryParams *url.Values
}

func NewGetAzureTargetReadPreferencesSpectraS3Request() *GetAzureTargetReadPreferencesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetAzureTargetReadPreferencesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithBucketId(bucketId string) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.bucketId = bucketId
    getAzureTargetReadPreferencesSpectraS3Request.queryParams.Set("bucket_id", bucketId)
    return getAzureTargetReadPreferencesSpectraS3Request
}
func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.pageLength = pageLength
    getAzureTargetReadPreferencesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getAzureTargetReadPreferencesSpectraS3Request
}
func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.pageOffset = pageOffset
    getAzureTargetReadPreferencesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getAzureTargetReadPreferencesSpectraS3Request
}
func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.pageStartMarker = pageStartMarker
    getAzureTargetReadPreferencesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getAzureTargetReadPreferencesSpectraS3Request
}
func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithReadPreference(readPreference TargetReadPreferenceType) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.readPreference = readPreference
    getAzureTargetReadPreferencesSpectraS3Request.queryParams.Set("read_preference", readPreference.String())
    return getAzureTargetReadPreferencesSpectraS3Request
}
func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithTargetId(targetId string) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.targetId = targetId
    getAzureTargetReadPreferencesSpectraS3Request.queryParams.Set("target_id", targetId)
    return getAzureTargetReadPreferencesSpectraS3Request
}


func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithLastPage() *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.queryParams.Set("last_page", "")
    return getAzureTargetReadPreferencesSpectraS3Request
}

func (GetAzureTargetReadPreferencesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) Path() string {
    return "/_rest_/azure_target_read_preference"
}

func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) QueryParams() *url.Values {
    return getAzureTargetReadPreferencesSpectraS3Request.queryParams
}

func (GetAzureTargetReadPreferencesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetAzureTargetReadPreferencesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetAzureTargetReadPreferencesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
