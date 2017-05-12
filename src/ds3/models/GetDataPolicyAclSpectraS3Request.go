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

type GetDataPolicyAclSpectraS3Request struct {
    dataPolicyAcl string
    queryParams *url.Values
}

func NewGetDataPolicyAclSpectraS3Request(dataPolicyAcl string) *GetDataPolicyAclSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDataPolicyAclSpectraS3Request{
        dataPolicyAcl: dataPolicyAcl,
        queryParams: queryParams,
    }
}




func (GetDataPolicyAclSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDataPolicyAclSpectraS3Request *GetDataPolicyAclSpectraS3Request) Path() string {
    return "/_rest_/data_policy_acl/" + getDataPolicyAclSpectraS3Request.dataPolicyAcl
}

func (getDataPolicyAclSpectraS3Request *GetDataPolicyAclSpectraS3Request) QueryParams() *url.Values {
    return getDataPolicyAclSpectraS3Request.queryParams
}

func (GetDataPolicyAclSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDataPolicyAclSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDataPolicyAclSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
