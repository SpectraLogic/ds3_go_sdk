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

type ModifyUserSpectraS3Request struct {
    defaultDataPolicyId string
    name *string
    secretKey *string
    userId string
    queryParams *url.Values
}

func NewModifyUserSpectraS3Request(userId string) *ModifyUserSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyUserSpectraS3Request{
        userId: userId,
        queryParams: queryParams,
    }
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) WithDefaultDataPolicyId(defaultDataPolicyId string) *ModifyUserSpectraS3Request {
    modifyUserSpectraS3Request.defaultDataPolicyId = defaultDataPolicyId
    modifyUserSpectraS3Request.queryParams.Set("default_data_policy_id", defaultDataPolicyId)
    return modifyUserSpectraS3Request
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) WithName(name *string) *ModifyUserSpectraS3Request {
    modifyUserSpectraS3Request.name = name
    if name != nil {
        modifyUserSpectraS3Request.queryParams.Set("name", *name)
    } else {
        modifyUserSpectraS3Request.queryParams.Set("name", "")
    }
    return modifyUserSpectraS3Request
}
func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) WithSecretKey(secretKey *string) *ModifyUserSpectraS3Request {
    modifyUserSpectraS3Request.secretKey = secretKey
    if secretKey != nil {
        modifyUserSpectraS3Request.queryParams.Set("secret_key", *secretKey)
    } else {
        modifyUserSpectraS3Request.queryParams.Set("secret_key", "")
    }
    return modifyUserSpectraS3Request
}


func (ModifyUserSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) Path() string {
    return "/_rest_/user/" + modifyUserSpectraS3Request.userId
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) QueryParams() *url.Values {
    return modifyUserSpectraS3Request.queryParams
}

func (ModifyUserSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyUserSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyUserSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
