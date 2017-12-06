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

type GetGroupMembersSpectraS3Request struct {
    GroupId *string
    LastPage bool
    MemberGroupId *string
    MemberUserId *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetGroupMembersSpectraS3Request() *GetGroupMembersSpectraS3Request {
    return &GetGroupMembersSpectraS3Request{
    }
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithGroupId(groupId string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.GroupId = &groupId
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithLastPage() *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.LastPage = true
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithMemberGroupId(memberGroupId string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.MemberGroupId = &memberGroupId
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithMemberUserId(memberUserId string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.MemberUserId = &memberUserId
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithPageLength(pageLength int) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.PageLength = &pageLength
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithPageOffset(pageOffset int) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.PageOffset = &pageOffset
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.PageStartMarker = &pageStartMarker
    return getGroupMembersSpectraS3Request
}

