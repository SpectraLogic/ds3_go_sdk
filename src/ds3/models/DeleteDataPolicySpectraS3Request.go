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

type DeleteDataPolicySpectraS3Request struct {
    dataPolicyId string
    queryParams *url.Values
}

func NewDeleteDataPolicySpectraS3Request(dataPolicyId string) *DeleteDataPolicySpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteDataPolicySpectraS3Request{
        dataPolicyId: dataPolicyId,
        queryParams: queryParams,
    }
}




func (DeleteDataPolicySpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteDataPolicySpectraS3Request *DeleteDataPolicySpectraS3Request) Path() string {
    return "/_rest_/data_policy/" + deleteDataPolicySpectraS3Request.dataPolicyId
}

func (deleteDataPolicySpectraS3Request *DeleteDataPolicySpectraS3Request) QueryParams() *url.Values {
    return deleteDataPolicySpectraS3Request.queryParams
}

func (DeleteDataPolicySpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteDataPolicySpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteDataPolicySpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
