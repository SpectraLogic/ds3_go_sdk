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

type GetDataPolicySpectraS3Request struct {
    dataPolicyId string
    queryParams *url.Values
}

func NewGetDataPolicySpectraS3Request(dataPolicyId string) *GetDataPolicySpectraS3Request {
    queryParams := &url.Values{}

    return &GetDataPolicySpectraS3Request{
        dataPolicyId: dataPolicyId,
        queryParams: queryParams,
    }
}




func (GetDataPolicySpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDataPolicySpectraS3Request *GetDataPolicySpectraS3Request) Path() string {
    return "/_rest_/data_policy/" + getDataPolicySpectraS3Request.dataPolicyId
}

func (getDataPolicySpectraS3Request *GetDataPolicySpectraS3Request) QueryParams() *url.Values {
    return getDataPolicySpectraS3Request.queryParams
}

func (GetDataPolicySpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDataPolicySpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDataPolicySpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
