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

type PutGlobalDataPolicyAclForGroupSpectraS3Request struct {
    groupId string
    queryParams *url.Values
}

func NewPutGlobalDataPolicyAclForGroupSpectraS3Request(groupId string) *PutGlobalDataPolicyAclForGroupSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("group_id", groupId)

    return &PutGlobalDataPolicyAclForGroupSpectraS3Request{
        groupId: groupId,
        queryParams: queryParams,
    }
}




func (PutGlobalDataPolicyAclForGroupSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putGlobalDataPolicyAclForGroupSpectraS3Request *PutGlobalDataPolicyAclForGroupSpectraS3Request) Path() string {
    return "/_rest_/data_policy_acl"
}

func (putGlobalDataPolicyAclForGroupSpectraS3Request *PutGlobalDataPolicyAclForGroupSpectraS3Request) QueryParams() *url.Values {
    return putGlobalDataPolicyAclForGroupSpectraS3Request.queryParams
}

func (PutGlobalDataPolicyAclForGroupSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutGlobalDataPolicyAclForGroupSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutGlobalDataPolicyAclForGroupSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
