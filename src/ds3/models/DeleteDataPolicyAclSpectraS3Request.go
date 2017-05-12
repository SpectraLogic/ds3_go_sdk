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

type DeleteDataPolicyAclSpectraS3Request struct {
    dataPolicyAcl string
    queryParams *url.Values
}

func NewDeleteDataPolicyAclSpectraS3Request(dataPolicyAcl string) *DeleteDataPolicyAclSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteDataPolicyAclSpectraS3Request{
        dataPolicyAcl: dataPolicyAcl,
        queryParams: queryParams,
    }
}




func (DeleteDataPolicyAclSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteDataPolicyAclSpectraS3Request *DeleteDataPolicyAclSpectraS3Request) Path() string {
    return "/_rest_/data_policy_acl/" + deleteDataPolicyAclSpectraS3Request.dataPolicyAcl
}

func (deleteDataPolicyAclSpectraS3Request *DeleteDataPolicyAclSpectraS3Request) QueryParams() *url.Values {
    return deleteDataPolicyAclSpectraS3Request.queryParams
}

func (DeleteDataPolicyAclSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteDataPolicyAclSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteDataPolicyAclSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
