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

type GetUsersSpectraS3Request struct {
    authId *string
    defaultDataPolicyId string
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    queryParams *url.Values
}

func NewGetUsersSpectraS3Request() *GetUsersSpectraS3Request {
    queryParams := &url.Values{}

    return &GetUsersSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithDefaultDataPolicyId(defaultDataPolicyId string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.defaultDataPolicyId = defaultDataPolicyId
    getUsersSpectraS3Request.queryParams.Set("default_data_policy_id", defaultDataPolicyId)
    return getUsersSpectraS3Request
}
func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithPageLength(pageLength int) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.pageLength = pageLength
    getUsersSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getUsersSpectraS3Request
}
func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithPageOffset(pageOffset int) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.pageOffset = pageOffset
    getUsersSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getUsersSpectraS3Request
}
func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.pageStartMarker = pageStartMarker
    getUsersSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithAuthId(authId *string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.authId = authId
    if authId != nil {
        getUsersSpectraS3Request.queryParams.Set("auth_id", *authId)
    } else {
        getUsersSpectraS3Request.queryParams.Set("auth_id", "")
    }
    return getUsersSpectraS3Request
}
func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithName(name *string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.name = name
    if name != nil {
        getUsersSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getUsersSpectraS3Request.queryParams.Set("name", "")
    }
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithLastPage() *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.queryParams.Set("last_page", "")
    return getUsersSpectraS3Request
}

func (GetUsersSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) Path() string {
    return "/_rest_/user"
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) QueryParams() *url.Values {
    return getUsersSpectraS3Request.queryParams
}

func (GetUsersSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetUsersSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetUsersSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
