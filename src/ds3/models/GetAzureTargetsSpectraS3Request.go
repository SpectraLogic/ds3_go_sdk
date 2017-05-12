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

type GetAzureTargetsSpectraS3Request struct {
    accountName *string
    defaultReadPreference TargetReadPreferenceType
    https bool
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    permitGoingOutOfSync bool
    quiesced Quiesced
    state TargetState
    queryParams *url.Values
}

func NewGetAzureTargetsSpectraS3Request() *GetAzureTargetsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetAzureTargetsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.defaultReadPreference = defaultReadPreference
    getAzureTargetsSpectraS3Request.queryParams.Set("default_read_preference", defaultReadPreference.String())
    return getAzureTargetsSpectraS3Request
}
func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithHttps(https bool) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.https = https
    getAzureTargetsSpectraS3Request.queryParams.Set("https", strconv.FormatBool(https))
    return getAzureTargetsSpectraS3Request
}
func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.pageLength = pageLength
    getAzureTargetsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getAzureTargetsSpectraS3Request
}
func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.pageOffset = pageOffset
    getAzureTargetsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getAzureTargetsSpectraS3Request
}
func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.pageStartMarker = pageStartMarker
    getAzureTargetsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getAzureTargetsSpectraS3Request
}
func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.permitGoingOutOfSync = permitGoingOutOfSync
    getAzureTargetsSpectraS3Request.queryParams.Set("permit_going_out_of_sync", strconv.FormatBool(permitGoingOutOfSync))
    return getAzureTargetsSpectraS3Request
}
func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.quiesced = quiesced
    getAzureTargetsSpectraS3Request.queryParams.Set("quiesced", quiesced.String())
    return getAzureTargetsSpectraS3Request
}
func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithState(state TargetState) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.state = state
    getAzureTargetsSpectraS3Request.queryParams.Set("state", state.String())
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithAccountName(accountName *string) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.accountName = accountName
    if accountName != nil {
        getAzureTargetsSpectraS3Request.queryParams.Set("account_name", *accountName)
    } else {
        getAzureTargetsSpectraS3Request.queryParams.Set("account_name", "")
    }
    return getAzureTargetsSpectraS3Request
}
func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithName(name *string) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.name = name
    if name != nil {
        getAzureTargetsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getAzureTargetsSpectraS3Request.queryParams.Set("name", "")
    }
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithLastPage() *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.queryParams.Set("last_page", "")
    return getAzureTargetsSpectraS3Request
}

func (GetAzureTargetsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) Path() string {
    return "/_rest_/azure_target"
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) QueryParams() *url.Values {
    return getAzureTargetsSpectraS3Request.queryParams
}

func (GetAzureTargetsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetAzureTargetsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetAzureTargetsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
