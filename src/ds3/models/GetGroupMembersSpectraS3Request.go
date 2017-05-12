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

type GetGroupMembersSpectraS3Request struct {
    groupId string
    memberGroupId string
    memberUserId string
    pageLength int
    pageOffset int
    pageStartMarker string
    queryParams *url.Values
}

func NewGetGroupMembersSpectraS3Request() *GetGroupMembersSpectraS3Request {
    queryParams := &url.Values{}

    return &GetGroupMembersSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithGroupId(groupId string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.groupId = groupId
    getGroupMembersSpectraS3Request.queryParams.Set("group_id", groupId)
    return getGroupMembersSpectraS3Request
}
func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithMemberGroupId(memberGroupId string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.memberGroupId = memberGroupId
    getGroupMembersSpectraS3Request.queryParams.Set("member_group_id", memberGroupId)
    return getGroupMembersSpectraS3Request
}
func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithMemberUserId(memberUserId string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.memberUserId = memberUserId
    getGroupMembersSpectraS3Request.queryParams.Set("member_user_id", memberUserId)
    return getGroupMembersSpectraS3Request
}
func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithPageLength(pageLength int) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.pageLength = pageLength
    getGroupMembersSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getGroupMembersSpectraS3Request
}
func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithPageOffset(pageOffset int) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.pageOffset = pageOffset
    getGroupMembersSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getGroupMembersSpectraS3Request
}
func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.pageStartMarker = pageStartMarker
    getGroupMembersSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getGroupMembersSpectraS3Request
}


func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithLastPage() *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.queryParams.Set("last_page", "")
    return getGroupMembersSpectraS3Request
}

func (GetGroupMembersSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) Path() string {
    return "/_rest_/group_member"
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) QueryParams() *url.Values {
    return getGroupMembersSpectraS3Request.queryParams
}

func (GetGroupMembersSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetGroupMembersSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetGroupMembersSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
