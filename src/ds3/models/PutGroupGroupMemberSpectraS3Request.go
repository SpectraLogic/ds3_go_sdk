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

type PutGroupGroupMemberSpectraS3Request struct {
    groupId string
    memberGroupId string
    queryParams *url.Values
}

func NewPutGroupGroupMemberSpectraS3Request(groupId string, memberGroupId string) *PutGroupGroupMemberSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("group_id", groupId)
    queryParams.Set("member_group_id", memberGroupId)

    return &PutGroupGroupMemberSpectraS3Request{
        groupId: groupId,
        memberGroupId: memberGroupId,
        queryParams: queryParams,
    }
}




func (PutGroupGroupMemberSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putGroupGroupMemberSpectraS3Request *PutGroupGroupMemberSpectraS3Request) Path() string {
    return "/_rest_/group_member"
}

func (putGroupGroupMemberSpectraS3Request *PutGroupGroupMemberSpectraS3Request) QueryParams() *url.Values {
    return putGroupGroupMemberSpectraS3Request.queryParams
}

func (PutGroupGroupMemberSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutGroupGroupMemberSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutGroupGroupMemberSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
