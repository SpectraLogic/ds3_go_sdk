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

type GetS3TargetsSpectraS3Request struct {
    accessKey *string
    dataPathEndPoint *string
    defaultReadPreference TargetReadPreferenceType
    https bool
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    permitGoingOutOfSync bool
    quiesced Quiesced
    region S3Region
    state TargetState
    queryParams *url.Values
}

func NewGetS3TargetsSpectraS3Request() *GetS3TargetsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetS3TargetsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.defaultReadPreference = defaultReadPreference
    getS3TargetsSpectraS3Request.queryParams.Set("default_read_preference", defaultReadPreference.String())
    return getS3TargetsSpectraS3Request
}
func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithHttps(https bool) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.https = https
    getS3TargetsSpectraS3Request.queryParams.Set("https", strconv.FormatBool(https))
    return getS3TargetsSpectraS3Request
}
func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.pageLength = pageLength
    getS3TargetsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getS3TargetsSpectraS3Request
}
func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.pageOffset = pageOffset
    getS3TargetsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getS3TargetsSpectraS3Request
}
func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.pageStartMarker = pageStartMarker
    getS3TargetsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getS3TargetsSpectraS3Request
}
func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.permitGoingOutOfSync = permitGoingOutOfSync
    getS3TargetsSpectraS3Request.queryParams.Set("permit_going_out_of_sync", strconv.FormatBool(permitGoingOutOfSync))
    return getS3TargetsSpectraS3Request
}
func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.quiesced = quiesced
    getS3TargetsSpectraS3Request.queryParams.Set("quiesced", quiesced.String())
    return getS3TargetsSpectraS3Request
}
func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithRegion(region S3Region) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.region = region
    getS3TargetsSpectraS3Request.queryParams.Set("region", region.String())
    return getS3TargetsSpectraS3Request
}
func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithState(state TargetState) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.state = state
    getS3TargetsSpectraS3Request.queryParams.Set("state", state.String())
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithAccessKey(accessKey *string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.accessKey = accessKey
    if accessKey != nil {
        getS3TargetsSpectraS3Request.queryParams.Set("access_key", *accessKey)
    } else {
        getS3TargetsSpectraS3Request.queryParams.Set("access_key", "")
    }
    return getS3TargetsSpectraS3Request
}
func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint *string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.dataPathEndPoint = dataPathEndPoint
    if dataPathEndPoint != nil {
        getS3TargetsSpectraS3Request.queryParams.Set("data_path_end_point", *dataPathEndPoint)
    } else {
        getS3TargetsSpectraS3Request.queryParams.Set("data_path_end_point", "")
    }
    return getS3TargetsSpectraS3Request
}
func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithName(name *string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.name = name
    if name != nil {
        getS3TargetsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getS3TargetsSpectraS3Request.queryParams.Set("name", "")
    }
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithLastPage() *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.queryParams.Set("last_page", "")
    return getS3TargetsSpectraS3Request
}

func (GetS3TargetsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) Path() string {
    return "/_rest_/s3_target"
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) QueryParams() *url.Values {
    return getS3TargetsSpectraS3Request.queryParams
}

func (GetS3TargetsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetS3TargetsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetS3TargetsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
