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

type GetDataPolicyAclsSpectraS3Request struct {
    dataPolicyId string
    groupId string
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetDataPolicyAclsSpectraS3Request() *GetDataPolicyAclsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDataPolicyAclsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.dataPolicyId = dataPolicyId
    getDataPolicyAclsSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return getDataPolicyAclsSpectraS3Request
}
func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithGroupId(groupId string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.groupId = groupId
    getDataPolicyAclsSpectraS3Request.queryParams.Set("group_id", groupId)
    return getDataPolicyAclsSpectraS3Request
}
func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithPageLength(pageLength int) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.pageLength = pageLength
    getDataPolicyAclsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDataPolicyAclsSpectraS3Request
}
func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithPageOffset(pageOffset int) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.pageOffset = pageOffset
    getDataPolicyAclsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDataPolicyAclsSpectraS3Request
}
func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.pageStartMarker = pageStartMarker
    getDataPolicyAclsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDataPolicyAclsSpectraS3Request
}
func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithUserId(userId string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.userId = userId
    getDataPolicyAclsSpectraS3Request.queryParams.Set("user_id", userId)
    return getDataPolicyAclsSpectraS3Request
}


func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithLastPage() *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.queryParams.Set("last_page", "")
    return getDataPolicyAclsSpectraS3Request
}

func (GetDataPolicyAclsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) Path() string {
    return "/_rest_/data_policy_acl"
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) QueryParams() *url.Values {
    return getDataPolicyAclsSpectraS3Request.queryParams
}

func (GetDataPolicyAclsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDataPolicyAclsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDataPolicyAclsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
