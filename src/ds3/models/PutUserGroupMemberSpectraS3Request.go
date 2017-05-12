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
)

type PutUserGroupMemberSpectraS3Request struct {
    groupId string
    memberUserId string
    queryParams *url.Values
}

func NewPutUserGroupMemberSpectraS3Request(groupId string, memberUserId string) *PutUserGroupMemberSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("group_id", groupId)
    queryParams.Set("member_user_id", memberUserId)

    return &PutUserGroupMemberSpectraS3Request{
        groupId: groupId,
        memberUserId: memberUserId,
        queryParams: queryParams,
    }
}




func (PutUserGroupMemberSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putUserGroupMemberSpectraS3Request *PutUserGroupMemberSpectraS3Request) Path() string {
    return "/_rest_/group_member"
}

func (putUserGroupMemberSpectraS3Request *PutUserGroupMemberSpectraS3Request) QueryParams() *url.Values {
    return putUserGroupMemberSpectraS3Request.queryParams
}

func (PutUserGroupMemberSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutUserGroupMemberSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutUserGroupMemberSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
