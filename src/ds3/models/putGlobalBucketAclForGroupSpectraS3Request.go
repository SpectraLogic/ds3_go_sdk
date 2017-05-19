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

type PutGlobalBucketAclForGroupSpectraS3Request struct {
    groupId string
    permission BucketAclPermission
    queryParams *url.Values
}

func NewPutGlobalBucketAclForGroupSpectraS3Request(groupId string, permission BucketAclPermission) *PutGlobalBucketAclForGroupSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("group_id", groupId)
    queryParams.Set("permission", permission.String())

    return &PutGlobalBucketAclForGroupSpectraS3Request{
        groupId: groupId,
        permission: permission,
        queryParams: queryParams,
    }
}




func (PutGlobalBucketAclForGroupSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putGlobalBucketAclForGroupSpectraS3Request *PutGlobalBucketAclForGroupSpectraS3Request) Path() string {
    return "/_rest_/bucket_acl"
}

func (putGlobalBucketAclForGroupSpectraS3Request *PutGlobalBucketAclForGroupSpectraS3Request) QueryParams() *url.Values {
    return putGlobalBucketAclForGroupSpectraS3Request.queryParams
}

func (PutGlobalBucketAclForGroupSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutGlobalBucketAclForGroupSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutGlobalBucketAclForGroupSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
