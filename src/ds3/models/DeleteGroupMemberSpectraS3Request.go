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

type DeleteGroupMemberSpectraS3Request struct {
    groupMember string
    queryParams *url.Values
}

func NewDeleteGroupMemberSpectraS3Request(groupMember string) *DeleteGroupMemberSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteGroupMemberSpectraS3Request{
        groupMember: groupMember,
        queryParams: queryParams,
    }
}




func (DeleteGroupMemberSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteGroupMemberSpectraS3Request *DeleteGroupMemberSpectraS3Request) Path() string {
    return "/_rest_/group_member/" + deleteGroupMemberSpectraS3Request.groupMember
}

func (deleteGroupMemberSpectraS3Request *DeleteGroupMemberSpectraS3Request) QueryParams() *url.Values {
    return deleteGroupMemberSpectraS3Request.queryParams
}

func (DeleteGroupMemberSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteGroupMemberSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteGroupMemberSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
