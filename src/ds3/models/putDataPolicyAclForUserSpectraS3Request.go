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

type PutDataPolicyAclForUserSpectraS3Request struct {
    dataPolicyId string
    userId string
    queryParams *url.Values
}

func NewPutDataPolicyAclForUserSpectraS3Request(dataPolicyId string, userId string) *PutDataPolicyAclForUserSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("data_policy_id", dataPolicyId)
    queryParams.Set("user_id", userId)

    return &PutDataPolicyAclForUserSpectraS3Request{
        dataPolicyId: dataPolicyId,
        userId: userId,
        queryParams: queryParams,
    }
}




func (PutDataPolicyAclForUserSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putDataPolicyAclForUserSpectraS3Request *PutDataPolicyAclForUserSpectraS3Request) Path() string {
    return "/_rest_/data_policy_acl"
}

func (putDataPolicyAclForUserSpectraS3Request *PutDataPolicyAclForUserSpectraS3Request) QueryParams() *url.Values {
    return putDataPolicyAclForUserSpectraS3Request.queryParams
}

func (PutDataPolicyAclForUserSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutDataPolicyAclForUserSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutDataPolicyAclForUserSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
