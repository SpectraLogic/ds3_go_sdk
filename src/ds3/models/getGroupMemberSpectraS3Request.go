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

type GetGroupMemberSpectraS3Request struct {
    groupMember string
    queryParams *url.Values
}

func NewGetGroupMemberSpectraS3Request(groupMember string) *GetGroupMemberSpectraS3Request {
    queryParams := &url.Values{}

    return &GetGroupMemberSpectraS3Request{
        groupMember: groupMember,
        queryParams: queryParams,
    }
}




func (GetGroupMemberSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getGroupMemberSpectraS3Request *GetGroupMemberSpectraS3Request) Path() string {
    return "/_rest_/group_member/" + getGroupMemberSpectraS3Request.groupMember
}

func (getGroupMemberSpectraS3Request *GetGroupMemberSpectraS3Request) QueryParams() *url.Values {
    return getGroupMemberSpectraS3Request.queryParams
}

func (GetGroupMemberSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetGroupMemberSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetGroupMemberSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
